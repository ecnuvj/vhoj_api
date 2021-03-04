package service

import (
	"context"
	"fmt"
	"github.com/ecnuvj/vhoj_api/model/adapter"
	"github.com/ecnuvj/vhoj_api/model/entity"
	"github.com/ecnuvj/vhoj_rpc/client/rpc_problem"
	"github.com/ecnuvj/vhoj_rpc/model/base"
	"github.com/ecnuvj/vhoj_rpc/model/problempb"
)

type IContestService interface {
	CreateContest(*entity.Contest) (*entity.Contest, error)
	ListContests(int32, int32) ([]*entity.Contest, *entity.Page, error)
	ShowContest(uint) (*entity.Contest, error)
	SearchContest(int32, int32, *entity.ContestSearchCondition) ([]*entity.Contest, *entity.Page, error)
}

var ContestService IContestService = &ContestServiceImpl{}

type ContestServiceImpl struct{}

func (c *ContestServiceImpl) CreateContest(contest *entity.Contest) (*entity.Contest, error) {
	request := &problempb.CreateContestRequest{
		Contest: adapter.EntityContestToRpcContest(contest),
	}
	resp, err := rpc_problem.ProblemServiceClient.CreateContest(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return nil, fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return adapter.RpcContestToEntityContest(resp.Contest), nil
}

func (c *ContestServiceImpl) ListContests(pageNo int32, pageSize int32) ([]*entity.Contest, *entity.Page, error) {
	request := &problempb.ListContestsRequest{
		PageNo:   pageNo,
		PageSize: pageSize,
	}
	resp, err := rpc_problem.ProblemServiceClient.ListContests(context.Background(), request)
	if err != nil {
		return nil, nil, err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return nil, nil, fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return adapter.RpcContestsToEntityContests(resp.Contests), &entity.Page{
		TotalCount: resp.TotalCount,
		TotalPages: resp.TotalPages,
	}, nil
}

func (c *ContestServiceImpl) ShowContest(contestId uint) (*entity.Contest, error) {
	request := &problempb.GetContestByIdRequest{
		ContestId: uint64(contestId),
	}
	resp, err := rpc_problem.ProblemServiceClient.GetContestById(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return nil, fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return adapter.RpcContestToEntityContest(resp.Contest), nil
}

func (c *ContestServiceImpl) SearchContest(pageNo int32, pageSize int32, condition *entity.ContestSearchCondition) ([]*entity.Contest, *entity.Page, error) {
	request := &problempb.SearchContestByConditionRequest{
		Status:      int32(condition.Status),
		Title:       condition.Title,
		CreatorName: condition.CreatorName,
		PageNo:      pageNo,
		PageSize:    pageSize,
	}
	resp, err := rpc_problem.ProblemServiceClient.SearchContestByCondition(context.Background(), request)
	if err != nil {
		return nil, nil, err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return nil, nil, fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return adapter.RpcContestsToEntityContests(resp.Contests), &entity.Page{
		TotalCount: resp.TotalCount,
		TotalPages: resp.TotalPages,
	}, nil
}
