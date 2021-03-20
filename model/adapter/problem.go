package adapter

import (
	"github.com/ecnuvj/vhoj_api/model/entity"
	"github.com/ecnuvj/vhoj_rpc/model/problempb"
	"strconv"
)

func RpcProblemToEntityProblem(problem *problempb.Problem) *entity.Problem {
	if problem == nil {
		return &entity.Problem{}
	}
	return &entity.Problem{
		ProblemId:    uint(problem.ProblemId),
		Title:        problem.Title,
		Description:  problem.Description,
		SampleInput:  problem.SampleInput,
		SampleOutput: problem.SampleOutput,
		Input:        problem.Input,
		Output:       problem.Output,
		TimeLimit:    problem.TimeLimit,
		MemoryLimit:  problem.MemoryLimit,
		Submitted:    problem.Submitted,
		Accepted:     problem.Accepted,
	}
}

func RpcProblemsToEntityProblems(problems []*problempb.Problem) []*entity.Problem {
	retProblems := make([]*entity.Problem, len(problems))
	for i, p := range problems {
		retProblems[i] = RpcProblemToEntityProblem(p)
	}
	return retProblems
}

func RpcContestProblemToEntityProblem(problem *problempb.ContestProblem) *entity.ContestProblem {
	if problem == nil {
		return &entity.ContestProblem{}
	}
	return &entity.ContestProblem{
		ContestId:    strconv.Itoa(int(problem.ContestId)),
		ProblemId:    strconv.Itoa(int(problem.ProblemId)),
		ProblemOrder: problem.ProblemOrder,
		Title:        problem.Title,
		Accepted:     uint(problem.Accepted),
		Submitted:    uint(problem.Submitted),
	}
}

func RpcContestProblemsToEntityContestProblems(problems []*problempb.ContestProblem) []*entity.ContestProblem {
	retProblems := make([]*entity.ContestProblem, len(problems))
	for i, p := range problems {
		retProblems[i] = RpcContestProblemToEntityProblem(p)
	}
	return retProblems
}

func EntityContestProblemToRpcContestProblem(problem *entity.ContestProblem) *problempb.ContestProblem {
	if problem == nil {
		return &problempb.ContestProblem{}
	}
	contestId, _ := strconv.Atoi(problem.ContestId)
	problemId, _ := strconv.Atoi(problem.ProblemId)
	return &problempb.ContestProblem{
		ContestId:    uint64(contestId),
		ProblemId:    uint64(problemId),
		ProblemOrder: problem.ProblemOrder,
		Title:        problem.Title,
		Accepted:     uint64(problem.Accepted),
		Submitted:    uint64(problem.Submitted),
	}
}

func EntityContestProblemsToRpcContestProblems(problems []*entity.ContestProblem) []*problempb.ContestProblem {
	retProblems := make([]*problempb.ContestProblem, len(problems))
	for i, p := range problems {
		retProblems[i] = EntityContestProblemToRpcContestProblem(p)
	}
	return retProblems
}
