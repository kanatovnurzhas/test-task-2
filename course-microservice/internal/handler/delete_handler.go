package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (ch *CourseHandler) DeleteCourse(ctx *fiber.Ctx) error {
	tempID := ctx.Params("id")
	id, _ := strconv.Atoi(tempID)
	err := ch.service.DeleteCourse(id)
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
