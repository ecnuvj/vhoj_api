package service

import (
	"context"
	"fmt"
	"github.com/ecnuvj/vhoj_api/model/adapter"
	"github.com/ecnuvj/vhoj_api/model/entity"
	"github.com/ecnuvj/vhoj_api/util"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/status_type"
	"github.com/ecnuvj/vhoj_rpc/client/rpc_problem"
	"github.com/ecnuvj/vhoj_rpc/client/rpc_submitter"
	"github.com/ecnuvj/vhoj_rpc/model/base"
	"github.com/ecnuvj/vhoj_rpc/model/problempb"
	"github.com/ecnuvj/vhoj_rpc/model/submitterpb"
	"sort"
	"time"
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
	ContestRank(uint, time.Time) (*entity.Rank, error)
	UpdateContestProblems(uint, []*entity.ContestProblem) ([]*entity.ContestProblem, error)
	GetUserContests(userId uint, pageNo int32, pageSize int32) ([]*entity.Contest, *entity.Page, error)
	GetContestParticipants(contestId uint) ([]*entity.User, error)
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
	retContest := adapter.RpcContestToEntityContest(resp.Contest)
	retContest.ProblemCount = len(retContest.Problems)
	return retContest, nil
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
	return adapter.RpcUsersToEntityUsers(resp.Users, true), nil
}

type tmpStatus struct {
	First    bool
	Accepted bool
	WaCount  int
	time     time.Time
}

func (c *ContestServiceImpl) ContestRank(contestId uint, startTime time.Time) (*entity.Rank, error) {
	submitterResp, err := rpc_submitter.SubmitServiceClient.GetContestSubmissions(context.Background(), &submitterpb.GetContestSubmissionsRequest{
		ContestId: uint64(contestId),
	})
	if err != nil {
		return nil, err
	}
	submissions := adapter.RpcSubmissionsToEntitySubmissions(submitterResp.Submissions, false)
	sort.Sort(entity.Submissions(submissions))
	userSubmissions := make(map[uint][]*entity.Submission)
	problemSubmissions := make(map[uint][]*entity.Submission)
	for _, s := range submissions {
		userSubmissions[s.UserId] = append(userSubmissions[s.UserId], s)
		problemSubmissions[s.ProblemId] = append(problemSubmissions[s.ProblemId], s)
	}
	problemResp, err := rpc_problem.ProblemServiceClient.GetContestProblems(context.Background(), &problempb.GetContestProblemsRequest{
		ContestId: uint64(contestId),
	})
	if err != nil {
		return nil, err
	}
	problems := adapter.RpcContestProblemsToEntityContestProblems(problemResp.ContestProblems)
	userResp, err := rpc_problem.ProblemServiceClient.GetContestParticipants(context.Background(), &problempb.GetContestParticipantsRequest{
		ContestId: uint64(contestId),
	})
	if err != nil {
		return nil, err
	}
	users := adapter.RpcUsersToEntityUsers(userResp.Users, false)
	userProblemMap := make(map[uint]map[uint]*tmpStatus)
	userScoreMap := make(map[uint]uint)
	userPenaltyMap := make(map[uint]uint)
	for _, u := range users {
		userProblemMap[u.UserId] = make(map[uint]*tmpStatus)
		for _, p := range problems {
			problemId, _ := util.StringToNumber(p.ProblemId)
			userProblemMap[u.UserId][uint(problemId)] = &tmpStatus{
				First:    false,
				Accepted: false,
				WaCount:  0,
				time:     time.Time{},
			}
		}
		for _, s := range userSubmissions[u.UserId] {
			if s.Result.Code == status_type.AC {
				if userProblemMap[u.UserId][s.ProblemId].Accepted {
					continue
				}
				userProblemMap[u.UserId][s.ProblemId].Accepted = true
				userProblemMap[u.UserId][s.ProblemId].time = s.SubmitTime
				userScoreMap[u.UserId]++
				userPenaltyMap[u.UserId] += uint(userProblemMap[u.UserId][s.ProblemId].WaCount*20 + int(s.SubmitTime.Sub(startTime).Minutes()))
			} else {
				if userProblemMap[u.UserId][s.ProblemId].Accepted {
					continue
				}
				userProblemMap[u.UserId][s.ProblemId].WaCount++
			}
		}
	}
	for _, sub := range problemSubmissions {
		for _, s := range sub {
			if s.Result.Code == status_type.AC {
				userProblemMap[s.UserId][s.ProblemId].First = true
				break
			}
		}
	}
	participants := make(entity.Participants, len(users))
	for i, u := range users {
		participants[i] = &entity.Participant{
			Username: u.Username,
			Rank:     0,
			Score:    userScoreMap[u.UserId],
			Penalty:  userPenaltyMap[u.UserId],
			Problems: make(map[string]*entity.ProblemStatus),
		}
		for _, pro := range problems {
			problemId, _ := util.StringToNumber(pro.ProblemId)
			var submitCount = userProblemMap[u.UserId][uint(problemId)].WaCount
			var status uint = 3
			if submitCount != 0 {
				status = 4
			}
			var submitTime string
			if userProblemMap[u.UserId][uint(problemId)].Accepted {
				status = 2
				submitCount++
				t := int(userProblemMap[u.UserId][uint(problemId)].time.Sub(startTime).Seconds())
				submitTime = fmt.Sprintf("%v:%v:%v", t/3600, (t%3600)/60, t%60)
			}
			if userProblemMap[u.UserId][uint(problemId)].First {
				status = 1
			}
			participants[i].Problems[pro.ProblemOrder] = &entity.ProblemStatus{
				ProblemId:   uint(problemId),
				SubmitCount: uint(submitCount),
				SubmitTime:  submitTime,
				Status:      status,
			}
		}
	}
	sort.Sort(participants)
	sort.Sort(entity.ContestProblems(problems))
	for i, p := range participants {
		p.Rank = uint(i + 1)
	}
	return &entity.Rank{
		Problems:     problems,
		Participants: participants,
	}, nil
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

func (c *ContestServiceImpl) GetUserContests(userId uint, pageNo int32, pageSize int32) ([]*entity.Contest, *entity.Page, error) {
	request := &problempb.GetUserContestsRequest{
		UserId:   uint64(userId),
		PageNo:   pageNo,
		PageSize: pageSize,
	}
	resp, err := rpc_problem.ProblemServiceClient.GetUserContests(context.Background(), request)
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

func (c *ContestServiceImpl) GetContestParticipants(contestId uint) ([]*entity.User, error) {
	request := &problempb.GetContestParticipantsRequest{
		ContestId: uint64(contestId),
	}
	resp, err := rpc_problem.ProblemServiceClient.GetContestParticipants(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return adapter.RpcUsersToEntityUsers(resp.Users, true), nil
}
