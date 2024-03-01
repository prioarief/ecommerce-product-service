package configs

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/prioarief/ecommerce-product-service/routers"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Config   *viper.Viper
	Validate *validator.Validate
	Log      *logrus.Logger
}

func Bootstrap(config *BootstrapConfig) {

	// setup route
	routeConfig := routers.RouteConfig{
		App: config.App,
	}

	routeConfig.Setup()
}
