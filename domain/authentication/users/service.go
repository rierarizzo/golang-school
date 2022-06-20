package users

import userDTO "github.com/kenethrrizzo/golang-school/dto/users"

type AuthService interface {
	Login(*User) (*userDTO.LoginResponse, error)
}
