package students

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kenethrrizzo/golang-school/domain/students"
	studentDTO "github.com/kenethrrizzo/golang-school/dto/students"
	"github.com/sirupsen/logrus"
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
			var responseItems = make([]studentDTO.StudentResponse, len(results))
			for i, element := range results {
				responseItems[i] = *toResponseModel(&element)
			}
			response := &studentDTO.ListResponse{
				Data: responseItems,
			}

			ex := c.MustGet("ex").(string)

			logrus.Info("Example received from middleware: " + ex)

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

			response := &studentDTO.StudentResponse{
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
