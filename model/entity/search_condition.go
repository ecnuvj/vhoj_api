package entity

import (
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/contest_status"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/language"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/status_type"
)

type ProblemSearchCondition struct {
	Title     string `json:"title" form:"title"`
	ProblemId uint   `json:"problem_id" form:"problem_id"`
}

type ContestSearchCondition struct {
	Status      contest_status.ContestStatus `json:"status" form:"status"`
	Title       string                       `json:"title" form:"title"`
	CreatorName string                       `json:"creator_name" form:"creator_name"`
}

type SubmissionSearchCondition struct {
	Username  string                           `json:"username" form:"username"`
	ProblemId uint                             `json:"problem_id" form:"problem_id"`
	Status    status_type.SubmissionStatusType `json:"status" form:"status"`
	Language  language.Language                `json:"language" form:"language"`
}
