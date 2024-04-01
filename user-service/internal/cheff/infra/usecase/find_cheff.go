package usecase

import "evaeats/user-service/internal/cheff/entity"

type GetAllCheffsOutputDto struct {
	Cheffs []*entity.Cheff `json:"cheffs"`
}

type GetAllCheffsUseCase struct {
	CheffRepository entity.CheffRepository
}

func NewGetAllCheffsUseCase(cheffRepository entity.CheffRepository) *GetAllCheffsUseCase {
	return &GetAllCheffsUseCase{CheffRepository: cheffRepository}
}

func (u *GetAllCheffsUseCase) Execute() (*GetAllCheffsOutputDto, error) {
	Cheffs, err := u.CheffRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return &GetAllCheffsOutputDto{Cheffs: Cheffs}, nil
}
