package models

import "gorm.io/gorm"

type Submission struct {
	gorm.Model
	TeamID       uint
	Team         Team
	EvaluationID uint
	Evaluation   Evaluation
	Status       string
	AttachedFile string
	StepID       uint
	Step         Step
}
