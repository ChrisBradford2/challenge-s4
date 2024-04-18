package models

import (
	"time"
)

type Step struct {
	Base
	Note         uint8     `json:"note" example:"5"`
	Status       string    `json:"status" example:"done"`
	DeadLineDate time.Time `json:"dead_line_date" example:"2025-12-31T23:59:59Z"`
}
