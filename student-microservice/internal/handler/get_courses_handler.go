package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kanatovnurzhas/test-task-2/student-microservice/internal/models"
)

func (st *StudentHandler) GetCourses(ctx *fiber.Ctx) error {
	name := ctx.Params("name")

	request := models.Request{
		Method: "GET",
		Url:    "http://localhost:7778/api/courses/getByStudent/:" + string(name),
	}
	newReq, err := http.NewRequest(request.Method, request.Url, nil)
	if err != nil {
		wrappedErr := fmt.Errorf("error is: %w", err)
		fmt.Println(wrappedErr)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"status":  fiber.StatusInternalServerError,
		})
	}
	client := http.Client{}
	response, err := client.Do(newReq)
	data, _ := ioutil.ReadAll(response.Body)
	var resp models.Response
	if err := json.Unmarshal(data, &resp); err != nil {
		wrappedErr := fmt.Errorf("error is: %w", err)
		fmt.Println(wrappedErr)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"status":  fiber.StatusInternalServerError,
		})
	}

	if err != nil {
		wrappedErr := fmt.Errorf("error is: %w", err)
		fmt.Println(wrappedErr)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"status":  fiber.StatusInternalServerError,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"courses": resp,
	})
}
