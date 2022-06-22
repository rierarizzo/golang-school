package students

import "github.com/sirupsen/logrus"

type StudentService interface {
	CreateStudent(*Student) (*Student, error)
	ListAllStudents() ([]Student, error)
	GetStudentById(uint) (*Student, error)
	DeleteStudent(uint) error
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

func (s *Service) GetStudentById(id uint) (*Student, error) {
	student, err := s.repo.GetStudentById(id)
	if err != nil {
		logrus.Error("Error has ocurred: ", err.Error())
		return nil, err
	}
	

	logrus.Info("Student name: ", student.Name)
	return student, nil
}

func (s *Service) DeleteStudent(id uint) error {
	return s.repo.DeleteStudent(id)
}

func NewService(repo StudentRepository) *Service {
	return &Service{repo}
}
