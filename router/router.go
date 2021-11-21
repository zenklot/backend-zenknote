package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zenklot/backend-zenknote/handler"
	"github.com/zenklot/backend-zenknote/helper"
	"github.com/zenklot/backend-zenknote/middleware"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

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
	note.Delete("/:id", middleware.Protected(helper.Config("KUNCI_RAHASIA")), handler.DeleteNote)
	note.Put("/:id", middleware.Protected(helper.Config("KUNCI_RAHASIA")), handler.PutNote)

	profile := api.Group("/profile")
	profile.Get("/", middleware.Protected(helper.Config("KUNCI_RAHASIA")), handler.GetProfile)
	profile.Put("/", middleware.Protected(helper.Config("KUNCI_RAHASIA")), handler.PutProfile)
}
