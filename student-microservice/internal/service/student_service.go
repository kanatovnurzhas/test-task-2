package service

import (
	"fmt"
	"global/pkg/kafka"

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
	// ConsumeKafka() ([]models.Response, error)
}

const path = "service"

type studentService struct {
	StudRepo    repository.IStudentRepo
	kafkaClient kafka.Messaging
}

func StudentServiceInit(repo repository.IStudentRepo, kafka kafka.Messaging) IStudentService {
	return &studentService{
		StudRepo:    repo,
		kafkaClient: kafka,
	}
}

func (st *studentService) ProduceKafka(name string) error {
	topic := "student_to_course"
	key := []byte("student-name")
	msg := []byte(name)
	err := st.kafkaClient.Write(topic, key, msg)
	if err != nil {
		return fmt.Errorf(path+"produce cafka: %w", err)
	}
	fmt.Println("Message sent successfully")
	return nil
}

// func (st *studentService) ConsumeKafka() ([]models.Response, error) {
// }

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
