package models

import "time"

type Cost struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	Amount      float32
	Payment_Id  int
	Payment     Payment  `gorm:"foreignKey:Payment_Id"`
	CreatedAt   time.Time
	CreatedBy 	int
	UpdatedAt   time.Time
	UpdatedBy   int
}