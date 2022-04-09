package entity

type SystemConfig struct {
	ID    int    `gorm:"primary_key;auto_increment;column:id"`
	Code  string `gorm:"column:code"`
	Value string `gorm:"column:value"`
}
