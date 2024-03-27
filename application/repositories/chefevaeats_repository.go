package repositories

import (
	"evaeats/domain"
	"fmt"

	"gorm.io/gorm"
)

type ChefEvaEatsRepository interface {
	Insert(chefEvaEatsImage, chefEvaEatsName, phoneNumber, address, locationID, profileID, registrationStatus string) (*domain.ChefEvaEats, error)
	Find(id string) (*domain.ChefEvaEats, error)
	Update(id string, chefEvaEats *domain.ChefEvaEats) (*domain.ChefEvaEats, error)
	GetOne() ([]*domain.ChefEvaEats, error)
	GetByID(id string) (map[string]interface{}, error)
	Delete(id string) error
}

type ChefEvaEatsRepositoryDb struct {
	Db *gorm.DB
}

func NewChefEvaEatsRepository(Db *gorm.DB) *ChefEvaEatsRepositoryDb {
	return &ChefEvaEatsRepositoryDb{Db: Db}
}

func (repo *ChefEvaEatsRepositoryDb) Insert(chefEvaEatsImage, chefEvaEatsName, phoneNumber, address, locationID, profileID, registrationStatus string) (*domain.ChefEvaEats, error) {
	var dbChef domain.ChefEvaEats
	repo.Db.Where("chef_evaeats_name = ?", chefEvaEatsName).First(&dbChef)
	if dbChef.ID != "" {
		return nil, fmt.Errorf("ChefEvaEats with the name %s already exists", chefEvaEatsName)
	}

	dbChef.Registration_status = "PENDING"
	if err := repo.Db.Create(&dbChef).Error; err != nil {
		return nil, err
	}

	chef := &domain.ChefEvaEats{
		ChefEvaEats_image:   chefEvaEatsImage,
		ChefEvaEats_name:    chefEvaEatsName,
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

func (repo *ChefEvaEatsRepositoryDb) Find(id string) (*domain.ChefEvaEats, error) {
	var chef domain.ChefEvaEats
	if err := repo.Db.First(&chef, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &chef, nil
}

func (repo *ChefEvaEatsRepositoryDb) Update(id string, chefEvaEats *domain.ChefEvaEats) (*domain.ChefEvaEats, error) {
	var dbChefEvaEats domain.ChefEvaEats

	if err := repo.Db.First(&dbChefEvaEats, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if err := repo.Db.Model(&dbChefEvaEats).Updates(chefEvaEats).Error; err != nil {
		return nil, err
	}

	return chefEvaEats, nil
}

func (repo *ChefEvaEatsRepositoryDb) GetOne() ([]map[string]interface{}, error) {
	// Retrieve ChefEvaEats objects with preloaded Profile
	chefEvaEats, err := repo.getChefEvaEatsWithProfile()
	if err != nil {
		return nil, err
	}

	// Map ChefEvaEats objects to maps
	restaurantsHolder := repo.mapChefEvaEatsToMaps(chefEvaEats)

	return restaurantsHolder, nil
}

// getChefEvaEatsWithProfile retrieves ChefEvaEats objects with preloaded Profile
func (repo *ChefEvaEatsRepositoryDb) getChefEvaEatsWithProfile() ([]*domain.ChefEvaEats, error) {
	var chefEvaEats []*domain.ChefEvaEats
	if err := repo.Db.Preload("Profile").Find(&chefEvaEats).Error; err != nil {
		return nil, err
	}
	return chefEvaEats, nil
}

// mapChefEvaEatsToMaps maps ChefEvaEats objects to maps
func (repo *ChefEvaEatsRepositoryDb) mapChefEvaEatsToMaps(chefEvaEats []*domain.ChefEvaEats) []map[string]interface{} {
	var restaurantsHolder []map[string]interface{}

	for _, chefEvaEat := range chefEvaEats {
		chefEvaEatsData := map[string]interface{}{
			"id":                 chefEvaEat.ID,
			"chef_evaeats_image": chefEvaEat.ChefEvaEats_image,
			"chef_evaeats_name":  chefEvaEat.ChefEvaEats_name,
			"phone_number":       chefEvaEat.Phone_number,
			"address":            chefEvaEat.Address,
			"location":           chefEvaEat.LocationID,
			"owner":              nil, // Placeholder for owner profile
		}

		// Check if ChefEvaEats has associated Profile
		if chefEvaEat.Profile.ID != "" {
			profileData := map[string]interface{}{
				"id": chefEvaEat.Profile.UserID,
				// Include other profile fields as needed
			}
			chefEvaEatsData["owner"] = profileData
		}

		// Append ChefEvaEatsData to restaurantsHolder
		restaurantsHolder = append(restaurantsHolder, chefEvaEatsData)
	}

	return restaurantsHolder
}

func (repo *ChefEvaEatsRepositoryDb) GetByID(id string) (map[string]interface{}, error) {
	var chef domain.ChefEvaEats
	if err := repo.Db.Preload("Profile").First(&chef, "id = ?", id).Error; err != nil {
		return nil, err
	}

	// Convert ChefEvaEats to map
	chefData := map[string]interface{}{
		"id":                chef.ID,
		"chef_evaeats_name": chef.ChefEvaEats_name,
		// Add other ChefEvaEats fields as needed
	}

	// Convert Profile to map
	profileData := map[string]interface{}{
		"id":         chef.Profile.ID,
		"first_name": chef.Profile.First_name,
		"last_name":  chef.Profile.Last_name,
		// Add other Profile fields as needed
	}

	// Construct the final map containing data from both ChefEvaEats and Profile
	data := map[string]interface{}{
		"chef_evaeats": chefData,
		"profile":      profileData,
	}

	return data, nil
}

func (repo *ChefEvaEatsRepositoryDb) GetAll() ([]map[string]interface{}, error) {
	// Retrieve all ChefEvaEats objects with preloaded Profile
	var chefs []*domain.ChefEvaEats
	if err := repo.Db.Preload("Profile").Find(&chefs).Error; err != nil {
		return nil, err
	}

	// Map ChefEvaEats objects to maps
	var chefsData []map[string]interface{}
	for _, chef := range chefs {
		chefData := map[string]interface{}{
			"id":                chef.ID,
			"chef_evaeats_name": chef.ChefEvaEats_name,
			// Add other ChefEvaEats fields as needed
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

func (repo *ChefEvaEatsRepositoryDb) Delete(id string) error {
	// Find the ChefEvaEats entity by ID
	var chef domain.ChefEvaEats
	if err := repo.Db.First(&chef, "id = ?", id).Error; err != nil {
		return err
	}

	// Delete the ChefEvaEats entity
	if err := repo.Db.Delete(&chef).Error; err != nil {
		return err
	}

	return nil
}
