package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zenklot/backend-zenknote/handler"
	"github.com/zenklot/backend-zenknote/helper"
	"github.com/zenklot/backend-zenknote/middleware"
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
	auth.Post("/forget-password", handler.PostForgetPassword)
	auth.Put("/forget-password", handler.PutForgetPassword)
	// auth.Get("/logout", handler.GetLogout)
	auth.Post("/refresh", handler.PostRefresh)

	note := api.Group("/note")
	api.Get("/notes", middleware.Protected(helper.Config("KUNCI_RAHASIA")), handler.GetNotes)
	note.Post("/", middleware.Protected(helper.Config("KUNCI_RAHASIA")), handler.PostNote)
	note.Get("/:id", middleware.Protected(helper.Config("KUNCI_RAHASIA")), handler.GetNoteById)
}
