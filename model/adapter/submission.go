package adapter

import (
	"github.com/ecnuvj/vhoj_api/model/entity"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/language"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/status_type"
	"github.com/ecnuvj/vhoj_rpc/model/submitterpb"
)

func RpcSubmissionToEntitySubmission(submission *submitterpb.Submission) *entity.Submission {
	if submission == nil {
		return &entity.Submission{}
	}
	return &entity.Submission{
		SubmissionId: uint(submission.SubmissionId),
		Username:     submission.Username,
		ProblemId:    uint(submission.ProblemId),
		UserId:       uint(submission.UserId),
		Result:       status_type.SubmissionStatusType(submission.Result),
		TimeCost:     submission.TimeCost,
		MemoryCost:   submission.MemoryCost,
		Language:     language.Language(submission.Language),
	}
}

func RpcSubmissionsToEntitySubmissions(submissions []*submitterpb.Submission) []*entity.Submission {
	retSubmissions := make([]*entity.Submission, len(submissions))
	for i, s := range submissions {
		retSubmissions[i] = RpcSubmissionToEntitySubmission(s)
	}
	return retSubmissions
}
