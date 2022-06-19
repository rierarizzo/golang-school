package students

type StudentRepository interface {
	CreateStudent(*Student) (*Student, error)
	ListAllStudents() ([]Student, error)
	GetStudentById(uint) (*Student, error)
}
