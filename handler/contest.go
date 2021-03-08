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

// @Tags contest
// @Summary 创建比赛
// @Description 创建比赛
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.CreateContestRequest true "request"s
// @Success 200 {object} contract.CreateContestResponse
// @Failure 400 {object} contract.CreateContestResponse
// @Router /contest/create [post]
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

// @Tags contest
// @Summary 比赛列表
// @Description 比赛列表
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.ListContestsRequest true "request"
// @Success 200 {object} contract.ListContestsResponse
// @Failure 400 {object} contract.ListContestsResponse
// @Router /contest/list [get]
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

// @Tags contest
// @Summary 比赛搜索
// @Description 比赛搜索
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.SearchContestRequest true "request"
// @Success 200 {object} contract.SearchContestResponse
// @Failure 400 {object} contract.SearchContestResponse
// @Router /contest/search [post]
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

// @Tags contest
// @Summary 加入比赛
// @Description 加入比赛
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.JoinContestRequest true "request"
// @Success 200 {object} contract.JoinContestResponse
// @Failure 400 {object} contract.JoinContestResponse
// @Router /contest/join [post]
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

// @Tags contest
// @Summary 比赛详情
// @Description 比赛详情
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.ShowContestRequest true "request"
// @Success 200 {object} contract.ShowContestResponse
// @Failure 400 {object} contract.ShowContestResponse
// @Router /contest/show [get]
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

// @Tags contest
// @Summary 比赛更新
// @Description 比赛更新
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.UpdateContestRequest true "request"
// @Success 200 {object} contract.UpdateContestResponse
// @Failure 400 {object} contract.UpdateContestResponse
// @Router /contest/update [post]
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

// @Tags contest
// @Summary 添加比赛题目
// @Description 添加比赛题目
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.AddContestProblemRequest true "request"
// @Success 200 {object} contract.AddContestProblemResponse
// @Failure 400 {object} contract.AddContestProblemResponse
// @Router /contest/problem/add [post]
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

// @Tags contest
// @Summary 删除比赛题目
// @Description 删除比赛题目
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.DeleteContestProblemRequest true "request"
// @Success 200 {object} contract.DeleteContestProblemResponse
// @Failure 400 {object} contract.DeleteContestProblemResponse
// @Router /contest/problem/delete [delete]
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

// @Tags contest
// @Summary 添加比赛管理
// @Description 添加比赛管理
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.AddContestAdminRequest true "request"
// @Success 200 {object} contract.AddContestAdminResponse
// @Failure 400 {object} contract.AddContestAdminResponse
// @Router /contest/admin/add [post]
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

// @Tags contest
// @Summary 删除比赛管理
// @Description 删除比赛管理
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.DeleteContestAdminRequest true "request"
// @Success 200 {object} contract.DeleteContestAdminResponse
// @Failure 400 {object} contract.DeleteContestAdminResponse
// @Router /contest/admin/delete [delete]
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

// @Tags contest
// @Summary 生成比赛用户
// @Description 生成比赛用户
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.GenerateParticipantsRequest true "request"
// @Success 200 {object} contract.GenerateParticipantsResponse
// @Failure 400 {object} contract.GenerateParticipantsResponse
// @Router /contest/generate [post]
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

// @Tags contest
// @Summary 比赛排名
// @Description 比赛排名
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.ContestRankRequest true "request"
// @Success 200 {object} contract.ContestRankResponse
// @Failure 400 {object} contract.ContestRankResponse
// @Router /contest/rank [post]
func ContestRank(c *gin.Context) {
	request := &contract.ContestRankRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ContestRankResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
}
