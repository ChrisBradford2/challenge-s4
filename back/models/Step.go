package models

import (
	"gorm.io/gorm"
	"time"
)

type Step struct {
	gorm.Model
	Note         uint8
	Status       string
	DeadLineDate time.Time
}
