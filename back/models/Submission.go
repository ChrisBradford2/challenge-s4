package models

type Submission struct {
	Base
	TeamID       uint       `json:"team_id"`
	Team         Team       `json:"team"`
	EvaluationID *uint      `json:"evaluation_id,omitempty"`
	Evaluation   Evaluation `json:"evaluation"`
	Status       string     `json:"status"`
	GitLink      string     `json:"git_link"`
	StepID       *uint      `json:"step_id,omitempty"`
	Step         Step       `json:"step,omitempty"`
	FileURL      string     `json:"file_url"`
}

type SubmissionCreate struct {
	Base
	TeamID  uint   `json:"team_id"`
	Team    Team   `json:"team"`
	Status  string `json:"status"`
	GitLink string `json:"git_link"`
	FileURL string `json:"file_url"`
}
