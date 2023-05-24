package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kanatovnurzhas/test-task-2/student-microservice/internal/service"
)

type StudentHandler struct {
	service service.IStudentService
}

func StudentHandlerInit(studentService service.IStudentService) *StudentHandler {
	return &StudentHandler{
		service: studentService,
	}
}

func (sh *StudentHandler) RegisterStudentRoutes(fb fiber.Router) {
	fb.Post("/students/create", sh.CreateStudent)
	fb.Delete("/students/delete/:id", sh.DeleteStudent)
	fb.Patch("/students/update/:id", sh.UpdateStudent)
	fb.Get("/students/get/:id", sh.GetStudent)
	fb.Get("/students/getAll", sh.GetAllStudent)
	fb.Get("/students/:name/courses", sh.GetCourses)
	fb.Get("/students/getByCourse/:name", sh.GetByCourse)
}
