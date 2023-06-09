package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kanatovnurzhas/test-task-2/student-microservice/internal/models"
)

func (st *StudentHandler) CreateStudent(ctx *fiber.Ctx) error {
	student := new(models.Student)

	if err := ctx.BodyParser(&student); err != nil {
		wrappedErr := fmt.Errorf("error is: %w", err)
		fmt.Println(wrappedErr)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"status":  fiber.StatusInternalServerError,
		})
	}
	err := st.service.CreateStudent(*student)
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
