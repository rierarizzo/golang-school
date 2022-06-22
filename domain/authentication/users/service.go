package users

import "github.com/kenethrrizzo/golang-school/domain/authentication/jwt"

type AuthService interface {
	Login(*User) (*jwt.Authentication, error)
}
