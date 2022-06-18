package main

import (
	"net/http"

	"github.com/kenethrrizzo/golang-school/config"
	"github.com/kenethrrizzo/golang-school/data/database"
	studentStore "github.com/kenethrrizzo/golang-school/data/students"
	"github.com/kenethrrizzo/golang-school/domain/students"
	router "github.com/kenethrrizzo/golang-school/router/http"
)

func main() {
	config := config.NewConfig()
	db, err := database.Connect(config.Database)
	if err != nil {
		panic(err)
	}

	studentsRepo := studentStore.New(db)
	studentsService := students.NewService(studentsRepo)

	httpRouter := router.NewHTTPHandler(studentsService)

	err = http.ListenAndServe(":"+config.Port, httpRouter)
	if err != nil {
		panic(err)
	}
}
