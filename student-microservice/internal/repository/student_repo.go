package repository

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"

	"github.com/kanatovnurzhas/test-task-nis/student-microservice/internal/models"
)

type IStudentRepo interface {
	CreateStudent(stud models.Student) error
	GetAll() ([]models.Student, error)
	GetStudentByID(id int) (models.Student, error)
	UpdateStudent(stud models.Student) error
	DeleteStudent(id int) error
	GetByCourse(name string) ([]models.Student, error)
}

type studentRepo struct {
	db *sql.DB
}

const path = "repository"

func StudentRepoInit(db *sql.DB) IStudentRepo {
	return &studentRepo{
		db: db,
	}
}

func (repo *studentRepo) CreateStudent(stud models.Student) error {
	query := `INSERT INTO student(name,age,email,grade,courses) VALUES($1,$2,$3,$4,$5)`
	coursesArray := pq.Array(stud.Courses)
	_, err := repo.db.Exec(query, stud.Name, stud.Age, stud.Email, stud.Grade, coursesArray)
	if err != nil {
		return fmt.Errorf(path+"create student: %w", err)
	}
	return nil
}

func (repo *studentRepo) GetAll() ([]models.Student, error) {
	rows, err := repo.db.Query(`SELECT * FROM student`)
	if err != nil {
		return nil, fmt.Errorf(path+"get all student: %w", err)
	}
	defer rows.Close()
	var students []models.Student
	for rows.Next() {
		student := models.Student{}
		if err = rows.Scan(&student.ID, &student.Name, &student.Age, &student.Email, &student.Grade, pq.Array(&student.Courses)); err != nil {
			return nil, fmt.Errorf(path+"get all post, scan: %w", err)
		}
		students = append(students, student)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf(path+"get all post, rows error: %w", err)
	}
	return students, nil
}

func (repo *studentRepo) GetStudentByID(id int) (models.Student, error) {
	rows, err := repo.db.Query(`SELECT * FROM student WHERE id = $1`, id)
	if err != nil {
		return models.Student{}, fmt.Errorf(path+"get student by id: %w", err)
	}
	defer rows.Close()
	student := models.Student{}
	for rows.Next() {
		if err = rows.Scan(&student.ID, &student.Name, &student.Age, &student.Email, &student.Grade, pq.Array(&student.Courses)); err != nil {
			return models.Student{}, fmt.Errorf(path+"get student by id, scan: %w", err)
		}
		if err = rows.Err(); err != nil {
			return models.Student{}, fmt.Errorf(path+"get student by id, rows error: %w", err)
		}

	}
	return student, nil
}

func (repo *studentRepo) UpdateStudent(stud models.Student) error {
	stmt, err := repo.db.Prepare("UPDATE student SET name=$1, age=$2, email=$3, grade=$4, courses=$5 WHERE id=$6")
	if err != nil {
		return fmt.Errorf(path+"update student: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(stud.Name, stud.Age, stud.Email, stud.Grade, stud.Courses, stud.ID)
	if err != nil {
		return fmt.Errorf(path+"update student, exec: %w", err)
	}

	return nil
}

func (repo *studentRepo) DeleteStudent(id int) error {
	query := `DELETE FROM student WHERE id=$1`
	_, err := repo.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf(path+"delete student, exec: %w", err)
	}
	return nil
}

func (repo *studentRepo) GetByCourse(name string) ([]models.Student, error) {
	rows, err := repo.db.Query("SELECT * FROM student WHERE $1 = ANY(courses)", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var student models.Student
		err := rows.Scan(&student.ID, &student.Name, &student.Age, &student.Email, &student.Grade, pq.Array(&student.Courses))
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	fmt.Println(students)

	return students, nil
}
