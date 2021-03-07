package contract

import (
	"github.com/ecnuvj/vhoj_api/model/entity"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/language"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/status_type"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/submission_status"
)

type SubmitCodeRequest struct {
	UserId     uint   `form:"user_id" json:"user_id" binding:"required"`
	ProblemId  uint   `form:"problem_id" json:"problem_id" binding:"required"`
	ContestId  uint   `form:"contest_id" json:"contest_id"`
	SourceCode string `form:"source_code" json:"source_code" binding:"required"`
	Language   int32  `form:"language" json:"language" binding:"required"`
}

type SubmitCodeResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}

type ReSubmitCodeRequest struct {
	SubmissionId uint `json:"submission_id" form:"submission_id" binding:"required"`
}

type ReSubmitCodeResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}

type ListSubmissionsRequest struct {
	Username  string                           `json:"username" form:"username"`
	ProblemId uint                             `json:"problem_id" form:"problem_id"`
	Status    status_type.SubmissionStatusType `json:"status" form:"status"`
	Language  language.Language                `json:"language" form:"language"`
	PageNo    int32                            `json:"page_no" form:"page_no" binding:"required"`
	PageSize  int32                            `json:"page_size" form:"page_size" binding:"required"`
}

type ListSubmissionsResponse struct {
	BaseResponse *BaseResponse        `json:"base_response"`
	Submissions  []*entity.Submission `json:"submissions"`
	PageInfo     *entity.Page         `json:"page_info"`
}

type CheckSubmissionStatusRequest struct {
	SubmissionId uint `json:"submission_id" form:"submission_id"`
}

type CheckSubmissionStatusResponse struct {
	SubmissionStatus *submission_status.SubmissionStatus `json:"submission_status"`
	BaseResponse     *BaseResponse                       `json:"base_response"`
}

type ShowSubmissionCodeRequest struct {
	SubmissionId uint `json:"submission_id" form:"submission_id" binding:"required"`
}

type ShowSubmissionCodeResponse struct {
	SubmissionCode string        `json:"submission_code"`
	BaseResponse   *BaseResponse `json:"base_response"`
}
