package main

import (
	db "github.com/AkulinIvan/Golang/config"
	routes "github.com/AkulinIvan/Golang/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db.PostgresConnection()

	app := fiber.New()

	routes.Setup(app)

}
