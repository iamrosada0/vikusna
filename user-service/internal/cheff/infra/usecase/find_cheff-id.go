package usecase

import "evaeats/user-service/internal/cheff/entity"

type GetCheffByIDInputDto struct {
	ID string `json:"id"`
}

type GetCheffByIDOutputDto struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	UserId string `json:"user_id" valid:"uuid" gorm:"type:uuid;primary_key"`

	CheffImage         string `json:"cheff_image"`
	PhoneNumber        string `json:"phone_number"`
	Address            string `json:"address"`
	LocationID         string `json:"location"`
	RegistrationStatus string `json:"registration_status"`
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
		ID:                 Cheff.ID,
		Name:               Cheff.Cheff_name,
		UserId:             Cheff.UserId,
		CheffImage:         Cheff.Cheff_image,
		PhoneNumber:        Cheff.Phone_number,
		Address:            Cheff.Address,
		LocationID:         Cheff.LocationID,
		RegistrationStatus: Cheff.Registration_status,
	}, nil
}
