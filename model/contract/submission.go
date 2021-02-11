package contract

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
}

type ReSubmitCodeResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}

type ListSubmissionsRequest struct {
}

type ListSubmissionsResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}
