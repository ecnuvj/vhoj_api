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

type IProblemService interface {
	ListProblems(int32, int32) ([]*entity.Problem, *entity.Page, error)
	ShowProblem(uint) (*entity.Problem, error)
	SearchProblem(int32, int32, *entity.SearchCondition) ([]*entity.Problem, error)
}

var ProblemService IProblemService = &ProblemServiceImpl{}

type ProblemServiceImpl struct{}

func (p *ProblemServiceImpl) ListProblems(pageNo int32, pageSize int32) ([]*entity.Problem, *entity.Page, error) {
	request := &problempb.ListProblemsRequest{
		PageNo:   pageNo,
		PageSize: pageSize,
	}
	resp, err := rpc_problem.ProblemServiceClient.ListProblems(context.Background(), request)
	if err != nil {
		return nil, nil, err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return nil, nil, fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return adapter.RpcProblemsToEntityProblems(resp.Problems), &entity.Page{
		TotalCount: resp.TotalCount,
		TotalPages: resp.TotalPages,
	}, nil
}

func (p *ProblemServiceImpl) ShowProblem(problemId uint) (*entity.Problem, error) {
	request := &problempb.GetProblemByIdRequest{
		ProblemId: uint64(problemId),
	}
	resp, err := rpc_problem.ProblemServiceClient.GetProblemById(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return nil, fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return adapter.RpcProblemToEntityProblem(resp.Problem), nil
}

func (p *ProblemServiceImpl) SearchProblem(pageNo int32, pageSize int32, condition *entity.SearchCondition) ([]*entity.Problem, error) {
	request := &problempb.SearchProblemByConditionRequest{
		Title:     condition.Title,
		ProblemId: uint64(condition.ProblemId),
		PageNo:    pageNo,
		PageSize:  pageSize,
	}
	resp, err := rpc_problem.ProblemServiceClient.SearchProblemByCondition(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return nil, fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return adapter.RpcProblemsToEntityProblems(resp.Problems), nil
}
