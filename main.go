package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/zenklot/backend-zenknote/database"
	"github.com/zenklot/backend-zenknote/helper"
	"github.com/zenklot/backend-zenknote/router"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	database.ConnectDB()
	var port string
	port = helper.Config("PORT")
	if port != "" {
		port = ":" + port
	} else {
		port = ":3000"
	}
	router.SetupRoutes(app)
	log.Fatal(app.Listen(port))
	// defer database.DB.Close()
}
