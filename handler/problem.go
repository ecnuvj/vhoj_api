package handler

import (
	"github.com/ecnuvj/vhoj_api/auth"
	"github.com/ecnuvj/vhoj_api/model/contract"
	"github.com/ecnuvj/vhoj_api/model/entity"
	"github.com/ecnuvj/vhoj_api/service"
	"github.com/ecnuvj/vhoj_api/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Tags problem
// @Summary 题目列表
// @Description 题目列表
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.ListProblemsRequest true "request"
// @Success 200 {object} contract.ListProblemsResponse
// @Failure 400 {object} contract.ListProblemsResponse
// @Router /problem/list [post]
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

// @Tags problem
// @Summary 题目详情
// @Description 题目详情
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.ShowProblemRequest true "request"
// @Success 200 {object} contract.ShowProblemResponse
// @Failure 400 {object} contract.ShowProblemResponse
// @Router /problem/show [get]
func ShowProblem(c *gin.Context) {
	request := &contract.ShowProblemRequest{}
	if err := c.ShouldBindQuery(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ShowProblemResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
	problemId, err := util.StringToNumber(request.ProblemId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.ShowProblemResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
	problem, err := service.ProblemService.ShowProblem(uint(problemId))
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

// @Tags problem
// @Summary 题目搜索
// @Description 题目搜索
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.SearchProblemRequest true "request"
// @Success 200 {object} contract.SearchProblemResponse
// @Failure 400 {object} contract.SearchProblemResponse
// @Router /problem/search [post]
func SearchProblem(c *gin.Context) {
	request := &contract.SearchProblemRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.SearchProblemResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
	problemId, err := util.StringToNumber(request.ProblemId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.SearchProblemResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
	problems, pageInfo, err := service.ProblemService.SearchProblem(request.PageNo, request.PageSize, &entity.ProblemSearchCondition{
		Title:     request.Title,
		ProblemId: uint(problemId),
	})
	token, exist := c.Get("auth")
	if exist {
		userId, _ := util.StringToNumber(token.(*auth.Claims).UserId)
		problems, _ = service.ProblemService.CheckUserProblemsStatus(problems, uint(userId), 0)
	}
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

// @Tags problem
// @Summary 随机提米
// @Description 随机题目
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.RandProblemRequest true "request"
// @Success 200 {object} contract.RandProblemResponse
// @Failure 400 {object} contract.RandProblemResponse
// @Router /problem/rand [get]
func RandProblem(c *gin.Context) {
	problemId, err := service.ProblemService.RandProblem()
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.RandProblemResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.RandProblemResponse{
		BaseResponse: util.NewSuccessResponse("success"),
		ProblemId:    problemId,
	})
}

func RawProblemList(c *gin.Context) {
	request := &contract.RawProblemListRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.RawProblemListResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
	problems, pageInfo, err := service.ProblemService.RawProblemList(request.PageNo, request.PageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.RawProblemListResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.RawProblemListResponse{
		BaseResponse: util.NewSuccessResponse("success"),
		RawProblems:  problems,
		PageInfo:     pageInfo,
	})
}

func CrawlProblem(c *gin.Context) {
	request := &contract.CrawlProblemRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.CrawlProblemResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
	rawId, err := service.ProblemService.CrawlProblem(request.RemoteOJ, request.RemoteProblemId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.CrawlProblemResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.CrawlProblemResponse{
		BaseResponse: util.NewSuccessResponse("success"),
		RawProblemId: rawId,
	})
}

func QueryCrawl(c *gin.Context) {
	rawIdStr := c.Param("rawId")
	rawId, err := util.StringToNumber(rawIdStr)
	if err != nil || rawId <= 0 {
		c.JSON(http.StatusBadRequest, &contract.QueryCrawlResponse{
			BaseResponse: util.NewFailureResponse("param error: %v", err),
		})
		return
	}
	status, err := service.ProblemService.QueryCrawl(uint(rawId))
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.QueryCrawlResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.QueryCrawlResponse{
		BaseResponse: util.NewSuccessResponse("success"),
		Status:       status,
	})
}
