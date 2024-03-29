package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func GenerateNewID() uint {
	return uint(time.Now().UnixNano())
}

type Promotion struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Type        int    `json:"type"`
}

func FetchPromoTypeFromMicroservice(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Failed to fetch promotion, status: %d", resp.StatusCode)
	}

	var response struct {
		Type int `json:"type"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return 0, err
	}

	fmt.Println("Return of function fetchPromoTypeFromMicroservice:", response.Type)

	return response.Type, nil
}

// func ContainsCheffPromotion(CheffsPromotion []entity.CheffPromotion, CheffID int) bool {
// 	for _, CheffPromotion := range CheffsPromotion {
// 		if int(CheffPromotion.CheffID) == CheffID {
// 			return true
// 		}
// 	}
// 	return false
// }
