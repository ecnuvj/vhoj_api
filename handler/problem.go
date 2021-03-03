package handler

import (
	"github.com/ecnuvj/vhoj_api/model/contract"
	"github.com/ecnuvj/vhoj_api/service"
	"github.com/ecnuvj/vhoj_api/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListProblems(c *gin.Context) {
	request := &contract.ListProblemsRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListProblemsResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
	problems, pageInfo, err := service.ProblemService.ListProblems(request.PageNo, request.PageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListProblemsResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.ListProblemsResponse{
		BaseResponse: util.NewSuccessResponse("success"),
		Problems:     problems,
		PageInfo:     pageInfo,
	})
}

func ShowProblem(c *gin.Context) {
	request := &contract.ShowProblemRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ShowProblemResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
	problem, err := service.ProblemService.ShowProblem(request.ProblemId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.ShowProblemResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.ShowProblemResponse{
		BaseResponse: util.NewSuccessResponse("success"),
		Problem:      problem,
	})
}

func SearchProblem(c *gin.Context) {
	request := &contract.SearchProblemRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.SearchProblemResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
	problems, pageInfo, err := service.ProblemService.SearchProblem(request.PageNo, request.PageSize, request.SearchCondition)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.SearchProblemResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.SearchProblemResponse{
		BaseResponse: util.NewSuccessResponse("success"),
		PageInfo:     pageInfo,
		Problems:     problems,
	})
}

func CheckUserProblemStatus(c *gin.Context) {
	request := &contract.CheckUserProblemStatusRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.CheckUserProblemStatusResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
}
