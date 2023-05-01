package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kanatovnurzhas/test-task-2/app2/internal/models"
)

func (ch *CourseHandler) UpdateCourse(ctx *fiber.Ctx) error {
	course := new(models.Course)

	if err := ctx.BodyParser(course); err != nil {
		wrappedErr := fmt.Errorf("error is: %w", err)
		fmt.Println(wrappedErr)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"status":  fiber.StatusInternalServerError,
		})
	}
	err := ch.service.UpdateCourse(*course)
	if err != nil {
		wrappedErr := fmt.Errorf("error is: %w", err)
		fmt.Println(wrappedErr)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"status":  fiber.StatusInternalServerError,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"status":  fiber.StatusOK,
	})
}
