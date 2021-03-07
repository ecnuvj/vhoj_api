package handler

import (
	"github.com/ecnuvj/vhoj_api/model/contract"
	"github.com/ecnuvj/vhoj_api/model/entity"
	"github.com/ecnuvj/vhoj_api/service"
	"github.com/ecnuvj/vhoj_api/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateContest(c *gin.Context) {
	request := &contract.CreateContestRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.CreateContestResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
	//todo real userid
	user, err := service.UserService.GetUserById(9)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.CreateContestResponse{
			BaseResponse: util.NewFailureResponse("find creator error: %v", err),
		})
		return
	}
	//creator 只需id
	contest := &entity.Contest{
		Title:       request.Title,
		Description: request.Description,
		StartTime:   time.Unix(request.StartTime, 0),
		EndTime:     time.Unix(request.EndTime, 0),
		Creator:     user,
		ProblemIds:  request.ProblemIds,
	}
	retContest, err := service.ContestService.CreateContest(contest)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.CreateContestResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.CreateContestResponse{
		Contest:      retContest,
		BaseResponse: util.NewSuccessResponse("success"),
	})
}

func ListContests(c *gin.Context) {
	request := &contract.ListContestsRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListContestsResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
	contests, pageInfo, err := service.ContestService.ListContests(request.PageNo, request.PageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListContestsResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.ListContestsResponse{
		BaseResponse: util.NewSuccessResponse("success"),
		Contests:     contests,
		PageInfo:     pageInfo,
	})
}

func SearchContest(c *gin.Context) {
	request := &contract.SearchContestRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.SearchContestResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
	contests, pageInfo, err := service.ContestService.SearchContest(request.PageNo, request.PageSize, request.SearchCondition)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.SearchContestResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.SearchContestResponse{
		Contests:     contests,
		BaseResponse: util.NewSuccessResponse("success"),
		PageInfo:     pageInfo,
	})
}

func JoinContest(c *gin.Context) {
	request := &contract.JoinContestRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.JoinContestResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
	err := service.ContestService.JoinContest(request.ContestId, request.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.JoinContestResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.JoinContestResponse{
		BaseResponse: util.NewSuccessResponse("success"),
	})
}

func ShowContest(c *gin.Context) {
	request := &contract.ShowContestRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ShowContestResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
	contest, err := service.ContestService.ShowContest(request.ContestId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.ShowContestResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.ShowContestResponse{
		Contest:      contest,
		BaseResponse: util.NewSuccessResponse("success"),
	})
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
	contest := &entity.Contest{
		ContestId:   request.ContestId,
		Title:       request.Title,
		Description: request.Description,
		StartTime:   time.Unix(request.StartTime, 0),
		EndTime:     time.Unix(request.EndTime, 0),
	}
	retContest, err := service.ContestService.UpdateContest(request.ContestId, contest)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.UpdateContestResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.UpdateContestResponse{
		Contest:      retContest,
		BaseResponse: util.NewSuccessResponse("success"),
	})
}

func AddContestProblem(c *gin.Context) {
	request := &contract.AddContestProblemRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.AddContestProblemResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
	err := service.ContestService.AddContestProblem(request.ContestId, request.ProblemId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.AddContestProblemResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.AddContestProblemResponse{
		BaseResponse: util.NewSuccessResponse("success"),
	})
}

func DeleteContestProblem(c *gin.Context) {
	request := &contract.DeleteContestProblemRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.DeleteContestProblemResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
	err := service.ContestService.DeleteContestProblem(request.ContestId, request.ProblemId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.DeleteContestProblemResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.DeleteContestProblemResponse{
		BaseResponse: util.NewSuccessResponse("success"),
	})
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
	err := service.ContestService.AddContestAdmin(request.ContestId, request.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.AddContestAdminResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.AddContestAdminResponse{
		BaseResponse: util.NewSuccessResponse("success"),
	})
}

func DeleteContestAdmin(c *gin.Context) {
	//todo 检查是否有权限修改
	request := &contract.DeleteContestAdminRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.DeleteContestAdminResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
	err := service.ContestService.DeleteContestAdmin(request.ContestId, request.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.DeleteContestAdminResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.DeleteContestAdminResponse{
		BaseResponse: util.NewSuccessResponse("success"),
	})
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
	users, err := service.ContestService.GenerateParticipants(request.ContestId, request.GenerateCount)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.GenerateParticipantsResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.GenerateParticipantsResponse{
		Users:        users,
		BaseResponse: util.NewSuccessResponse("success"),
	})
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
