package main

import (
	"App/database/connection"
	"App/router"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func main() {
	err := connection.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer connection.DB.Close()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(static.New("./static"))
	router.Router(app)
	log.Fatal(app.Listen(":2345"))
}
