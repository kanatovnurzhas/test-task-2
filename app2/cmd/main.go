package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kanatovnurzhas/test-task-2/app2/internal/handler"
	"github.com/kanatovnurzhas/test-task-2/app2/internal/repository"
	"github.com/kanatovnurzhas/test-task-2/app2/internal/service"
	"github.com/kanatovnurzhas/test-task-2/pkg"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := pkg.ConnectionDB()
	if err != nil {
		wrappedErr := fmt.Errorf("connection db refused: %w", err)
		fmt.Println(wrappedErr)
		return
	}
	fmt.Println("Connection success!")
	err = CreateTable(db)
	if err != nil {
		wrappedErr := fmt.Errorf("create table: %w", err)
		fmt.Println(wrappedErr)
		return
	}
	fmt.Println("Create table success!")
	cr := repository.CourseRepoInit(db)
	cs := service.CourseServiceInit(cr)
	ch := handler.CourseHandlerInit(cs)

	server := fiber.New()

	logger := logrus.New()

	server.Use(func(ctx *fiber.Ctx) error {
		logger.Infof("Incoming request: %s %s", ctx.Method(), ctx.Path())
		return ctx.Next()
	})

	basePath := server.Group("/api")
	ch.RegisterCourseRoutes(basePath)

	data, _ := json.MarshalIndent(server.Stack(), "", " ")
	fmt.Println(string(data))

	if err := server.Listen(":7778"); err != nil {
		logger.Fatalf("Error starting server: %s", err)
	}
}

const studentTable = `
				CREATE TABLE IF NOT EXISTS course (
					id SERIAL PRIMARY KEY,
					name TEXT NOT NULL UNIQUE,
					teacher TEXT NOT NULL,
					students TEXT[]
					);
				`

// функция для создания таблицы
func CreateTable(db *sql.DB) error {
	_, err := db.Query(studentTable)
	if err != nil {
		return fmt.Errorf("create table: %w", err)
	}
	return nil
}
