package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (st *StudentHandler) GetAllStudent(ctx *fiber.Ctx) error {
	students, err := st.service.GetAll()
	if err != nil {
		wrappedErr := fmt.Errorf("error is: %w", err)
		fmt.Println(wrappedErr)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"status":  fiber.StatusInternalServerError,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":  true,
		"students": students,
	})
}
