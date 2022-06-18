package students

import (
	"github.com/gin-gonic/gin"
	studentDomain "github.com/kenethrrizzo/golang-school/domain/students"
)

type StudentValidator struct {
	Name    string `binding:"required" json:"name"`
	Surname string `binding:"required" json:"surname"`
}

func Bind(c *gin.Context) (*studentDomain.Student, error) {
	var json StudentValidator

	if err := c.ShouldBindJSON(&json); err != nil {
		return nil, err
	}

	student := &studentDomain.Student{
		Name:    json.Name,
		Surname: json.Surname,
	}

	return student, nil
}
