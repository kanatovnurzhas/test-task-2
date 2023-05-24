package service

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kanatovnurzhas/test-task-2/course-microservice/internal/models"
	"github.com/kanatovnurzhas/test-task-2/course-microservice/internal/repository"
	"github.com/kanatovnurzhas/test-task-2/pkg/kafka"
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
	AnswerKafka() error
}

type courseService struct {
	CourseRepo  repository.ICourseRepo
	kafkaClient kafka.Messaging
	chName      chan []byte
	chAnswer    chan []byte
}

func CourseServiceInit(repo repository.ICourseRepo, kafka kafka.Messaging, chName, chAnswer chan []byte) ICourseService {
	return &courseService{
		CourseRepo:  repo,
		kafkaClient: kafka,
		chName:      chName,
		chAnswer:    chAnswer,
	}
}

func (cr *courseService) ProduceKafka(name string) error {
	topic := "course-to-stud"
	key := []byte("course-name")
	msg := []byte(name)
	err := cr.kafkaClient.Write(topic, key, msg)
	if err != nil {
		return fmt.Errorf(path+"produce kafka: %w", err)
	}
	fmt.Println("Message sent successfully")
	return nil
}

func (cr *courseService) AnswerKafka() error {
	for {
		name := <-cr.chName
		fmt.Println("Answer kafka get!: ", name)
		courses, err := cr.CourseRepo.GetByStudent(string(name))
		if err != nil {
			return err
		}
		log.Printf("Курсы на которые записан студент: %+v", courses)

		topic := "answer-for-stud"
		key := []byte("answer")
		message, err := json.Marshal(courses)
		err = cr.kafkaClient.Write(topic, key, message)
		fmt.Println("Zapis proshla")
		if err != nil {
			return fmt.Errorf(path+"produce kafka: %w", err)
		}
		fmt.Println("Message sent successfully")
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
