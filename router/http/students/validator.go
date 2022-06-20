package students

import (
	"github.com/gin-gonic/gin"
	studentDomain "github.com/kenethrrizzo/golang-school/domain/students"
	studentDTO "github.com/kenethrrizzo/golang-school/dto/students"
)

func Bind(c *gin.Context) (*studentDomain.Student, error) {
	var json studentDTO.StudentRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		return nil, err
	}

	student := &studentDomain.Student{
		Name:    json.Name,
		Surname: json.Surname,
	}

	return student, nil
}
