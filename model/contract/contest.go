package contract

import (
	"github.com/ecnuvj/vhoj_api/model/entity"
)

type CreateContestRequest struct {
	Title       string                   `json:"title" form:"title" binding:"required"`
	Description string                   `json:"description" form:"description"`
	StartTime   int64                    `json:"start_time" time_format:"unix" form:"start_time"  binding:"required"`
	EndTime     int64                    `json:"end_time" time_format:"unix" form:"end_time" binding:"required"`
	Problems    []*entity.ContestProblem `json:"problems" form:"problems"`
}

type CreateContestResponse struct {
	Contest      *entity.Contest `json:"contest"`
	BaseResponse *BaseResponse   `json:"base_response"`
}

type ListContestsRequest struct {
	PageNo   int32 `json:"page_no" form:"page_no" binding:"required"`
	PageSize int32 `json:"page_size" form:"page_size" binding:"required"`
}

type ListContestsResponse struct {
	BaseResponse *BaseResponse     `json:"base_response"`
	Contests     []*entity.Contest `json:"contests"`
	PageInfo     *entity.Page      `json:"page_info"`
}

type SearchContestRequest struct {
	SearchCondition *entity.ContestSearchCondition `json:"search_condition" form:"search_condition" binding:"required"`
	PageNo          int32                          `json:"page_no" form:"page_no" binding:"required"`
	PageSize        int32                          `json:"page_size" form:"page_size" binding:"required"`
}

type SearchContestResponse struct {
	Contests     []*entity.Contest `json:"contests"`
	BaseResponse *BaseResponse     `json:"base_response"`
	PageInfo     *entity.Page      `json:"page_info"`
}

type AddContestProblemRequest struct {
	ContestId uint `json:"contest_id" form:"contest_id" binding:"required"`
	ProblemId uint `json:"problem_id" form:"problem_id" binding:"required"`
}

type AddContestProblemResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}

type DeleteContestProblemRequest struct {
	ContestId uint `json:"contest_id" form:"contest_id" binding:"required"`
	ProblemId uint `json:"problem_id" form:"contest_id" binding:"required"`
}

type DeleteContestProblemResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}

type JoinContestRequest struct {
	UserId    uint `json:"user_id" form:"user_id" binding:"required"`
	ContestId uint `json:"contest_id" form:"contest_id" binding:"required"`
}

type JoinContestResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}

type ShowContestRequest struct {
	ContestId uint `json:"contest_id" form:"contest_id" binding:"required"`
}

type ShowContestResponse struct {
	Contest      *entity.Contest `json:"contest"`
	BaseResponse *BaseResponse   `json:"base_response"`
}

type UpdateContestRequest struct {
	ContestId   uint   `json:"contest_id" form:"contest_id" binding:"required"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	StartTime   int64  `json:"start_time" form:"start_time" time_format:"unix"`
	EndTime     int64  `json:"end_time" form:"end_time" time_format:"unix"`
}

type UpdateContestResponse struct {
	Contest      *entity.Contest `json:"contest"`
	BaseResponse *BaseResponse   `json:"base_response"`
}

type AddContestAdminRequest struct {
	ContestId uint `json:"contest_id" form:"contest_id" binding:"required"`
	UserId    uint `json:"user_id" form:"user_id" binding:"required"`
}

type AddContestAdminResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}

type DeleteContestAdminRequest struct {
	ContestId uint `json:"contest_id" form:"contest_id" binding:"required"`
	UserId    uint `json:"user_id" form:"user_id" binding:"required"`
}

type DeleteContestAdminResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}

type GenerateParticipantsRequest struct {
	ContestId     uint  `json:"contest_id" form:"contest_id" binding:"required"`
	GenerateCount int32 `json:"generate_count" form:"generate_count" binding:"required"`
}

type GenerateParticipantsResponse struct {
	Users        []*entity.User `json:"users"`
	BaseResponse *BaseResponse  `json:"base_response"`
}

type ContestRankRequest struct {
	ContestId string `json:"contest_id" form:"contest_id" binding:"required"`
}

type ContestRankResponse struct {
	Rank         *entity.Rank  `json:"rank"`
	BaseResponse *BaseResponse `json:"base_response"`
}

type UpdateContestProblemsRequest struct {
	ContestId string                   `json:"contest_id" form:"contest_id" binding:"required"`
	Problems  []*entity.ContestProblem `json:"problems" form:"problems" binding:"required"`
}

type UpdateContestProblemsResponse struct {
	Problems     []*entity.ContestProblem `json:"problems"`
	BaseResponse *BaseResponse            `json:"base_response"`
}

type GetUserContestsRequest struct {
	PageNo   int32 `json:"page_no" form:"page_no" binding:"required"`
	PageSize int32 `json:"page_size" form:"page_size" binding:"required"`
}

type GetUserContestsResponse struct {
	Contests     []*entity.Contest `json:"contests" form:"contests"`
	PageInfo     *entity.Page      `json:"page_info"`
	BaseResponse *BaseResponse     `json:"base_response"`
}
