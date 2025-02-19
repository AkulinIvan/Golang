package main

import (
	db "github.com/AkulinIvan/go-test/config"
	routes "github.com/AkulinIvan/go-test/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {
	db.PostgresConnection()

	app := fiber.New()
	app.Use(app)

	routes.Setup(app)

}
