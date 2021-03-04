package entity

import "github.com/ecnuvj/vhoj_common/pkg/common/constants/contest_status"

type ProblemSearchCondition struct {
	Title     string `json:"title" form:"title"`
	ProblemId uint   `json:"problem_id" form:"problem_id"`
}

type ContestSearchCondition struct {
	Status      contest_status.ContestStatus `json:"status" form:"status"`
	Title       string                       `json:"title" form:"title"`
	CreatorName string                       `json:"creator_name" form:"creator_name"`
}
