package usecase

import (
	"evaeats/user-service/internal/dish/entity"
)

type CreateDishCategoryUseCase struct {
	DishCategoryRepository entity.DishCategoryRepository
}

type CreateDishCategoryInputDto struct {
	Names []string `json:"names"`
}

func NewCreateDishCategoryUseCase(DishCategoryRepository entity.DishCategoryRepository) *CreateDishCategoryUseCase {
	return &CreateDishCategoryUseCase{DishCategoryRepository: DishCategoryRepository}
}

func (u *CreateDishCategoryUseCase) Execute(input CreateDishCategoryInputDto) ([]*CreateDishCategoryOutputDto, error) {
	var outputs []*CreateDishCategoryOutputDto

	for _, name := range input.Names {
		newCategory, err := entity.NewDishCategory(name)
		if err != nil {
			return nil, err
		}
		err = u.DishCategoryRepository.Create(newCategory)
		if err != nil {
			return nil, err
		}
		output := &CreateDishCategoryOutputDto{
			ID:   newCategory.ID,
			Name: newCategory.Name,
		}

		outputs = append(outputs, output)
	}

	return outputs, nil
}

type DeleteDishCategoryUseCase struct {
	DishCategoryRepository entity.DishCategoryRepository
}

func NewDeleteDishCategoryUseCase(DishCategoryRepository entity.DishCategoryRepository) *DeleteDishCategoryUseCase {
	return &DeleteDishCategoryUseCase{DishCategoryRepository: DishCategoryRepository}
}

func (u *DeleteDishCategoryUseCase) Execute(id string) error {
	// Call DishCategoryRepository to delete the DishCategory by ID
	err := u.DishCategoryRepository.DeleteByID(id)
	if err != nil {
		return err
	}

	return nil
}

type UpdateDishCategoryUseCase struct {
	DishCategoryRepository entity.DishCategoryRepository
}

func NewUpdateDishCategoryUseCase(DishCategoryRepository entity.DishCategoryRepository) *UpdateDishCategoryUseCase {
	return &UpdateDishCategoryUseCase{DishCategoryRepository: DishCategoryRepository}
}

func (u *UpdateDishCategoryUseCase) Execute(input UpdateDishCategoryInputDto) (*UpdateDishCategoryOutputDto, error) {
	// Retrieve the existing DishCategory
	existingCategory, err := u.DishCategoryRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	// Update the existing DishCategory entity
	existingCategory.Name = input.Name

	// Call DishCategoryRepository to update the DishCategory
	err = u.DishCategoryRepository.Update(existingCategory)
	if err != nil {
		return nil, err
	}

	// Construct output DTO using updated DishCategory
	output := &UpdateDishCategoryOutputDto{
		ID:   existingCategory.ID,
		Name: existingCategory.Name,
	}

	return output, nil
}

type GetDishCategoryByIDUseCase struct {
	DishCategoryRepository entity.DishCategoryRepository
}

func NewGetDishCategoryByIDUseCase(DishCategoryRepository entity.DishCategoryRepository) *GetDishCategoryByIDUseCase {
	return &GetDishCategoryByIDUseCase{DishCategoryRepository: DishCategoryRepository}
}

func (u *GetDishCategoryByIDUseCase) Execute(id string) (*GetDishCategoryByIDOutputDto, error) {
	// Retrieve the DishCategory by ID
	category, err := u.DishCategoryRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Construct output DTO using retrieved DishCategory
	output := &GetDishCategoryByIDOutputDto{
		ID:   category.ID,
		Name: category.Name,
	}

	return output, nil
}

type GetAllDishCategoriesUseCase struct {
	DishCategoryRepository entity.DishCategoryRepository
}

func NewGetAllDishCategoriesUseCase(DishCategoryRepository entity.DishCategoryRepository) *GetAllDishCategoriesUseCase {
	return &GetAllDishCategoriesUseCase{DishCategoryRepository: DishCategoryRepository}
}

func (u *GetAllDishCategoriesUseCase) Execute() ([]*GetAllDishCategoriesOutputDto, error) {
	// Retrieve all DishCategories
	categories, err := u.DishCategoryRepository.FindAll()
	if err != nil {
		return nil, err
	}

	// Construct output DTOs using retrieved DishCategories
	var output []*GetAllDishCategoriesOutputDto
	for _, category := range categories {
		output = append(output, &GetAllDishCategoriesOutputDto{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return output, nil
}

// Define input and output DTOs for the use cases

type CreateDishCategoryOutputDto struct {
	ID   string `json:"category_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Name string `json:"name"`
}

type UpdateDishCategoryInputDto struct {
	ID   string `json:"category_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Name string `json:"name"`
}

type UpdateDishCategoryOutputDto struct {
	ID   string `json:"category_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Name string `json:"name"`
}

type GetDishCategoryByIDOutputDto struct {
	ID   string `json:"category_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Name string `json:"name"`
}

type GetAllDishCategoriesOutputDto struct {
	ID   string `json:"category_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Name string `json:"name"`
}
