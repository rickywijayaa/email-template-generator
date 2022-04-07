package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(input LoginInput) (UserFormatter, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) Login(input LoginInput) (UserFormatter, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return UserFormatter{}, err
	}

	if user.ID == 0 {
		return UserFormatter{}, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return UserFormatter{}, errors.New("Incorrect Password")
	}

	response := UserFormat(user, user.Token)

	return response, nil
}
