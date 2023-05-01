package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (ch *CourseHandler) GetByStudent(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	fmt.Println(name[1:])
	courses, err := ch.service.GetByStudent(name[1:])
	if err != nil {
		wrappedErr := fmt.Errorf("error is: %w", err)
		fmt.Println(wrappedErr)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"status":  fiber.StatusInternalServerError,
		})
	}
	fmt.Println("courses:", courses)
	return ctx.JSON(fiber.Map{
		"courses": courses,
	})
}
