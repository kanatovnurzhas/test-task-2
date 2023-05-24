package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kanatovnurzhas/test-task-2/pkg/database"
	"github.com/kanatovnurzhas/test-task-2/pkg/kafka"
	"github.com/kanatovnurzhas/test-task-2/student-microservice/internal/handler"
	"github.com/kanatovnurzhas/test-task-2/student-microservice/internal/repository"
	"github.com/kanatovnurzhas/test-task-2/student-microservice/internal/service"
	"github.com/sirupsen/logrus"
)

var topics = []string{"course-to-stud", "answer-for-stud", "stud-to-course", "answer-for-course"}

func main() {
	db, err := database.ConnectionDB()
	if err != nil {
		wrappedErr := fmt.Errorf("connection db refused: %w", err)
		fmt.Println(wrappedErr)
		return
	}
	log.Println("Connection success!")
	err = CreateTable(db)
	if err != nil {
		wrappedErr := fmt.Errorf("create table: %w", err)
		fmt.Println(wrappedErr)
		return
	}
	log.Println("Create table success!")

	channelName := make(chan []byte)
	channelAnswer := make(chan []byte)
	kafkaClient := kafka.NewKafkaClient("localhost:9092", topics, context.Background(), channelName, channelAnswer)

	go kafkaClient.Read("answer-for-stud")

	sr := repository.StudentRepoInit(db)
	ss := service.StudentServiceInit(sr, kafkaClient, channelName, channelAnswer)
	sh := handler.StudentHandlerInit(ss)

	server := fiber.New()

	logger := logrus.New()

	server.Use(func(ctx *fiber.Ctx) error {
		logger.Infof("Incoming request: %s %s", ctx.Method(), ctx.Path())
		return ctx.Next()
	})

	basePath := server.Group("/api")
	sh.RegisterStudentRoutes(basePath)

	// data, _ := json.MarshalIndent(server.Stack(), "", " ")
	// fmt.Println(string(data))

	if err := server.Listen(":7777"); err != nil {
		logger.Fatalf("Error starting server: %s", err)
	}
}

const studentTable = `
				CREATE TABLE IF NOT EXISTS student (
					id SERIAL PRIMARY KEY,
					name TEXT NOT NULL,
					age INTEGER NOT NULL,
					email TEXT NOT NULL UNIQUE,
					grade TEXT NOT NULL,
					courses TEXT[]
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
