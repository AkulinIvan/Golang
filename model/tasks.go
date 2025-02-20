package model

import (
	"fmt"
	"strings"
	"time"
)

// Определение перечисления
type StatusType int

const (
	NewStatus StatusType = iota + 1
	InProgressStatus
	DoneStatus
)

var statusMap = map[string]StatusType{
	"new":         NewStatus,
	"in_progress": InProgressStatus,
	"done":        DoneStatus,
}

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

func ParseStatus(statusStr string) (StatusType, error) {
	statusStr = strings.ToLower(strings.TrimSpace(statusStr))
	s, ok := statusMap[statusStr]
	if !ok {
		return 0, fmt.Errorf("invalid status: %s", statusStr)
	}
	return s, nil
}

type Tasks struct {
	Id          uint32     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      StatusType `json:"status"`
	Created_at  time.Time  `json:"created_at"`
	Updated_at  time.Time  `json:"updated_at"`
}
