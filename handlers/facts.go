package handlers

import (
	"github.com/delwar/database"
	"github.com/delwar/model"
	"github.com/gofiber/fiber/v2"
)

func CreateFact(c *fiber.Ctx) error {
	fact := new(model.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}

func ListFacts(c *fiber.Ctx) error {
	facts := []model.Fact{}
	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}