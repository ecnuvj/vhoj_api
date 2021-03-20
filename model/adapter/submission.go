package adapter

import (
	"github.com/ecnuvj/vhoj_api/model/entity"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/language"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/status_type"
	"github.com/ecnuvj/vhoj_rpc/model/submitterpb"
)

func RpcSubmissionToEntitySubmission(submission *submitterpb.Submission, withCode bool) *entity.Submission {
	if submission == nil {
		return &entity.Submission{}
	}
	result := entity.SubmissionResult{
		Code: status_type.SubmissionStatusType(submission.Result),
		Text: status_type.CodeToTextMap[status_type.SubmissionStatusType(submission.Result)],
	}
	lang := entity.SubmissionLanguage{
		Code: language.Language(submission.Language),
		Text: language.CodeToTextMap[language.Language(submission.Language)],
	}
	retSubmission := &entity.Submission{
		SubmissionId: uint(submission.SubmissionId),
		Username:     submission.Username,
		ProblemId:    uint(submission.ProblemId),
		UserId:       uint(submission.UserId),
		Result:       result,
		TimeCost:     submission.TimeCost,
		MemoryCost:   submission.MemoryCost,
		Language:     lang,
		SubmitTime:   submission.SubmitTime.AsTime(),
	}
	if withCode {
		retSubmission.Code = submission.Code
	}
	return retSubmission
}

func RpcSubmissionsToEntitySubmissions(submissions []*submitterpb.Submission, withCode bool) []*entity.Submission {
	retSubmissions := make([]*entity.Submission, len(submissions))
	for i, s := range submissions {
		retSubmissions[i] = RpcSubmissionToEntitySubmission(s, withCode)
	}
	return retSubmissions
}
