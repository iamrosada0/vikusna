package usecase

import "evaeats/user-service/internal/cheff/entity"

type UpdateCheffInputDto struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	CheffImage         string `json:"cheff_image"`
	PhoneNumber        string `json:"phone_number"`
	Address            string `json:"address"`
	LocationID         string `json:"location"`
	RegistrationStatus string `json:"registration_status"`
}

type UpdateCheffOutputDto struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	CheffImage         string `json:"cheff_image"`
	PhoneNumber        string `json:"phone_number"`
	Address            string `json:"address"`
	LocationID         string `json:"location"`
	RegistrationStatus string `json:"registration_status"`
}

type UpdateCheffUseCase struct {
	CheffRepository entity.CheffRepository
}

func NewUpdateCheffUseCase(cheffRepository entity.CheffRepository) *UpdateCheffUseCase {
	return &UpdateCheffUseCase{CheffRepository: cheffRepository}
}

func (u *UpdateCheffUseCase) Execute(input UpdateCheffInputDto) (*UpdateCheffOutputDto, error) {
	// Busca o Cheff existente pelo ID fornecido
	existingCheff, err := u.CheffRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	// Atualiza os campos do Cheff com os novos valores fornecidos
	existingCheff.Update(
		input.Name,
		input.CheffImage,
		input.PhoneNumber,
		input.Address,
		input.LocationID,
		input.RegistrationStatus,
	)

	// Atualiza o Cheff no repositório
	err = u.CheffRepository.Update(existingCheff)
	if err != nil {
		return nil, err
	}

	// Retorna os dados atualizados do Cheff na saída
	return &UpdateCheffOutputDto{
		ID:                 existingCheff.ID,
		Name:               existingCheff.Cheff_name,
		CheffImage:         existingCheff.Cheff_image,
		PhoneNumber:        existingCheff.Phone_number,
		Address:            existingCheff.Address,
		LocationID:         existingCheff.LocationID,
		RegistrationStatus: existingCheff.Registration_status,
	}, nil
}
