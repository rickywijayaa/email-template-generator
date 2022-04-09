package systemConfig

import (
	"email-template-generator/entity"

	"gorm.io/gorm"
)

type Repository interface {
	Create(systemConfig entity.SystemConfig) (entity.SystemConfig, error)
	Update(systemConfig entity.SystemConfig) (entity.SystemConfig, error)
	FindByCode(code string) ([]entity.SystemConfig, error)
	FindByCodeAndValue(code string, value string) (entity.SystemConfig, error)
	FindByID(ID int) (entity.SystemConfig, error)
	Delete(systemConfig entity.SystemConfig) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(systemConfig entity.SystemConfig) (entity.SystemConfig, error) {
	err := r.db.Save(&systemConfig).Error
	if err != nil {
		return systemConfig, err
	}

	return systemConfig, nil
}

func (r *repository) Update(systemConfig entity.SystemConfig) (entity.SystemConfig, error) {
	err := r.db.Save(&systemConfig).Error
	if err != nil {
		return systemConfig, err
	}

	return systemConfig, nil
}

func (r *repository) FindByCode(code string) ([]entity.SystemConfig, error) {
	var systemConfig []entity.SystemConfig
	err := r.db.Where("code = ?", code).Find(&systemConfig).Error
	if err != nil {
		return systemConfig, err
	}

	return systemConfig, nil
}

func (r *repository) FindByCodeAndValue(code string, value string) ([]entity.SystemConfig, error) {
	var systemConfig []entity.SystemConfig
	err := r.db.Where("code = ? AND value = ?", code, value).Find(&systemConfig).Error
	if err != nil {
		return systemConfig, err
	}

	return systemConfig, nil
}

func (r *repository) FindByID(ID int) (entity.SystemConfig, error) {
	var systemConfig entity.SystemConfig
	err := r.db.Where("id = ?", ID).Find(&systemConfig).Error
	if err != nil {
		return systemConfig, err
	}

	return systemConfig, nil
}

func (r *repository) Delete(systemConfig entity.SystemConfig) (bool, error) {
	err := r.db.Delete(&systemConfig).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
