package students

type StudentService interface {
	CreateStudent(*Student) (*Student, error)
	ListAllStudents() ([]Student, error)
}

type Service struct {
	repo StudentRepository
}

func (s *Service) CreateStudent(student *Student) (*Student, error) {
	return s.repo.CreateStudent(student)
}

func (s *Service) ListAllStudents() ([]Student, error) {
	return s.repo.ListAllStudents()
}

func NewService(repo StudentRepository) *Service {
	return &Service{repo}
}
