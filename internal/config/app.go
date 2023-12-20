package config

import (
	"be/internal/delivery/http"
	"be/internal/delivery/http/middleware"
	"be/internal/delivery/http/route"
	"be/internal/repository"
	"be/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {

	// setup repositories
	userRepository := repository.NewUserRepository(config.Log)
	businessRepository := repository.NewBusinesseRepository(config.Log)

	// setup use cases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)
	businessUseCase := usecase.NewBusinesseUseCase(config.DB, config.Log, config.Validate, businessRepository)

	// setup controller
	userController := http.NewUserController(userUseCase, config.Log)
	businessController := http.NewBusinessController(businessUseCase, config.Log)

	// setup middleware
	authMiddleware := middleware.NewAuth(userUseCase)

	routeConfig := route.RouteConfig{
		App:                 config.App,
		UserController:      userController,
		BusinesseController: businessController,
		AuthMiddleware:      authMiddleware,
	}
	routeConfig.Setup()
}
