package students

import (
	domain "github.com/kenethrrizzo/golang-school/domain/students"
)

func toDBModel(entity *domain.Student) *Student {
	return &Student{
		Id:      entity.Id,
		Name:    entity.Name,
		Surname: entity.Surname,
	}
}

func toDomainModel(entity *Student) *domain.Student {
	return &domain.Student{
		Id:        entity.Id,
		Name:      entity.Name,
		Surname:   entity.Surname,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
