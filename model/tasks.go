package model

import (
	"time"
)

// Определение перечисления
type StatusType string

const (
	NewStatus        StatusType = "new"
	InProgressStatus StatusType = "in_progress"
	DoneStatus       StatusType = "done"
)

func (s StatusType) String() string {
	switch s {
	case NewStatus:
		return "new"
	case InProgressStatus:
		return "in_progress"
	case DoneStatus:
		return "done"
	default:
		return "new"
	}
}

type Tasks struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      StatusType `json:"status"`
	Created_at  time.Time  `json:"created_at"`
	Updated_at  time.Time  `json:"updated_at"`
}
