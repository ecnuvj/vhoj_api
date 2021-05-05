package entity

import (
	"time"
)

type RawProblem struct {
	RawProblemId    uint64    `json:"raw_problem_id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	SampleInput     string    `json:"sample_input"`
	SampleOutput    string    `json:"sample_output"`
	Input           string    `json:"input"`
	Output          string    `json:"output"`
	Hint            string    `json:"hint"`
	RemoteOj        string    `json:"remote_oj"`
	RemoteProblemId string    `json:"remote_problem_id"`
	RemoteSubmitId  string    `json:"remote_submit_id"`
	TimeLimit       string    `json:"time_limit"`
	MemoryLimit     string    `json:"memory_limit"`
	Spj             string    `json:"spj"`
	Std             string    `json:"std"`
	Source          string    `json:"source"`
	GroupId         uint64    `json:"group_id"`
	UpdatedAt       time.Time `json:"updated_at"`
	ProblemId       uint64    `json:"problem_id"`
}
