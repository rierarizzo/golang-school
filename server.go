package main

import (
	"net/http"

	"github.com/kenethrrizzo/golang-school/config"
	"github.com/kenethrrizzo/golang-school/data/database"
)

func main() {
	config := config.NewConfig()
	_, err := database.Connect(config.Database)
	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		panic(err)
	}
}
