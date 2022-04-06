package entity

import "time"

type User struct {
	ID        int       `gorm:"primary_key;auto_increment;column:id"`
	Name      string    `gorm:"column:nama"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	Token     string    `gorm:"column:token"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
