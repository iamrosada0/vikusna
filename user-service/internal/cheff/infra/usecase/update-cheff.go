package usecase

import "evaeats/user-service/internal/cheff/entity"

type UpdateCheffInputDto struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateCheffOutputDto struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateCheffUseCase struct {
	CheffRepository entity.CheffRepository
}

func NewUpdateCheffUseCase(cheffRepository entity.CheffRepository) *UpdateCheffUseCase {
	return &UpdateCheffUseCase{CheffRepository: cheffRepository}
}

func (u *UpdateCheffUseCase) Execute(input UpdateCheffInputDto) (*UpdateCheffOutputDto, error) {
	existingCheff, err := u.CheffRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	existingCheff.Update(
		input.Name,
		input.Email,
	)

	err = u.CheffRepository.Update(existingCheff)
	if err != nil {
		return nil, err
	}

	return &UpdateCheffOutputDto{
		ID:    existingCheff.ID,
		Name:  existingCheff.Name,
		Email: existingCheff.Email,
	}, nil
}
