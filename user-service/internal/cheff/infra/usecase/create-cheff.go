package usecase

import "evaeats/user-service/internal/cheff/entity"

type CreateCheffInputDto struct {
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

type CreateCheffOutputDto struct {
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

type CreateCheffUseCase struct {
	CheffRepository entity.CheffRepository
}

func NewCreateCheffUseCase(cheffRepository entity.CheffRepository) *CreateCheffUseCase {
	return &CreateCheffUseCase{CheffRepository: cheffRepository}
}

func (u *CreateCheffUseCase) Execute(input CreateCheffInputDto) (*CreateCheffOutputDto, error) {
	// Create a new chef using the input data
	cheff := entity.NewCheff(
		input.CheffImage,
		input.CheffName,
		input.PhoneNumber,
		input.Address,
		input.LocationID,
		input.LocationPreference,
		input.CookingExperience,
		input.Specialties,
		input.Certifications,
	)

	// Add the chef to the repository
	err := u.CheffRepository.Create(cheff)
	if err != nil {
		return nil, err
	}

	// Return the data of the created chef in the output
	return &CreateCheffOutputDto{
		ID:                 cheff.ID,
		CheffName:          cheff.Cheff_name,
		CheffImage:         cheff.Cheff_image,
		PhoneNumber:        cheff.Phone_number,
		Address:            cheff.Address,
		LocationID:         cheff.LocationID,
		RegistrationStatus: cheff.Registration_status,
		LocationPreference: cheff.LocationPreference,
		CookingExperience:  cheff.CookingExperience,
		Specialties:        cheff.Specialties,
		Certifications:     cheff.Certifications,
	}, nil
}
