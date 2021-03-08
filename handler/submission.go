package handler

import (
	"github.com/ecnuvj/vhoj_api/model/contract"
	"github.com/ecnuvj/vhoj_api/model/entity"
	"github.com/ecnuvj/vhoj_api/service"
	"github.com/ecnuvj/vhoj_api/util"
	"github.com/gin-gonic/gin"
	"net/http"
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
	submitId, err := service.SubmitService.SubmitCode(request)
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
	submissions, pageInfo, err := service.SubmitService.ListSubmission(request.PageNo, request.PageSize, &entity.SubmissionSearchCondition{
		Username:  request.Username,
		ProblemId: request.ProblemId,
		Status:    request.Status,
		Language:  request.Language,
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

func CheckSubmissionStatus(c *gin.Context) {
	request := &contract.CheckSubmissionStatusRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.CheckSubmissionStatusResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
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
// @Router /status/list [post]
func ShowSubmissionCode(c *gin.Context) {
	//todo 检查权限
	request := &contract.ShowSubmissionCodeRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ShowSubmissionCodeResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	code, err := service.SubmitService.GetSubmissionCode(request.SubmissionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.ShowSubmissionCodeResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.ShowSubmissionCodeResponse{
		SubmissionCode: code,
		BaseResponse:   util.NewSuccessResponse("success"),
	})
}
