package students

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kenethrrizzo/golang-school/domain/students"
)

func NewRoutesFactory(group *gin.RouterGroup) func(service students.StudentService) {
	studentRoutesFactory := func(service students.StudentService) {

		// Get all students
		group.GET("/", func(c *gin.Context) {
			results, err := service.ListAllStudents()
			if err != nil {
				c.Error(err)
				return
			}
			var responseItems = make([]StudentResponse, len(results))
			for i, element := range results {
				responseItems[i] = *toResponseModel(&element)
			}
			response := &ListResponse{
				Data: responseItems,
			}
			c.JSON(http.StatusOK, response)
		})

		// Get student by id
		group.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.Error(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			result, err := service.GetStudentById(uint(id))
			if err != nil {
				c.Error(err)
				return
			}

			response := &StudentResponse{
				Id:      int(result.Id),
				Name:    result.Name,
				Surname: result.Surname,
			}
			c.JSON(http.StatusOK, response)
		})

		// Create new student
		group.POST("/", func(c *gin.Context) {
			student, err := Bind(c)
			if err != nil {
				c.Error(err)
				return
			}
			newStudent, err := service.CreateStudent(student)
			if err != nil {
				c.Error(err)
				return
			}
			c.JSON(http.StatusCreated, *toResponseModel(newStudent))
		})

	}
	return studentRoutesFactory
}
