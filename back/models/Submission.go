package models

type Submission struct {
	Base
	TeamID       uint
	Team         Team
	EvaluationID uint
	Evaluation   Evaluation
	Status       string
	AttachedFile string
	StepID       uint
	Step         Step
}
