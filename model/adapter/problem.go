package adapter

import (
	"github.com/ecnuvj/vhoj_api/model/entity"
	"github.com/ecnuvj/vhoj_rpc/model/problempb"
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
