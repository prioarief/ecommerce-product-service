package main

import (
	"fmt"
	"strconv"

	"github.com/prioarief/ecommerce-product-service/configs"
)

func main() {
	viperConfig := configs.NewViper()
	app := configs.NewFiber(viperConfig)
	db := configs.NewDatabase(viperConfig)
	validate := configs.NewValidator(viperConfig)
	log := configs.NewLogger(viperConfig)
	// db := configs.GormConfiguration(viperConfig, log)

	port, err := strconv.Atoi(viperConfig.GetString("APP_PORT"))
	if err != nil {
		log.WithError(err).Error("cannot get APP PORT")
	}

	configs.Bootstrap(&configs.BootstrapConfig{
		App:      app,
		DB:       db,
		Config:   viperConfig,
		Validate: validate,
		Log:      log,
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
