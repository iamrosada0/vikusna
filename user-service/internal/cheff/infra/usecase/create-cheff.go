package usecase

import "evaeats/user-service/internal/cheff/entity"

type CreateCheffInputDto struct {
	Name               string `json:"name"`
	CheffImage         string `json:"cheff_image"`
	PhoneNumber        string `json:"phone_number"`
	Address            string `json:"address"`
	LocationID         string `json:"location"`
	RegistrationStatus string `json:"registration_status"`
}

type CreateCheffOutputDto struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	CheffImage         string `json:"cheff_image"`
	PhoneNumber        string `json:"phone_number"`
	Address            string `json:"address"`
	LocationID         string `json:"location"`
	RegistrationStatus string `json:"registration_status"`
}

type CreateCheffUseCase struct {
	CheffRepository entity.CheffRepository
}

func NewCreateCheffUseCase(cheffRepository entity.CheffRepository) *CreateCheffUseCase {
	return &CreateCheffUseCase{CheffRepository: cheffRepository}
}

func (u *CreateCheffUseCase) Execute(input CreateCheffInputDto) (*CreateCheffOutputDto, error) {
	// Criar um novo Cheff usando os dados de entrada
	Cheff := entity.NewCheff(
		input.CheffImage,
		input.Name,
		input.PhoneNumber,
		input.Address,
		input.LocationID,
	)

	// Adicionar o Cheff ao repositório
	err := u.CheffRepository.Create(Cheff)
	if err != nil {
		return nil, err
	}

	// Retornar os dados do Cheff criado na saída
	return &CreateCheffOutputDto{
		ID:                 Cheff.ID,
		Name:               Cheff.Cheff_name,
		CheffImage:         Cheff.Cheff_image,
		PhoneNumber:        Cheff.Phone_number,
		Address:            Cheff.Address,
		LocationID:         Cheff.LocationID,
		RegistrationStatus: Cheff.Registration_status,
	}, nil
}
