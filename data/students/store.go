package students

import (
	studentDomain "github.com/kenethrrizzo/golang-school/domain/students"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Store {
	db.AutoMigrate(&Student{})
	return &Store{db}
}

func (s Store) CreateStudent(student *studentDomain.Student) (*studentDomain.Student, error) {
	entity := toDBModel(student)

	if err := s.db.Create(entity).Error; err != nil {
		return nil, err
	}

	return toDomainModel(entity), nil
}

func (s Store) ListAllStudents() ([]studentDomain.Student, error) {
	var results []Student

	err := s.db.Find(&results).Error
	if err != nil {
		return nil, err
	}

	var students = make([]studentDomain.Student, len(results))

	for i, element := range results {
		students[i] = *toDomainModel(&element)
	}

	return students, nil
}

func (s Store) GetStudentById(id uint) (*studentDomain.Student, error) {
	result := Student{Id: id}

	err := s.db.Find(&result).Error
	if err != nil {
		return nil, err
	}

	student := *toDomainModel(&result)

	return &student, nil
}

func (s Store) DeleteStudent(id uint) error {
	result := Student{Id: id}

	err := s.db.Delete(&result).Error
	if err != nil {
		return err
	}

	return nil
}