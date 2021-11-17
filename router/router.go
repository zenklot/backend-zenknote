package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zenklot/backend-zenknote/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// employee := api.Group("/employee")
	// employee.Get("/", handler.GetEmployee)
	// employee.Post("/", handler.PostEmployee)
	// employee.Put("/:id", handler.PutEmployee)
	// employee.Delete("/:id", handler.DeleteEmployee)

	auth := api.Group("/auth")
	auth.Post("/login", handler.PostLogin)
	auth.Post("/register", handler.PostRegister)
	// auth.Post("/forget-password", handler.PostForgetPassword)
	// auth.Put("/forget-password", handler.PutForgetPassword)
	// auth.Get("/logout", handler.GetLogout)
	// auth.Post("/refresh", handler.PostRefresh)
}