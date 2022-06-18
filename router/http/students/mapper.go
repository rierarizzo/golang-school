package students

import studentDomain "github.com/kenethrrizzo/golang-school/domain/students"

func toResponseModel(student *studentDomain.Student) *StudentResponse {
	return &StudentResponse{
		Id:      int(student.Id),
		Name:    student.Name,
		Surname: student.Surname,
	}
}
