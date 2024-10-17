package main

import (
	"akramfirmansyah/fm-dinamic-api/controllers"
	"akramfirmansyah/fm-dinamic-api/routers"
	"akramfirmansyah/fm-dinamic-api/utils"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	utils.LoadEnv()
	ConnectDatabase()
	Migrate()
}

func main() {
	app := fiber.New(fiber.Config{
		AppName:     "FM Dinamic",
		JSONEncoder: jsoniter.Marshal,
		JSONDecoder: jsoniter.Unmarshal,
	})

	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to FM Dinamic!")
	})

	app.Post("/login", controllers.Login)

	api := app.Group("/api")

	routers.RegisterPlayerRoutes(api.Group("/player"))

	app.Listen(":" + os.Getenv("PORT"))
}
