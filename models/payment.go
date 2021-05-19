package models

import "time"

type Payment struct {
	ID        uint `gorm:"primaryKey"`
	Method    string
	Amount    float32
	CreatedAt time.Time
	CreatedBy int
	PaidBy    int
	Meta      string
}
