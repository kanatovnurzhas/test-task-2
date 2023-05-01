package repository

import (
	"database/sql"
	"test-task-2/app1/models"
)

type IStudentRepo interface {
	CreateStudent(stud models.Student) error
	GetAll() ([]models.Student, error)
	GetStudentByID(id int) (models.Student, error)
	UpdateStudent(stud models.Student) error
	DeleteStudent(id int) error
}

type studentRepo struct {
	db *sql.DB
}

func StudentRepoInit(db *sql.DB) IStudentRepo {
	return &studentRepo{db}
}

func (repo *studentRepo) CreateStudent(stud models.Student) error {
	query := `INSERT INTO students(name,age,email,grade,courses)`
}
