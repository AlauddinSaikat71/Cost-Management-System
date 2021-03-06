package models

import "time"

type Cost struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	Amount      float32
	Date        time.Time
	Payment_Id  int
	Payment     Payment `gorm:"foreignKey:Payment_Id"`
	CreatedAt   time.Time
	CreatedBy   int `gorm:"foreignKey:CreatedBy"`
	UpdatedAt   time.Time
	UpdatedBy   int
}
