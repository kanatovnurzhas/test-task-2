package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (sh *StudentHandler) GetByCourse(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	fmt.Println(name[1:])
	students, err := sh.service.GetByCourse(name[1:])
	if err != nil {
		wrappedErr := fmt.Errorf("error is: %w", err)
		fmt.Println(wrappedErr)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"status":  fiber.StatusInternalServerError,
		})
	}
	return ctx.JSON(fiber.Map{
		"students": students,
	})
}
