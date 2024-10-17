package routers

import (
	"akramfirmansyah/fm-dinamic-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterPlayerRoutes(router fiber.Router) {
	router.Post("/", controllers.CreatePlayer)
	router.Get("/", controllers.GetAllPlayers)
	router.Get("/:uuid", controllers.GetPlayerByUUID)
	router.Put("/:uuid", controllers.UpdatePlayer)
	router.Delete("/:uuid", controllers.DeletePlayer)
}
