package students

type StudentRepository interface {
	CreateStudent(*Student) (*Student, error)
	ListAllStudents() ([]Student, error)
}