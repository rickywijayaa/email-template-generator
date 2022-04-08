package user

import (
	"email-template-generator/entity"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(input LoginInput) (UserFormatter, error)
	GetUserByID(ID int) (UserFormatter, error)
	ChangePassword(currentUser entity.User, input ChangePasswordInput) (UserFormatter, error)
	UpdateToken(input UserFormatter) (UserFormatter, error)
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

func (s *service) GetUserByID(ID int) (UserFormatter, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return UserFormatter{}, err
	}

	if user.ID == 0 {
		return UserFormatter{}, errors.New("User not found")
	}

	response := UserFormat(user, user.Token)

	return response, nil
}

func (s *service) ChangePassword(currentUser entity.User, input ChangePasswordInput) (UserFormatter, error) {
	user, err := s.repository.FindByID(currentUser.ID)
	if err != nil {
		return UserFormatter{}, err
	}

	if user.ID == 0 {
		return UserFormatter{}, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.OldPassword))
	if err != nil {
		return UserFormatter{}, errors.New("Incorrect Old Password")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return UserFormatter{}, errors.New("Something wrong please contact administrator")
	}

	user.Password = string(password)
	newUser, err := s.repository.Update(user)
	if err != nil {
		return UserFormatter{}, err
	}

	response := UserFormat(newUser, newUser.Token)

	return response, nil
}

func (s *service) UpdateToken(input UserFormatter) (UserFormatter, error) {
	user, err := s.repository.FindByID(input.ID)
	if err != nil {
		return UserFormatter{}, err
	}

	if user.ID == 0 {
		return UserFormatter{}, errors.New("User not found")
	}

	user.Token = input.Token
	newUser, err := s.repository.Update(user)
	if err != nil {
		return UserFormatter{}, err
	}

	response := UserFormat(newUser, newUser.Token)

	return response, nil
}
