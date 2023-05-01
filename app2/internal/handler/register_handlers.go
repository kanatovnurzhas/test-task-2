package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kanatovnurzhas/test-task-2/app2/internal/service"
)

type CourseHandler struct {
	service service.ICourseService
}

func CourseHandlerInit(studentService service.ICourseService) *CourseHandler {
	return &CourseHandler{
		service: studentService,
	}
}

func (ch *CourseHandler) RegisterCourseRoutes(fb fiber.Router) {
	fb.Post("/courses/create", ch.CreateCourse)
	fb.Delete("/courses/delete/:id", ch.DeleteCourse)
	fb.Patch("/courses/update/:id", ch.UpdateCourse)
	fb.Get("/courses/get/:id", ch.GetCourse)
	fb.Get("/courses/getAll", ch.GetAllCourse)
	fb.Get("/courses/:name/students", ch.GetStudents)
	fb.Get("/courses/getByStudent/:name", ch.GetByStudent)
}
