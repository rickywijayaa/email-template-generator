package systemConfig

import (
	"email-template-generator/entity"
	"errors"
	"strings"
)

type Service interface {
	CreateSystemConfig(input CreateSystemConfigInput) (SystemConfigFormatter, error)
	UpdateSystemConfig(inputID GeneralIdInput, input UpdateSystemConfigInput) (SystemConfigFormatter, error)
	GetSystemConfigByCode(code string) ([]SystemConfigFormatter, error)
	DeleteSystemConfig(inputID GeneralIdInput) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateSystemConfig(input CreateSystemConfigInput) (SystemConfigFormatter, error) {
	if strings.Contains(input.Code, " ") {
		strings.TrimSpace(input.Code)
	}

	if strings.Contains(input.Value, " ") {
		strings.TrimSpace(input.Value)
	}

	isExist, err := s.repository.FindByCodeAndValue(input.Code, input.Value)
	if err != nil {
		return SystemConfigFormatter{}, err
	}

	if isExist.ID != 0 {
		return SystemConfigFormatter{}, errors.New("Cannot create with same systemconfig")
	}

	systemConfig := entity.SystemConfig{
		Code:  input.Code,
		Value: input.Value,
	}

	newSystemConfig, err := s.repository.Create(systemConfig)

	response := SystemConfigFormat(newSystemConfig)

	return response, nil
}

func (s *service) UpdateSystemConfig(inputID GeneralIdInput, input UpdateSystemConfigInput) (SystemConfigFormatter, error) {
	systemConfig, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return SystemConfigFormatter{}, err
	}

	systemConfig.Code = input.Code
	systemConfig.Value = input.Value

	updatedSystemConfig, err := s.repository.Update(systemConfig)

	response := SystemConfigFormat(updatedSystemConfig)

	return response, nil
}

func (s *service) GetSystemConfigByCode(code string) ([]SystemConfigFormatter, error) {
	systemConfig, err := s.repository.FindByCode(code)
	if err != nil {
		return []SystemConfigFormatter{}, err
	}

	if len(systemConfig) == 0 {
		return []SystemConfigFormatter{}, errors.New("Systemconfig not found")
	}

	response := SystemConfigsFormat(systemConfig)

	return response, nil
}

func (s *service) DeleteSystemConfig(inputID GeneralIdInput) (bool, error) {
	systemConfig, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return false, err
	}

	if systemConfig.ID == 0 {
		return false, errors.New("Systemconfig not found")
	}

	_, err = s.repository.Delete(systemConfig)
	if err != nil {
		return false, err
	}

	return true, nil
}
