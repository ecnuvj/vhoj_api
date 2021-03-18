package entity

import "github.com/ecnuvj/vhoj_common/pkg/common/constants/user_problem_status"

type Problem struct {
	ProblemId    uint                                  `json:"problem_id"`
	Title        string                                `json:"title"`
	Description  string                                `json:"description"`
	SampleInput  string                                `json:"sample_input"`
	SampleOutput string                                `json:"sample_output"`
	Input        string                                `json:"input"`
	Output       string                                `json:"output"`
	TimeLimit    string                                `json:"time_limit"`
	MemoryLimit  string                                `json:"memory_limit"`
	Submitted    int64                                 `json:"submitted"`
	Accepted     int64                                 `json:"accepted"`
	Status       user_problem_status.UserProblemStatus `json:"status"`
}
