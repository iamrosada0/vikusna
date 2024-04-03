package main

import (
	"database/sql"
	"evaeats/user-service/api"
	cheffEntity "evaeats/user-service/internal/cheff/entity"
	dishEntity "evaeats/user-service/internal/dish/entity"
	dishcategoryEntity "evaeats/user-service/internal/dishcategory/entity"
	paymentEntity "evaeats/user-service/internal/payment/entity"
	userEntity "evaeats/user-service/internal/user/entity"
	"fmt"
	"net/http"

	cheffRepo "evaeats/user-service/internal/cheff/infra/repository"
	cheffUsecase "evaeats/user-service/internal/cheff/infra/usecase"

	dishRepo "evaeats/user-service/internal/dish/infra/repository"
	dishUsecase "evaeats/user-service/internal/dish/infra/usecase"

	dishcategoryRepo "evaeats/user-service/internal/dishcategory/infra/repository"
	dishcategoryUsecase "evaeats/user-service/internal/dishcategory/infra/usecase"

	paymentRepo "evaeats/user-service/internal/payment/infra/repository"
	paymentUsecase "evaeats/user-service/internal/payment/infra/usecase"

	userRepo "evaeats/user-service/internal/user/infra/repository"
	userUsecase "evaeats/user-service/internal/user/infra/usecase"

	notificationRepo "evaeats/user-service/internal/notification/infra/repository"
	notificationUsecase "evaeats/user-service/internal/notification/infra/usecase"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	dbPath := "./user-service/db/main.db"
	sqlDB, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	_, err = os.Stat(dbPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll("./user-service/db", os.ModePerm)
		if err != nil {
			panic(err)
		}

		file, err := os.Create(dbPath)
		if err != nil {
			panic(err)
		}
		file.Close()
	}

	// Create Gorm connection
	gormDB, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Auto-migrate necessary tables
	err = gormDB.AutoMigrate(
		&cheffEntity.Cheff{},
		&dishEntity.Dish{},
		&dishcategoryEntity.DishCategory{},
		&paymentEntity.Payment{},
		&userEntity.User{},
	)
	if err != nil {
		panic(err)
	}

	userRepo := userRepo.NewUserRepositoryPostgres(gormDB)
	cheffRepo := cheffRepo.NewCheffRepositoryPostgres(gormDB)
	dishRepo := dishRepo.NewDishRepositoryPostgres(gormDB)
	dishCategoryRepo := dishcategoryRepo.NewDishCategoryRepositoryPostgres(gormDB)
	paymentRepo := paymentRepo.NewPaymentRepositoryPostgres(gormDB)
	notificationRepo := notificationRepo.NewNotificationRepositoryPostgres(gormDB)

	// Create use cases

	createUserUC := userUsecase.NewCreateUserUseCase(userRepo)
	deleteUserUC := userUsecase.NewDeleteUserUseCase(userRepo)
	getUserByIDUC := userUsecase.NewGetUserByIDUseCase(userRepo)
	updateUserUC := userUsecase.NewUpdateUserUseCase(userRepo)
	getAllUsersUC := userUsecase.NewGetAllUsersUseCase(userRepo)

	createCheffUC := cheffUsecase.NewCreateCheffUseCase(cheffRepo)
	deleteCheffUC := cheffUsecase.NewDeleteCheffUseCase(cheffRepo)
	getCheffByIDUC := cheffUsecase.NewGetCheffByIDUseCase(cheffRepo)
	updateCheffUC := cheffUsecase.NewUpdateCheffUseCase(cheffRepo)
	getAllCheffsUC := cheffUsecase.NewGetAllCheffsUseCase(cheffRepo)

	createDishUC := dishUsecase.NewCreateDishUseCase(dishRepo)
	deleteDishUC := dishUsecase.NewDeleteDishUseCase(dishRepo)
	getDishByIDUC := dishUsecase.NewGetDishByIDUseCase(dishRepo)
	updateDishUC := dishUsecase.NewUpdateDishUseCase(dishRepo)
	getAllDishesUC := dishUsecase.NewGetAllDishsUseCase(dishRepo)

	createDishCategoryUC := dishcategoryUsecase.NewCreateDishCategoryUseCase(dishCategoryRepo)
	deleteDishCategoryUC := dishcategoryUsecase.NewDeleteDishCategoryUseCase(dishCategoryRepo)
	getDishCategoryByIDUC := dishcategoryUsecase.NewGetDishCategoryByIDUseCase(dishCategoryRepo)
	updateDishCategoryUC := dishcategoryUsecase.NewUpdateDishCategoryUseCase(dishCategoryRepo)
	getAllDishCategoriesUC := dishcategoryUsecase.NewGetAllDishCategoriesUseCase(dishCategoryRepo)

	createPaymentUC := paymentUsecase.NewCreatePaymentUseCase(paymentRepo)
	deletePaymentUC := paymentUsecase.NewDeletePaymentUseCase(paymentRepo)
	getPaymentByIDUC := paymentUsecase.NewGetPaymentByIDUseCase(paymentRepo)
	updatePaymentUC := paymentUsecase.NewUpdatePaymentUseCase(paymentRepo)
	getAllPaymentsUC := paymentUsecase.NewGetAllPaymentsUseCase(paymentRepo)

	createNotificationUC := notificationUsecase.NewCreateNotificationUseCase(notificationRepo)
	deleteNotificationUC := notificationUsecase.NewDeleteNotificationUseCase(notificationRepo)
	getNotificationByIDUC := notificationUsecase.NewGetNotificationsUseCase(notificationRepo)
	updateNotificationUC := notificationUsecase.NewUpdateNotificationUseCase(notificationRepo)

	// Create handlers
	userHandlers := api.NewUserHandlers(createUserUC, getAllUsersUC, deleteUserUC, getUserByIDUC, updateUserUC)
	cheffHandlers := api.NewCheffHandlers(createCheffUC, getAllCheffsUC, deleteCheffUC, getCheffByIDUC, updateCheffUC)
	dishHandlers := api.NewDishHandlers(createDishUC, getAllDishesUC, deleteDishUC, getDishByIDUC, updateDishUC)
	dishCategoryHandlers := api.NewDishCategoryHandlers(createDishCategoryUC, getAllDishCategoriesUC, deleteDishCategoryUC, getDishCategoryByIDUC, updateDishCategoryUC)
	paymentHandlers := api.NewPaymentHandlers(createPaymentUC, getAllPaymentsUC, getPaymentByIDUC, updatePaymentUC, deletePaymentUC)

	notificationHandlers := api.NewNotificationHandlers(createNotificationUC, getNotificationByIDUC, deleteNotificationUC, updateNotificationUC)

	// Set up Gin router
	router := gin.Default()

	// Set up routes
	userHandlers.SetupRoutes(router)
	cheffHandlers.SetupRoutes(router)
	dishHandlers.SetupRoutes(router)
	dishCategoryHandlers.SetupRoutes(router)
	paymentHandlers.SetupRoutes(router)

	notificationHandlers.SetupRoutes(router)

	// Start the server
	err = http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Println(err)
	}
}
