package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (st *StudentHandler) GetCoursesKafka(ctx *fiber.Ctx) error {
	name := ctx.Params("name")

	err := st.service.ProduceKafka(name)
	if err != nil {
		wrappedErr := fmt.Errorf("error is: %w", err)
		fmt.Println(wrappedErr)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"status":  fiber.StatusInternalServerError,
		})
	}
	// return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
	// 	"success": true,
	// 	"courses": "fuck you",
	// })

	courses, err := st.service.AnswerForStud()
	fmt.Println(courses)
	if err != nil || len(courses.Course) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"courses": "fuck you",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"courses": courses,
	})
}
