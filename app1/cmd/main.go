package main

import (
	"log"
	"test-task-2/app1/server/http"
)

func main() {
	if err := http.RunServer(); err != nil {
		log.Fatal("salamaleikum")
	}
}
