package usecase

import "evaeats/user-service/internal/cheff/entity"

type DeleteCheffInputDto struct {
	ID string `json:"id"`
}

type DeleteCheffOutputDto struct {
	ID string `json:"id"`
}

type DeleteCheffUseCase struct {
	CheffRepository entity.CheffRepository
}

func NewDeleteCheffUseCase(cheffRepository entity.CheffRepository) *DeleteCheffUseCase {
	return &DeleteCheffUseCase{CheffRepository: cheffRepository}
}

func (u *DeleteCheffUseCase) Execute(input DeleteCheffInputDto) (*DeleteCheffOutputDto, error) {
	err := u.CheffRepository.DeleteByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &DeleteCheffOutputDto{ID: input.ID}, nil
}
