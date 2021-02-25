package contract

import (
	"github.com/ecnuvj/vhoj_api/model/entity"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/user_problem_status"
)

type ListProblemsRequest struct {
	PageNo   int32 `json:"page_no" form:"page_no" binding:"required"`
	PageSize int32 `json:"page_size" form:"page_size" binding:"required"`
}

type ListProblemsResponse struct {
	BaseResponse *BaseResponse     `json:"base_response"`
	Problems     []*entity.Problem `json:"problems"`
	Page         *entity.Page      `json:"page"`
}

type ShowProblemRequest struct {
	ProblemId uint `json:"problem_id" form:"problem_id" binding:"required"`
}

type ShowProblemResponse struct {
	BaseResponse *BaseResponse   `json:"base_response"`
	Problem      *entity.Problem `json:"problem"`
}

type SearchProblemRequest struct {
	ProblemId uint   `json:"problem_id" form:"problem_id"`
	Title     string `json:"title" form:"title"`
}

type SearchProblemResponse struct {
	BaseResponse *BaseResponse     `json:"base_response"`
	Problems     []*entity.Problem `json:"problems"`
}

type CheckUserProblemStatusRequest struct {
	ProblemId uint `json:"problem_id" form:"problem_id" binding:"required"`
	UserId    uint `json:"user_id" form:"user_id" binding:"required"`
}

type CheckUserProblemStatusResponse struct {
	BaseResponse *BaseResponse                         `json:"base_response"`
	Status       user_problem_status.UserProblemStatus `json:"status"`
}
