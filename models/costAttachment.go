package models

import (
	"time"
)

type CostAttachment struct {
	ID           uint `gorm:"primaryKey"`
	CostId       int
	Cost         Cost `gorm:"foreignKey:CostId"`
	AttachmentId int
	Attachment   Attachment `gorm:"foreignKey:AttachmentId"`
	CreatedAt    time.Time
	CreatedBy    int
}
