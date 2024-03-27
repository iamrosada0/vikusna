package repositories

import (
	"evaeats/domain"
	"fmt"

	"gorm.io/gorm"
)

type HomeChefEatsRepository interface {
	Insert(HomeChefEatsImage, HomeChefEatsName, phoneNumber, address, locationID, profileID, registrationStatus string) (*domain.HomeChefEats, error)
	Find(id string) (*domain.HomeChefEats, error)
	Update(id string, HomeChefEats *domain.HomeChefEats) (*domain.HomeChefEats, error)
	GetOne() ([]*domain.HomeChefEats, error)
	GetByID(id string) (map[string]interface{}, error)
	Delete(id string) error
}

type HomeChefEatsRepositoryDb struct {
	Db *gorm.DB
}

func NewHomeChefEatsRepository(Db *gorm.DB) *HomeChefEatsRepositoryDb {
	return &HomeChefEatsRepositoryDb{Db: Db}
}

func (repo *HomeChefEatsRepositoryDb) Insert(HomeChefEatsImage, HomeChefEatsName, phoneNumber, address, locationID, profileID, registrationStatus string) (*domain.HomeChefEats, error) {
	var dbChef domain.HomeChefEats
	repo.Db.Where("homeChefEats_name = ?", HomeChefEatsName).First(&dbChef)
	if dbChef.ID != "" {
		return nil, fmt.Errorf("HomeChefEats with the name %s already exists", HomeChefEatsName)
	}

	dbChef.Registration_status = "PENDING"
	if err := repo.Db.Create(&dbChef).Error; err != nil {
		return nil, err
	}

	chef := &domain.HomeChefEats{
		HomeChefEats_image:  HomeChefEatsImage,
		HomeChefEats_name:   HomeChefEatsName,
		Phone_number:        phoneNumber,
		Address:             address,
		LocationID:          locationID,
		ProfileID:           profileID,
		Registration_status: registrationStatus,
	}
	if err := repo.Db.Create(chef).Error; err != nil {
		return nil, err
	}
	return chef, nil
}

func (repo *HomeChefEatsRepositoryDb) Find(id string) (*domain.HomeChefEats, error) {
	var chef domain.HomeChefEats
	if err := repo.Db.First(&chef, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &chef, nil
}

func (repo *HomeChefEatsRepositoryDb) Update(id string, HomeChefEats *domain.HomeChefEats) (*domain.HomeChefEats, error) {
	var dbHomeChefEats domain.HomeChefEats

	if err := repo.Db.First(&dbHomeChefEats, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if err := repo.Db.Model(&dbHomeChefEats).Updates(HomeChefEats).Error; err != nil {
		return nil, err
	}

	return HomeChefEats, nil
}

func (repo *HomeChefEatsRepositoryDb) GetOne() ([]map[string]interface{}, error) {
	// Retrieve HomeChefEats objects with preloaded Profile
	HomeChefEats, err := repo.getHomeChefEatsWithProfile()
	if err != nil {
		return nil, err
	}

	// Map HomeChefEats objects to maps
	restaurantsHolder := repo.mapHomeChefEatsToMaps(HomeChefEats)

	return restaurantsHolder, nil
}

// getHomeChefEatsWithProfile retrieves HomeChefEats objects with preloaded Profile
func (repo *HomeChefEatsRepositoryDb) getHomeChefEatsWithProfile() ([]*domain.HomeChefEats, error) {
	var HomeChefEats []*domain.HomeChefEats
	if err := repo.Db.Preload("Profile").Find(&HomeChefEats).Error; err != nil {
		return nil, err
	}
	return HomeChefEats, nil
}

// mapHomeChefEatsToMaps maps HomeChefEats objects to maps
func (repo *HomeChefEatsRepositoryDb) mapHomeChefEatsToMaps(HomeChefEats []*domain.HomeChefEats) []map[string]interface{} {
	var restaurantsHolder []map[string]interface{}

	for _, chefEvaEat := range HomeChefEats {
		HomeChefEatsData := map[string]interface{}{
			"id":                 chefEvaEat.ID,
			"homeChefEats_image": chefEvaEat.HomeChefEats_image,
			"homeChefEats_name":  chefEvaEat.HomeChefEats_name,
			"phone_number":       chefEvaEat.Phone_number,
			"address":            chefEvaEat.Address,
			"location":           chefEvaEat.LocationID,
			"owner":              nil, // Placeholder for owner profile
		}

		// Check if HomeChefEats has associated Profile
		if chefEvaEat.Profile.ID != "" {
			profileData := map[string]interface{}{
				"id": chefEvaEat.Profile.UserID,
				// Include other profile fields as needed
			}
			HomeChefEatsData["owner"] = profileData
		}

		// Append HomeChefEatsData to restaurantsHolder
		restaurantsHolder = append(restaurantsHolder, HomeChefEatsData)
	}

	return restaurantsHolder
}

func (repo *HomeChefEatsRepositoryDb) GetByID(id string) (map[string]interface{}, error) {
	var chef domain.HomeChefEats
	if err := repo.Db.Preload("Profile").First(&chef, "id = ?", id).Error; err != nil {
		return nil, err
	}

	// Convert HomeChefEats to map
	chefData := map[string]interface{}{
		"id":                chef.ID,
		"homeChefEats_name": chef.HomeChefEats_name,
		// Add other HomeChefEats fields as needed
	}

	// Convert Profile to map
	profileData := map[string]interface{}{
		"id":         chef.Profile.ID,
		"first_name": chef.Profile.First_name,
		"last_name":  chef.Profile.Last_name,
		// Add other Profile fields as needed
	}

	// Construct the final map containing data from both HomeChefEats and Profile
	data := map[string]interface{}{
		"homeChefEats": chefData,
		"profile":      profileData,
	}

	return data, nil
}

func (repo *HomeChefEatsRepositoryDb) GetAll() ([]map[string]interface{}, error) {
	// Retrieve all HomeChefEats objects with preloaded Profile
	var chefs []*domain.HomeChefEats
	if err := repo.Db.Preload("Profile").Find(&chefs).Error; err != nil {
		return nil, err
	}

	// Map HomeChefEats objects to maps
	var chefsData []map[string]interface{}
	for _, chef := range chefs {
		chefData := map[string]interface{}{
			"id":                chef.ID,
			"homeChefEats_name": chef.HomeChefEats_name,
			// Add other HomeChefEats fields as needed
		}

		profileData := map[string]interface{}{
			"id":         chef.Profile.ID,
			"first_name": chef.Profile.First_name,
			"last_name":  chef.Profile.Last_name,
			// Add other Profile fields as needed
		}

		chefData["profile"] = profileData

		chefsData = append(chefsData, chefData)
	}

	return chefsData, nil
}

func (repo *HomeChefEatsRepositoryDb) Delete(id string) error {
	// Find the HomeChefEats entity by ID
	var chef domain.HomeChefEats
	if err := repo.Db.First(&chef, "id = ?", id).Error; err != nil {
		return err
	}

	// Delete the HomeChefEats entity
	if err := repo.Db.Delete(&chef).Error; err != nil {
		return err
	}

	return nil
}

func (repo *HomeChefEatsRepositoryDb) GetOnex() ([]map[string]interface{}, error) {
	// Retrieve HomeChefEats objects with preloaded Profile
	var HomeChefEats []*domain.HomeChefEats
	if err := repo.Db.Preload("Profile").Find(&HomeChefEats).Error; err != nil {
		return nil, err
	}

	// Map HomeChefEats objects to maps
	var restaurantsHolder []map[string]interface{}
	for _, chefEvaEat := range HomeChefEats {
		restaurantData := map[string]interface{}{
			"id":                 chefEvaEat.ID,
			"homeChefEats_image": chefEvaEat.HomeChefEats_image,
			"homeChefEats_name":  chefEvaEat.HomeChefEats_name,
			"phone_number":       chefEvaEat.Phone_number,
			"address":            chefEvaEat.Address,
			"location":           chefEvaEat.LocationID,
			"profile": map[string]interface{}{ // Embed profile data
				"user_id":       chefEvaEat.Profile.UserID,
				"first_name":    chefEvaEat.Profile.First_name,
				"last_name":     chefEvaEat.Profile.Last_name,
				"profile_image": chefEvaEat.Profile.Profile_image,
				"user_type":     chefEvaEat.Profile.User_type,
				"pro_type":      chefEvaEat.Profile.Pro_type,
				"user_name":     chefEvaEat.Profile.User_name,
			},
			"registration_status": chefEvaEat.Registration_status,
		}
		restaurantsHolder = append(restaurantsHolder, restaurantData)
	}

	return restaurantsHolder, nil
}
