package usecase

import "evaeats/user-service/internal/cheff/entity"

type GetCheffByIDInputDto struct {
	ID uint `json:"id"`
}

type GetCheffByIDOutputDto struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GetCheffByIDUseCase struct {
	CheffRepository entity.CheffRepository
}

func NewGetCheffByIDUseCase(cheffRepository entity.CheffRepository) *GetCheffByIDUseCase {
	return &GetCheffByIDUseCase{CheffRepository: cheffRepository}
}

func (u *GetCheffByIDUseCase) Execute(input GetCheffByIDInputDto) (*GetCheffByIDOutputDto, error) {
	Cheff, err := u.CheffRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &GetCheffByIDOutputDto{
		ID:    Cheff.ID,
		Name:  Cheff.Name,
		Email: Cheff.Email,
	}, nil
}
