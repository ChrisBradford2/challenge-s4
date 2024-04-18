package models

import (
	"time"
)

type Step struct {
	Base
	Note         uint8
	Status       string
	DeadLineDate time.Time
}
