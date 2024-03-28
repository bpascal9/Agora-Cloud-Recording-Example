package main

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/raysandeep/Estimator-App/api/router"
	"github.com/raysandeep/Estimator-App/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func healthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func main() {
	// Set global configuration
	viper.SetConfigName("config.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Panicln(fmt.Errorf("fatal error config file: %s", err))
	}
	viper.AutomaticEnv()

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", healthCheck)
	api.MountRoutes(app)

	app.Listen(":" + viper.GetString("PORT"))
}
