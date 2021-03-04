package entity

import (
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/contest_status"
	"time"
)

type Contest struct {
	ContestId   uint                         `json:"contest_id" form:"contest_id"`
	Title       string                       `json:"title" form:"title" binding:"required"`
	Description string                       `json:"description" form:"description" binding:"required"`
	StartTime   time.Time                    `json:"start_time" form:"start_time" binding:"required"`
	EndTime     time.Time                    `json:"end_time" form:"end_time" binding:"required"`
	Creator     *User                        `json:"creator" form:"creator"`
	ProblemIds  []uint                       `json:"problem_ids" form:"problem_ids"`
	Problems    []*Problem                   `json:"problems" form:"problems"`
	Status      contest_status.ContestStatus `json:"status" form:"status"`
}
