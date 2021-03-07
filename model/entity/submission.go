package entity

import (
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/language"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/status_type"
)

type Submission struct {
	SubmissionId uint                             `json:"submission_id"`
	Username     string                           `json:"username"`
	ProblemId    uint                             `json:"problem_id"`
	UserId       uint                             `json:"user_id"`
	Result       status_type.SubmissionStatusType `json:"result"`
	TimeCost     int64                            `json:"time_cost"`
	MemoryCost   int64                            `json:"memory_cost"`
	Language     language.Language                `json:"language"`
}
