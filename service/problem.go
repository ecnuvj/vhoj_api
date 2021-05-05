package service

import (
	"context"
	"fmt"
	"github.com/ecnuvj/vhoj_api/model/adapter"
	"github.com/ecnuvj/vhoj_api/model/entity"
	"github.com/ecnuvj/vhoj_api/util"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/user_problem_status"
	"github.com/ecnuvj/vhoj_rpc/client/rpc_problem"
	"github.com/ecnuvj/vhoj_rpc/client/rpc_remote"
	"github.com/ecnuvj/vhoj_rpc/client/rpc_submitter"
	"github.com/ecnuvj/vhoj_rpc/model/base"
	"github.com/ecnuvj/vhoj_rpc/model/problempb"
	"github.com/ecnuvj/vhoj_rpc/model/remotepb"
	"github.com/ecnuvj/vhoj_rpc/model/submitterpb"
)

type IProblemService interface {
	ListProblems(int32, int32) ([]*entity.Problem, *entity.Page, error)
	ShowProblem(uint) (*entity.Problem, error)
	SearchProblem(int32, int32, *entity.ProblemSearchCondition) ([]*entity.Problem, *entity.Page, error)
	CheckUserProblemsStatus(problems []*entity.Problem, userId uint, contestId uint) ([]*entity.Problem, error)
	CheckUserContestProblemsStatus(problems []*entity.ContestProblem, userId uint, contestId uint) ([]*entity.ContestProblem, error)
	RandProblem() (uint, error)
	RawProblemList(pageNo int32, pageSize int32) ([]*entity.RawProblem, *entity.Page, error)
	CrawlProblem(remoteOj remote_oj.RemoteOJ, problemId string) (uint, error)
	QueryCrawl(rawId uint) (int32, error)
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

func (p *ProblemServiceImpl) SearchProblem(pageNo int32, pageSize int32, condition *entity.ProblemSearchCondition) ([]*entity.Problem, *entity.Page, error) {
	request := &problempb.SearchProblemByConditionRequest{
		Title:     condition.Title,
		ProblemId: uint64(condition.ProblemId),
		PageNo:    pageNo,
		PageSize:  pageSize,
	}
	resp, err := rpc_problem.ProblemServiceClient.SearchProblemByCondition(context.Background(), request)
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

func (p *ProblemServiceImpl) CheckUserProblemsStatus(problems []*entity.Problem, userId uint, contestId uint) ([]*entity.Problem, error) {
	for _, p := range problems {
		p.Status, _ = checkUserProblemStatus(userId, p.ProblemId, contestId)
	}
	return problems, nil
}

func (p *ProblemServiceImpl) CheckUserContestProblemsStatus(problems []*entity.ContestProblem, userId uint, contestId uint) ([]*entity.ContestProblem, error) {
	for _, p := range problems {
		problemId, _ := util.StringToNumber(p.ProblemId)
		p.Status, _ = checkUserProblemStatus(userId, uint(problemId), contestId)
	}
	return problems, nil
}

func checkUserProblemStatus(userId uint, problemId uint, contestId uint) (user_problem_status.UserProblemStatus, error) {
	request := &submitterpb.CheckUserProblemStatusRequest{
		UserId:    uint64(userId),
		ProblemId: uint64(problemId),
		ContestId: uint64(contestId),
	}
	resp, err := rpc_submitter.SubmitServiceClient.CheckUserProblemStatus(context.Background(), request)
	if err != nil {
		return user_problem_status.UNKNOWN, err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return user_problem_status.UNKNOWN, fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return user_problem_status.UserProblemStatus(resp.Status), nil
}

func (p *ProblemServiceImpl) RandProblem() (uint, error) {
	resp, err := rpc_problem.ProblemServiceClient.RandProblem(context.Background(), &problempb.RandProblemRequest{})
	if err != nil {
		return 0, err
	}
	return uint(resp.ProblemId), nil
}

func (p *ProblemServiceImpl) RawProblemList(pageNo int32, pageSize int32) ([]*entity.RawProblem, *entity.Page, error) {
	resp, err := rpc_problem.ProblemServiceClient.RawProblemList(context.Background(), &problempb.RawProblemListRequest{
		PageNo:   pageNo,
		PageSize: pageSize,
	})
	if err != nil {
		return nil, nil, err
	}
	return adapter.RpcRawProblemsToEntityRawProblems(resp.RawProblems), &entity.Page{
		TotalCount: resp.TotalCount,
		TotalPages: resp.TotalPages,
	}, nil
}

func (p *ProblemServiceImpl) CrawlProblem(remoteOj remote_oj.RemoteOJ, problemId string) (uint, error) {
	resp, err := rpc_remote.RemoteServiceClient.CrawlProblem(context.Background(), &remotepb.CrawlProblemRequest{
		RemoteOj:        int32(remoteOj),
		RemoteProblemId: problemId,
		Enforce:         true,
	})
	if err != nil {
		return 0, err
	}
	if resp.BaseResponse == nil {
		return 0, fmt.Errorf("rpc remote error, no resp")
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return 0, fmt.Errorf(resp.BaseResponse.Message)
	}
	return uint(resp.RawProblemId), nil
}

func (p *ProblemServiceImpl) QueryCrawl(rawId uint) (int32, error) {
	resp, err := rpc_remote.RemoteServiceClient.QueryCrawlResult(context.Background(), &remotepb.QueryCrawlResultRequest{
		RawId: uint64(rawId),
	})
	if err != nil {
		return 0, err
	}
	if resp.BaseResponse == nil {
		return 0, fmt.Errorf("rpc remote error, no resp")
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return 0, fmt.Errorf(resp.BaseResponse.Message)
	}
	return resp.Status, nil
}
