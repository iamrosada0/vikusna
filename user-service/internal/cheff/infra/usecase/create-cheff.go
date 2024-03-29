package usecase

import "evaeats/user-service/internal/cheff/entity"

type CreateCheffInputDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateCheffOutputDto struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateCheffUseCase struct {
	CheffRepository entity.CheffRepository
}

func NewCreateCheffUseCase(cheffRepository entity.CheffRepository) *CreateCheffUseCase {
	return &CreateCheffUseCase{CheffRepository: cheffRepository}
}

func (u *CreateCheffUseCase) Execute(input CreateCheffInputDto) (*CreateCheffOutputDto, error) {
	Cheff := entity.NewCheff(
		input.Name,
		input.Email,
	)

	err := u.CheffRepository.Create(Cheff)
	if err != nil {
		return nil, err
	}

	return &CreateCheffOutputDto{
		ID:    Cheff.ID,
		Name:  Cheff.Name,
		Email: Cheff.Email,
	}, nil
}
