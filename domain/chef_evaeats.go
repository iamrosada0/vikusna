package domain

// ChefEvaEats is like a restaurant
type ChefEvaEats struct {
	ID                  string `json:"chef_evaeats_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	ChefEvaEats_image   string `json:"chef_evaeats_image"`
	ChefEvaEats_name    string `json:"chef_evaeats_name"`
	Phone_number        string `json:"phone_number"`
	Address             string `json:"address"`
	LocationID          string `json:"location"` // geo coordnates, change this interface
	ProfileID           string `json:"owner"`
	Registration_status string `json:"registration_status" validate:"eq=PENDING|eq=ACCEPTED|eq=REJECTED"`
}
