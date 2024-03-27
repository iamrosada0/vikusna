package domain

// HomeChefEats is like a restaurant
type HomeChefEats struct {
	ID                  string  `json:"homeChefEats_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	HomeChefEats_image  string  `json:"homeChefEats_image"`
	HomeChefEats_name   string  `json:"homeChefEats_name"`
	Phone_number        string  `json:"phone_number"`
	Address             string  `json:"address"`
	LocationID          string  `json:"location"`
	ProfileID           string  `json:"profile_id"`                          // Foreign key for Profile
	Profile             Profile `json:"profile" gorm:"foreignKey:ProfileID"` // Define the association with Profile
	Registration_status string  `json:"registration_status" validate:"eq=PENDING|eq=ACCEPTED|eq=REJECTED"`
}
