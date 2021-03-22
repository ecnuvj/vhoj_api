package adapter

import (
	"github.com/ecnuvj/vhoj_api/model/entity"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/contest_status"
	"github.com/ecnuvj/vhoj_rpc/model/problempb"
	"github.com/golang/protobuf/ptypes"
	"time"
)

//not contain problem entity
func EntityContestToRpcContest(contest *entity.Contest) *problempb.Contest {
	if contest == nil {
		return &problempb.Contest{}
	}
	startTime, _ := ptypes.TimestampProto(contest.StartTime)
	endTime, _ := ptypes.TimestampProto(contest.EndTime)
	problemIds := make([]uint64, len(contest.ProblemIds))
	for i, p := range contest.ProblemIds {
		problemIds[i] = uint64(p)
	}
	return &problempb.Contest{
		ContestId:   uint64(contest.ContestId),
		Title:       contest.Title,
		Description: contest.Description,
		StartTime:   startTime,
		EndTime:     endTime,
		Creator:     EntityUserToRpcUser(contest.Creator),
		ProblemIds:  problemIds,
	}
}

func RpcContestToEntityContest(contest *problempb.Contest) *entity.Contest {
	if contest == nil {
		return &entity.Contest{}
	}
	problemIds := make([]uint, len(contest.ProblemIds))
	for i, p := range contest.ProblemIds {
		problemIds[i] = uint(p)
	}
	var status contest_status.ContestStatus
	now := time.Now()
	if now.Before(contest.StartTime.AsTime()) {
		status = contest_status.SCHEDULED
	} else if now.After(contest.EndTime.AsTime()) {
		status = contest_status.ENDED
	} else {
		status = contest_status.RUNNING
	}
	return &entity.Contest{
		ContestId:    uint(contest.ContestId),
		Title:        contest.Title,
		Description:  contest.Description,
		StartTime:    contest.StartTime.AsTime(),
		EndTime:      contest.EndTime.AsTime(),
		Creator:      RpcUserToEntityUser(contest.Creator, false),
		ProblemIds:   problemIds,
		Problems:     RpcContestProblemsToEntityContestProblems(contest.Problems),
		Status:       status,
		ProblemCount: len(contest.ProblemIds),
	}
}

func RpcContestsToEntityContests(contests []*problempb.Contest) []*entity.Contest {
	retContests := make([]*entity.Contest, len(contests))
	for i, c := range contests {
		retContests[i] = RpcContestToEntityContest(c)
	}
	return retContests
}
