package service

import (
	"fmt"
	"global/pkg/kafka"

	"github.com/kanatovnurzhas/test-task-2/course-microservice/internal/models"
	"github.com/kanatovnurzhas/test-task-2/course-microservice/internal/repository"
)

const path = "service"

type ICourseService interface {
	CreateCourse(course models.Course) error
	GetAll() ([]models.Course, error)
	GetCourseByID(id int) (models.Course, error)
	UpdateCourse(course models.Course) error
	DeleteCourse(id int) error
	GetByStudent(name string) ([]models.Course, error)
	ProduceKafka(name string) error
}

type courseService struct {
	CourseRepo  repository.ICourseRepo
	kafkaClient kafka.Messaging
}

func CourseServiceInit(repo repository.ICourseRepo, kafka kafka.Messaging) ICourseService {
	return &courseService{
		CourseRepo:  repo,
		kafkaClient: kafka,
	}
}

func (cr *courseService) ProduceKafka(name string) error {
	topic := "course_to_student"
	key := []byte("course-name")
	msg := []byte(name)
	err := cr.kafkaClient.Write(topic, key, msg)
	if err != nil {
		return fmt.Errorf(path+"produce cafka: %w", err)
	}
	fmt.Println("Message sent successfully")
	return nil
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
