package entity

import (
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/language"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/status_type"
	"time"
)

type Submission struct {
	SubmissionId uint               `json:"submission_id"`
	Username     string             `json:"username"`
	ProblemId    uint               `json:"problem_id"`
	UserId       uint               `json:"user_id"`
	Result       SubmissionResult   `json:"result"`
	TimeCost     int64              `json:"time_cost"`
	MemoryCost   int64              `json:"memory_cost"`
	Language     SubmissionLanguage `json:"language"`
	SubmitTime   time.Time          `json:"submit_time"`
	Code         string             `json:"code"`
}

type SubmissionResult struct {
	Code status_type.SubmissionStatusType `json:"code"`
	Text string                           `json:"text"`
}

type SubmissionLanguage struct {
	Code language.Language `json:"code"`
	Text string            `json:"text"`
}
