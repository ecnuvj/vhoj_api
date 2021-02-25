package handler

import (
	"github.com/ecnuvj/vhoj_api/model/contract"
	"github.com/ecnuvj/vhoj_api/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateContest(c *gin.Context) {
	request := &contract.CreateContestRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.CreateContestResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
}

func ListContests(c *gin.Context) {
	request := &contract.ListContestsRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListContestsResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
}

func JoinContest(c *gin.Context) {
	request := &contract.JoinContestRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.JoinContestResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
}

func ShowContest(c *gin.Context) {
	request := &contract.ShowContestRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ShowContestResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
}

func UpdateContest(c *gin.Context) {
	//todo 检查是否有权限修改
	request := &contract.UpdateContestRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.UpdateContestResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
}

func AddContestAdmin(c *gin.Context) {
	//todo 检查是否有权限修改
	request := &contract.AddContestAdminRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.AddContestAdminResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
}

func GenerateParticipants(c *gin.Context) {
	//todo 检查是否有权限生成
	request := &contract.GenerateParticipantsRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.GenerateParticipantsResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
}

func ContestRank(c *gin.Context) {
	request := &contract.ContestRankRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ContestRankResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
}
