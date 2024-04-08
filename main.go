package main

import (
	"log"
	"meeting3/database"
	"meeting3/middleware"
	"meeting3/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()
	database.ConnectDatabase()

	app.Use(
		cors.New(),
		logger.New(),
	)

	//Kumpulan route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World")
	})

	app.Get("/hello", routes.Hello)
	app.Get("/heelo/:id", routes.HelloParams)
	app.Post("/register", routes.Register)
	app.Post("/login", routes.Login)
	app.Put("/update", middleware.JwtProtect(), routes.UpdateDataUser)

	log.Fatal(app.Listen(":3000"))

}
