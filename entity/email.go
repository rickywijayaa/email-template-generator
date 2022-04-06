package entity

import "time"

type Email struct {
	ID           int       `gorm:"primary_key;auto_increment;column:id"`
	TemplateCode string    `gorm:"column:template_code"`
	Body         string    `gorm:"column:body"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}
