package service

import (
	"fmt"

	"github.com/kanatovnurzhas/test-task-2/course-microservice/internal/models"
	"github.com/kanatovnurzhas/test-task-2/course-microservice/internal/repository"
)

type ICourseService interface {
	CreateCourse(course models.Course) error
	GetAll() ([]models.Course, error)
	GetCourseByID(id int) (models.Course, error)
	UpdateCourse(course models.Course) error
	DeleteCourse(id int) error
	GetByStudent(name string) ([]models.Course, error)
}

type courseService struct {
	CourseRepo repository.ICourseRepo
}

func CourseServiceInit(repo repository.ICourseRepo) ICourseService {
	return &courseService{
		CourseRepo: repo,
	}
}

func (cr *courseService) CreateCourse(course models.Course) error {
	// какая то бизнес логика
	return cr.CourseRepo.CreateCourse(course)
}

func (cr *courseService) GetAll() ([]models.Course, error) {
	// какая то бизнес логика
	return cr.CourseRepo.GetAll()
}

func (cr *courseService) GetCourseByID(id int) (models.Course, error) {
	// какая то бизнес логика
	return cr.CourseRepo.GetCourseByID(id)
}

func (cr *courseService) UpdateCourse(course models.Course) error {
	// какая то бизнес логика
	return cr.CourseRepo.UpdateCourse(course)
}

func (cr *courseService) DeleteCourse(id int) error {
	// какая то бизнес логика
	return cr.CourseRepo.DeleteCourse(id)
}

func (cr *courseService) GetByStudent(name string) ([]models.Course, error) {
	// какая то бизнес логика
	fmt.Println("service:" + name)
	return cr.CourseRepo.GetByStudent(name)
}
