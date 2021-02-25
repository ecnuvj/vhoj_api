package handler

import (
	"github.com/ecnuvj/vhoj_api/model/contract"
	"github.com/ecnuvj/vhoj_api/service"
	"github.com/ecnuvj/vhoj_api/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

func ReSubmitCode(c *gin.Context) {
	request := &contract.ReSubmitCodeRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ReSubmitCodeResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
}

func ListSubmissions(c *gin.Context) {
	request := &contract.ListSubmissionsRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListSubmissionsResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
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

func ShowSubmissionCode(c *gin.Context) {
	//todo 检查权限
	request := &contract.ShowSubmissionCodeRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ShowSubmissionCodeResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
}
