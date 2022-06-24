package users

import "github.com/kenethrrizzo/golang-school/domain/dto"

type AuthService interface {
	Login(*User) (*dto.Authentication, error)
}
