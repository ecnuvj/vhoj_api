package handler

import (
	"github.com/ecnuvj/vhoj_api/model/contract"
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

}

func ShowProblem(c *gin.Context) {
	request := &contract.ShowProblemRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ShowProblemResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
}

func SearchProblem(c *gin.Context) {
	request := &contract.SearchProblemRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.SearchProblemResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
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
