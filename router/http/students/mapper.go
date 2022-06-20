package students

import (
	studentDomain "github.com/kenethrrizzo/golang-school/domain/students"
	studentDTO "github.com/kenethrrizzo/golang-school/dto/students"
)

func toResponseModel(student *studentDomain.Student) *studentDTO.StudentResponse {
	return &studentDTO.StudentResponse{
		Id:      int(student.Id),
		Name:    student.Name,
		Surname: student.Surname,
	}
}
