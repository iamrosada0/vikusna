package usecase

import "evaeats/user-service/internal/cheff/entity"

type UpdateCheffInputDto struct {
	ID                 string `json:"cheff_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	CheffImage         string `json:"cheff_image"`
	CheffName          string `json:"cheff_name"`
	PhoneNumber        string `json:"phone_number"`
	Address            string `json:"address"`
	LocationID         string `json:"location"`
	RegistrationStatus string `json:"registration_status" validate:"eq=PENDING|eq=ACCEPTED|eq=REJECTED"`
	LocationPreference string `json:"location_preference"`
	CookingExperience  string `json:"cooking_experience"`
	Specialties        string `json:"specialties"`
	Certifications     string `json:"certifications"`
}

type UpdateCheffOutputDto struct {
	ID                 string `json:"cheff_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	CheffImage         string `json:"cheff_image"`
	CheffName          string `json:"cheff_name"`
	PhoneNumber        string `json:"phone_number"`
	Address            string `json:"address"`
	LocationID         string `json:"location"`
	RegistrationStatus string `json:"registration_status" validate:"eq=PENDING|eq=ACCEPTED|eq=REJECTED"`
	LocationPreference string `json:"location_preference"`
	CookingExperience  string `json:"cooking_experience"`
	Specialties        string `json:"specialties"`
	Certifications     string `json:"certifications"`
}

type UpdateCheffUseCase struct {
	CheffRepository entity.CheffRepository
}

func NewUpdateCheffUseCase(cheffRepository entity.CheffRepository) *UpdateCheffUseCase {
	return &UpdateCheffUseCase{CheffRepository: cheffRepository}
}

func (u *UpdateCheffUseCase) Execute(input UpdateCheffInputDto) (*UpdateCheffOutputDto, error) {
	// Get the existing chef by the provided ID
	existingCheff, err := u.CheffRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	// Update the chef's fields with the new provided values
	existingCheff.Update(
		input.CheffImage,
		input.CheffName,
		input.PhoneNumber,
		input.Address,
		input.LocationID,
		input.RegistrationStatus,
		input.LocationPreference,
		input.CookingExperience,
		input.Specialties,
		input.Certifications,
	)

	// Update the chef in the repository
	err = u.CheffRepository.Update(existingCheff)
	if err != nil {
		return nil, err
	}

	// Return the updated chef's data in the output
	return &UpdateCheffOutputDto{
		ID:                 existingCheff.ID,
		CheffName:          existingCheff.Cheff_name,
		CheffImage:         existingCheff.Cheff_image,
		PhoneNumber:        existingCheff.Phone_number,
		Address:            existingCheff.Address,
		LocationID:         existingCheff.LocationID,
		RegistrationStatus: existingCheff.Registration_status,
		LocationPreference: existingCheff.LocationPreference,
		CookingExperience:  existingCheff.CookingExperience,
		Specialties:        existingCheff.Specialties,
		Certifications:     existingCheff.Certifications,
	}, nil
}
