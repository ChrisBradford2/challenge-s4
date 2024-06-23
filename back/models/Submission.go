package models

type Submission struct {
	Base
	TeamID       uint       `json:"team_id"`
	Team         Team       `json:"team"`
	EvaluationID uint       `json:"evaluation_id"`
	Evaluation   Evaluation `json:"evaluation"`
	Status       string     `json:"status"`
	GitLink      string     `json:"git_link"`
	StepID       uint       `json:"step_id"`
	Step         Step       `json:"step"`
}

type SubmissionCreate struct {
	TeamID       uint       `json:"team_id"`
	Team         Team       `json:"team"`
	EvaluationID uint       `json:"evaluation_id"`
	Evaluation   Evaluation `json:"evaluation"`
	Status       string     `json:"status"`
	GitLink      string     `json:"git_link"`
}
