package routes

import(
	"backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router){
    users := router.Group("/users")	
	users.Get("/", controllers.GetAllUsers)
	users.Get("/:id", controllers.GetUserById)
	users.Post("/", controllers.CreateUser)
	users.Put("/:id", controllers.UpdateUser)
	users.Delete("/:id", controllers.DeleteUser)
}