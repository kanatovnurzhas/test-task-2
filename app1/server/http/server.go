package http

import (
	"github.com/gofiber/fiber/v2"
	"test-task-2/app1/handler"
)

func RunServer() error {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Post("/api/students/create", handler.CreateStudent)
	app.Delete("api/students/delete/:id", handler.DeleteStudent)
	app.Patch("api/students/update/:id", handler.UpdateStudent)
	app.Get("api/students/get/:id", handler.GetStudent)
	app.Get("api/students/getAll", handler.GetAllStudent)
	return app.Listen(":8888")
}
