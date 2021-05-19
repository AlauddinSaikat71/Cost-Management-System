package models

import (
	"time"
)

type Attachment struct {
	ID        uint `gorm:"primaryKey"`
	FilePath  string
	FileType  string
	CreatedAt time.Time
	CreatedBy int
}
