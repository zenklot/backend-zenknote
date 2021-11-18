package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/zenklot/backend-zenknote/database"
	"github.com/zenklot/backend-zenknote/router"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	database.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
	// defer database.DB.Close()
}
