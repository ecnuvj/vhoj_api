package entity

import "github.com/ecnuvj/vhoj_common/pkg/common/constants/user_problem_status"

type ContestProblem struct {
	ContestId    string                                `json:"contest_id" form:"contest_id"`
	ProblemId    string                                `json:"problem_id" form:"problem_id"`
	ProblemOrder string                                `json:"problem_order" form:"problem_order"`
	Title        string                                `json:"title" form:"title"`
	Accepted     uint                                  `json:"accepted" form:"accepted"`
	Submitted    uint                                  `json:"submitted" form:"submitted"`
	Status       user_problem_status.UserProblemStatus `json:"status" form:"status"`
}

type ContestProblems []*ContestProblem

func (c ContestProblems) Len() int {
	return len(c)
}

func (c ContestProblems) Less(i, j int) bool {
	return c[i].ProblemOrder < c[j].ProblemOrder
}

func (c ContestProblems) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
