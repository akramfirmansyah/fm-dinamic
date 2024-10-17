package controllers

import "github.com/gofiber/fiber/v2"

func CreatePlayer(c *fiber.Ctx) error {
	return c.SendString("Hi...")
}

func GetAllPlayers(c *fiber.Ctx) error {
	return c.SendString("Get all players")
}

func GetPlayerByUUID(c *fiber.Ctx) error {
	return c.SendString("Get by UUID")
}

func UpdatePlayer(c *fiber.Ctx) error {
	return c.SendString("Update player")
}

func DeletePlayer(c *fiber.Ctx) error {
	return c.SendString("Delete player")
}
