package service

import (
	"github.com/kanatovnurzhas/test-task-2/app1/internal/models"
	"github.com/kanatovnurzhas/test-task-2/app1/internal/repository"
)

type IStudentService interface {
	CreateStudent(stud models.Student) error
	GetAll() ([]models.Student, error)
	GetStudentByID(id int) (models.Student, error)
	UpdateStudent(stud models.Student) error
	DeleteStudent(id int) error
	GetByCourse(name string) ([]models.Student, error)
}

type studentService struct {
	StudRepo repository.IStudentRepo
}

func StudentServiceInit(repo repository.IStudentRepo) IStudentService {
	return &studentService{
		StudRepo: repo,
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
