package repository

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"

	"github.com/kanatovnurzhas/test-task-2/app2/internal/models"
)

type ICourseRepo interface {
	CreateCourse(course models.Course) error
	GetAll() ([]models.Course, error)
	GetCourseByID(id int) (models.Course, error)
	UpdateCourse(stud models.Course) error
	DeleteCourse(id int) error
	GetByStudent(name string) ([]models.Course, error)
}

type courseRepo struct {
	db *sql.DB
}

const path = "repository"

func CourseRepoInit(db *sql.DB) ICourseRepo {
	return &courseRepo{
		db: db,
	}
}

func (repo *courseRepo) CreateCourse(course models.Course) error {
	query := `INSERT INTO course(name,teacher,students) VALUES($1,$2,$3)`
	studentsArray := pq.Array(course.Students)
	_, err := repo.db.Exec(query, course.Name, course.Teacher, studentsArray)
	if err != nil {
		return fmt.Errorf(path+"create Course: %w", err)
	}
	return nil
}

func (repo *courseRepo) GetAll() ([]models.Course, error) {
	rows, err := repo.db.Query(`SELECT * FROM course`)
	if err != nil {
		return nil, fmt.Errorf(path+"get all course: %w", err)
	}
	defer rows.Close()
	var courses []models.Course
	for rows.Next() {
		course := models.Course{}
		if err = rows.Scan(&course.ID, &course.Name, course.Teacher, pq.Array(&course.Students)); err != nil {
			return nil, fmt.Errorf(path+"get all course, scan: %w", err)
		}
		courses = append(courses, course)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf(path+"get all course, rows error: %w", err)
	}
	return courses, nil
}

func (repo *courseRepo) GetCourseByID(id int) (models.Course, error) {
	rows, err := repo.db.Query(`SELECT * FROM course WHERE id = $1`, id)
	if err != nil {
		return models.Course{}, fmt.Errorf(path+"get course by id: %w", err)
	}
	defer rows.Close()
	course := models.Course{}
	for rows.Next() {
		if err = rows.Scan(&course.ID, course.Name, course.Teacher, pq.Array(&course.Students)); err != nil {
			return models.Course{}, fmt.Errorf(path+"get course by id, scan: %w", err)
		}
		if err = rows.Err(); err != nil {
			return models.Course{}, fmt.Errorf(path+"get course by id, rows error: %w", err)
		}

	}
	return course, nil
}

func (repo *courseRepo) UpdateCourse(course models.Course) error {
	stmt, err := repo.db.Prepare("UPDATE course SET name=$1, teacher=$2, students=$3 WHERE id=$4")
	if err != nil {
		return fmt.Errorf(path+"update course: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(course.Name, course.Teacher, course.Teacher, course.Students)
	if err != nil {
		return fmt.Errorf(path+"update course, exec: %w", err)
	}

	return nil
}

func (repo *courseRepo) DeleteCourse(id int) error {
	query := `DELETE FROM course WHERE id=$1`
	_, err := repo.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf(path+"delete course, exec: %w", err)
	}
	return nil
}

func (repo *courseRepo) GetByStudent(name string) ([]models.Course, error) {
	rows, err := repo.db.Query("SELECT * FROM course WHERE $1 = ANY(students)", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []models.Course
	for rows.Next() {
		var course models.Course
		err := rows.Scan(&course.ID, &course.Name, &course.Teacher, pq.Array(&course.Students))
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	fmt.Println(courses)

	return courses, nil
}
