package handler

import (
	"github.com/ecnuvj/vhoj_api/auth"
	"github.com/ecnuvj/vhoj_api/model/contract"
	"github.com/ecnuvj/vhoj_api/model/entity"
	"github.com/ecnuvj/vhoj_api/service"
	"github.com/ecnuvj/vhoj_api/util"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/language"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/role"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/status_type"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Tags problem
// @Summary 提交代码
// @Description 提交代码
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.SubmitCodeRequest true "request"
// @Success 200 {object} contract.SubmitCodeResponse
// @Failure 400 {object} contract.SubmitCodeResponse
// @Router /problem/submit [post]
func SubmitCode(c *gin.Context) {
	request := &contract.SubmitCodeRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.SubmitCodeResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	problemId, err := util.StringToNumber(request.ProblemId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.SubmitCodeResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	lang, err := util.StringToNumber(request.Language)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.SubmitCodeResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	contestId, err := util.StringToNumber(request.ContestId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.SubmitCodeResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	token, exist := c.Get("auth")
	if !exist {
		c.JSON(http.StatusUnauthorized, &contract.UserInfoResponse{
			BaseResponse: util.NewFailureResponse("auth token not exist"),
		})
		return
	}
	claims := token.(*auth.Claims)
	userId, _ := strconv.Atoi(claims.UserId)
	param := &service.SubmitCodeParam{
		ProblemId:  uint64(problemId),
		UserId:     uint64(userId),
		Username:   claims.Username,
		Language:   int32(lang),
		ContestId:  uint64(contestId),
		SourceCode: request.SourceCode,
	}
	submitId, err := service.SubmitService.SubmitCode(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.SubmitCodeResponse{
			BaseResponse: util.NewFailureResponse("submit code error: %v", err),
		})
		return
	}
	if submitId <= 0 {
		c.JSON(http.StatusBadRequest, &contract.SubmitCodeResponse{
			BaseResponse: util.NewFailureResponse("submit id error"),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.SubmitCodeResponse{
		BaseResponse: util.NewSuccessResponse("submit success, submit id: %v", submitId),
	})
}

// @Tags status
// @Summary 重新提交代码
// @Description 重新提交代码
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.ReSubmitCodeRequest true "request"
// @Success 200 {object} contract.ReSubmitCodeResponse
// @Failure 400 {object} contract.ReSubmitCodeResponse
// @Router /status/reSubmit [post]
func ReSubmitCode(c *gin.Context) {
	request := &contract.ReSubmitCodeRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ReSubmitCodeResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	err := service.SubmitService.ReSubmitCode(request.SubmissionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.ReSubmitCodeResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.ReSubmitCodeResponse{
		BaseResponse: util.NewSuccessResponse("success"),
	})
}

// @Tags status
// @Summary 提交列表
// @Description 提交列表
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.ListSubmissionsRequest true "request"
// @Success 200 {object} contract.ListSubmissionsResponse
// @Failure 400 {object} contract.ListSubmissionsResponse
// @Router /status/list [post]
func ListSubmissions(c *gin.Context) {
	request := &contract.ListSubmissionsRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListSubmissionsResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	problemId, err := util.StringToNumber(request.ProblemId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListSubmissionsResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	status, err := util.StringToNumber(request.Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListSubmissionsResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	lang, err := util.StringToNumber(request.Language)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListSubmissionsResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	submissions, pageInfo, err := service.SubmitService.ListSubmission(request.PageNo, request.PageSize, &entity.SubmissionSearchCondition{
		Username:  request.Username,
		ProblemId: uint(problemId),
		Status:    status_type.SubmissionStatusType(status),
		Language:  language.Language(lang),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListSubmissionsResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.ListSubmissionsResponse{
		BaseResponse: util.NewSuccessResponse("success"),
		Submissions:  submissions,
		PageInfo:     pageInfo,
	})
}

// @Tags status
// @Summary 查看代码
// @Description 查看代码
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.ShowSubmissionCodeRequest true "request"
// @Success 200 {object} contract.ShowSubmissionCodeResponse
// @Failure 400 {object} contract.ShowSubmissionCodeResponse
// @Router /status/showCode [post]
func ShowSubmissionCode(c *gin.Context) {
	//todo 检查权限
	request := &contract.ShowSubmissionCodeRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ShowSubmissionCodeResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	submissionId, err := util.StringToNumber(request.SubmissionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.ShowSubmissionCodeResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	submission, err := service.SubmitService.GetSubmission(uint(submissionId))
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.ShowSubmissionCodeResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	token, exist := c.Get("auth")
	if !exist {
		c.JSON(http.StatusUnauthorized, &contract.ShowSubmissionCodeResponse{
			BaseResponse: util.NewFailureResponse("auth token not exist"),
		})
		return
	}
	user := token.(*auth.Claims)
	userId, _ := util.StringToNumber(user.UserId)
	var ok = true
	if userId != int(submission.UserId) {
		ok = false
		for _, r := range user.Roles {
			if r == role.CODE_REVIEWER || r == role.ADMIN_USER {
				ok = true
				break
			}
		}
	}
	if ok {
		c.JSON(http.StatusOK, &contract.ShowSubmissionCodeResponse{
			Submission:   submission,
			BaseResponse: util.NewSuccessResponse("success"),
		})
	} else {
		c.JSON(http.StatusForbidden, &contract.ShowSubmissionCodeResponse{
			BaseResponse: util.NewFailureResponse("forbidden"),
		})
	}

}
