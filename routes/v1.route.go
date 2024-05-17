package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-todo-app/controllers"
	"go-todo-app/middleware"
)

func v1Route(app *fiber.App) {

	user := app.Group("/users")

	user.Post("/register", controllers.UserHandlerCreate)
	user.Post("/login", controllers.LoginHandler)
	user.Get("/", middleware.AuthMiddleware, middleware.CheckAdminMiddleware, controllers.UserHandlerGetAll)
	user.Put("/", middleware.AuthMiddleware, controllers.UpdateUserById)
	user.Delete("/:id", middleware.AuthMiddleware, middleware.CheckAdminMiddleware, controllers.DeleteUserById)
}
