package search

import (
	"awesome/models"
	"awesome/utils"

	"github.com/gofiber/fiber/v2"
)

func HandleSearch(ctx *fiber.Ctx) error {
	searchBody := struct {
		Text string `json:"text" validate:"required"`
	}{}
	if err := ctx.BodyParser(&searchBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Could Not Parse Body")
	}

	var resources []models.Resource
	db, err := utils.GetDbFromRequestOrigin(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Could Not Find Resources")
	}
	err = db.Where("name ILIKE ?", "%"+searchBody.Text+"%").Or("description ILIKE ?", "%"+searchBody.Text+"%").Find(&resources).Error
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Could Not Find Resources")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"resources": resources,
	})
}
