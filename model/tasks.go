package model

import (
	"time"

	"gorm.io/gorm"
)

// Определение перечисления
type Status string

const (
	New         Status = "new"
	In_progress Status = "in_progress"
	Done        Status = "done"
)

type Tasks struct {
	gorm.Model
	Id          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"size:64"`
	Description string    `gorm:"size:255"`
	Status      Status    `gorm:"default:new"`
	Created_at  time.Time `gorm:"index"`
	Updated_at  time.Time `gorm:"index"`
}
