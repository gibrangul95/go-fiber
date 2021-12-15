package main

import (
	"github.com/gibrangul95/go-fiber/router"
	"github.com/gibrangul95/go-fiber/database"
	"github.com/gofiber/fiber/v2"
	"encoding/json"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	database.ConnectDB()

	router.SetupRoutes(app)

	app.Listen(":3000")
}