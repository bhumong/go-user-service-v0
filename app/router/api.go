package router

import (
	"github.com/bhumong/go-user-service-v0/app/handler"
	"github.com/bhumong/go-user-service-v0/app/middleware"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupApi(app *fiber.App) {

	api := app.Group("/api", middleware.ApiKey)
	api.Post("/login", handler.Login)
	userApi := api.Group("/", middleware.ApiKey, middleware.AuthReq)

	// must login
	userApi.Get("/users/:id", handler.GetUser)
	userApi.Delete("/users/:id", handler.DeleteUser)
	userApi.Patch("/users/:id", handler.UpdateUser)
	userApi.Get("/users", handler.GetUsers)
	userApi.Post("/users", handler.CreateUser)
}
