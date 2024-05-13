package model

import "gorm.io/gorm"

type (
	Transaction struct {
		ID    uint   `gorm:"primarykey"`
		Key   string `gorm:"column:key"`
		Value string `gorm:"column:value"`
	}
)

func (Transaction) TableName() string {
	return "payment"
}
func (m *Transaction) Create(db *gorm.DB) error {
	return db.Create(&m).Error
}
