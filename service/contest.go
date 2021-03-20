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
	JoinContest(uint, uint) error
	UpdateContest(uint, *entity.Contest) (*entity.Contest, error)
	AddContestProblem(uint, uint) error
	DeleteContestProblem(uint, uint) error
	AddContestAdmin(uint, uint) error
	DeleteContestAdmin(uint, uint) error
	GenerateParticipants(uint, int32) ([]*entity.User, error)
	ContestRank(uint) (*entity.Rank, error)
	UpdateContestProblems(uint, []*entity.ContestProblem) ([]*entity.ContestProblem, error)
}

var ContestService IContestService = &ContestServiceImpl{}

type ContestServiceImpl struct{}

func (c *ContestServiceImpl) CreateContest(contest *entity.Contest) (*entity.Contest, error) {
	request := &problempb.CreateContestRequest{
		Contest:         adapter.EntityContestToRpcContest(contest),
		ContestProblems: adapter.EntityContestProblemsToRpcContestProblems(contest.Problems),
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

func (c *ContestServiceImpl) JoinContest(contestId uint, userId uint) error {
	request := &problempb.AddContestParticipantRequest{
		ContestId: uint64(contestId),
		UserId:    uint64(userId),
	}
	resp, err := rpc_problem.ProblemServiceClient.AddContestParticipant(context.Background(), request)
	if err != nil {
		return err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return nil
}

func (c *ContestServiceImpl) UpdateContest(contestId uint, contest *entity.Contest) (*entity.Contest, error) {
	request := &problempb.UpdateContestRequest{
		ContestId: uint64(contestId),
		Contest:   adapter.EntityContestToRpcContest(contest),
	}
	resp, err := rpc_problem.ProblemServiceClient.UpdateContest(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return nil, fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return adapter.RpcContestToEntityContest(resp.Contest), nil
}

func (c *ContestServiceImpl) AddContestProblem(contestId uint, problemId uint) error {
	request := &problempb.AddContestProblemRequest{
		ContestId: uint64(contestId),
		ProblemId: uint64(problemId),
	}
	resp, err := rpc_problem.ProblemServiceClient.AddContestProblem(context.Background(), request)
	if err != nil {
		return err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return nil
}

func (c *ContestServiceImpl) DeleteContestProblem(contestId uint, problemId uint) error {
	request := &problempb.DeleteContestProblemRequest{
		ContestId: uint64(contestId),
		ProblemId: uint64(problemId),
	}
	resp, err := rpc_problem.ProblemServiceClient.DeleteContestProblem(context.Background(), request)
	if err != nil {
		return err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return nil
}

func (c *ContestServiceImpl) AddContestAdmin(contestId uint, userId uint) error {
	request := &problempb.AddContestAdminRequest{
		ContestId: uint64(contestId),
		UserId:    uint64(userId),
	}
	resp, err := rpc_problem.ProblemServiceClient.AddContestAdmin(context.Background(), request)
	if err != nil {
		return err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return nil
}

func (c *ContestServiceImpl) DeleteContestAdmin(contestId uint, userId uint) error {
	request := &problempb.DeleteContestAdminRequest{
		ContestId: uint64(contestId),
		UserId:    uint64(userId),
	}
	resp, err := rpc_problem.ProblemServiceClient.DeleteContestAdmin(context.Background(), request)
	if err != nil {
		return err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return nil
}

func (c *ContestServiceImpl) GenerateParticipants(contestId uint, count int32) ([]*entity.User, error) {
	request := &problempb.GenerateContestParticipantsRequest{
		ContestId:     uint64(contestId),
		GenerateCount: count,
	}
	resp, err := rpc_problem.ProblemServiceClient.GenerateContestParticipants(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return nil, fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return adapter.RpcUsersToEntityUsers(resp.Users), nil
}

func (c *ContestServiceImpl) ContestRank(contestId uint) (*entity.Rank, error) {
	//submitterResp, err := rpc_submitter.SubmitServiceClient.GetContestSubmissions(context.Background(), &submitterpb.GetContestSubmissionsRequest{
	//	ContestId: uint64(contestId),
	//})
	//if err != nil {
	//	return nil, err
	//}
	//submissions := adapter.RpcSubmissionsToEntitySubmissions(submitterResp.Submissions, false)
	//problemResp, err := rpc_problem.ProblemServiceClient.GetContestProblems(context.Background(), &problempb.GetContestProblemsRequest{
	//	ContestId: uint64(contestId),
	//})
	//if err != nil {
	//	return nil, err
	//}
	//problems := adapter.
	panic("no")
}

func (c *ContestServiceImpl) UpdateContestProblems(contestId uint, problems []*entity.ContestProblem) ([]*entity.ContestProblem, error) {
	request := &problempb.UpdateContestProblemsRequest{
		ContestId:       uint64(contestId),
		ContestProblems: adapter.EntityContestProblemsToRpcContestProblems(problems),
	}
	resp, err := rpc_problem.ProblemServiceClient.UpdateContestProblems(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return nil, fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return adapter.RpcContestProblemsToEntityContestProblems(resp.ContestProblems), nil
}
