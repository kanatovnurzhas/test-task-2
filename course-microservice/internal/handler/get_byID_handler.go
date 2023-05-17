package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (ch *CourseHandler) GetCourse(ctx *fiber.Ctx) error {
	tempID := ctx.Params("id")
	ID, _ := strconv.Atoi(tempID)
	course, err := ch.service.GetCourseByID(ID)
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
		"course":  course,
	})
}
