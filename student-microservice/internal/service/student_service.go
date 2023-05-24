package service

import (
	"encoding/json"
	"fmt"

	"github.com/kanatovnurzhas/test-task-2/pkg/kafka"
	"github.com/kanatovnurzhas/test-task-2/student-microservice/internal/models"
	"github.com/kanatovnurzhas/test-task-2/student-microservice/internal/repository"
)

type IStudentService interface {
	CreateStudent(stud models.Student) error
	GetAll() ([]models.Student, error)
	GetStudentByID(id int) (models.Student, error)
	UpdateStudent(stud models.Student) error
	DeleteStudent(id int) error
	GetByCourse(name string) ([]models.Student, error)
	ProduceKafka(name string) error
	AnswerForStud() (models.Response, error)
	// ConsumeKafka() ([]models.Response, error)
}

const path = "service"

type studentService struct {
	StudRepo    repository.IStudentRepo
	kafkaClient kafka.Messaging
	chName      chan []byte
	chAnswer    chan []byte
}

func StudentServiceInit(repo repository.IStudentRepo, kafka kafka.Messaging, chName, chAnswer chan []byte) IStudentService {
	return &studentService{
		StudRepo:    repo,
		kafkaClient: kafka,
		chName:      chName,
		chAnswer:    chAnswer,
	}
}

func (st *studentService) ProduceKafka(name string) error {
	topic := "stud-to-course"
	key := []byte("student")
	msg := []byte(name)
	err := st.kafkaClient.Write(topic, key, msg)
	if err != nil {
		return fmt.Errorf(path+"produce kafka: %w", err)
	}
	fmt.Println("Message sent successfully: ", string(msg))
	return nil
}

func (st *studentService) AnswerForStud() (models.Response, error) {
	for {
		fmt.Println("answer for stud")
		coursesByte := <-st.chAnswer

		var courses models.Response
		err := json.Unmarshal(coursesByte, &courses.Course)
		if err != nil {
			return courses, err
		}
		return courses, nil
	}
}

func (st *studentService) CreateStudent(stud models.Student) error {
	// какая то бизнес логика
	return st.StudRepo.CreateStudent(stud)
}

func (st *studentService) GetAll() ([]models.Student, error) {
	// какая то бизнес логика
	return st.StudRepo.GetAll()
}

func (st *studentService) GetStudentByID(id int) (models.Student, error) {
	// какая то бизнес логика
	return st.StudRepo.GetStudentByID(id)
}

func (st *studentService) UpdateStudent(stud models.Student) error {
	// какая то бизнес логика
	return st.StudRepo.UpdateStudent(stud)
}

func (st *studentService) DeleteStudent(id int) error {
	// какая то бизнес логика
	return st.StudRepo.DeleteStudent(id)
}

func (st *studentService) GetByCourse(name string) ([]models.Student, error) {
	// какая то бизнес логика
	return st.StudRepo.GetByCourse(name)
}
