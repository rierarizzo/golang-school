package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kenethrrizzo/golang-school/domain/students"
	studentRoutes "github.com/kenethrrizzo/golang-school/router/http/students"
)

func NewHTTPHandler(studentService students.StudentService) http.Handler {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")

	router.Use(cors.New(config))

	api := router.Group("/api")

	studentGroup := api.Group("/students")

	studentRoutes.NewRoutesFactory(studentGroup)(studentService)

	return router
}
