package service

import (
	"context"
	"fmt"
	"github.com/ecnuvj/vhoj_api/model/adapter"
	"github.com/ecnuvj/vhoj_api/model/entity"
	"github.com/ecnuvj/vhoj_rpc/client/rpc_submitter"
	"github.com/ecnuvj/vhoj_rpc/model/base"
	"github.com/ecnuvj/vhoj_rpc/model/submitterpb"
)

type SubmitCodeParam struct {
	ProblemId  uint64
	UserId     uint64
	Username   string
	Language   int32
	ContestId  uint64
	SourceCode string
}

type ISubmitService interface {
	SubmitCode(*SubmitCodeParam) (uint, error)
	ReSubmitCode(uint) error
	ListSubmission(int32, int32, *entity.SubmissionSearchCondition) ([]*entity.Submission, *entity.Page, error)
	GetSubmission(uint) (*entity.Submission, error)
}

var SubmitService ISubmitService = &SubmitServiceImpl{}

type SubmitServiceImpl struct{}

func (s *SubmitServiceImpl) SubmitCode(param *SubmitCodeParam) (uint, error) {
	request := &submitterpb.SubmitCodeRequest{
		ProblemId:  param.ProblemId,
		UserId:     param.UserId,
		Username:   param.Username,
		Language:   param.Language,
		ContestId:  param.ContestId,
		SourceCode: param.SourceCode,
	}
	resp, err := rpc_submitter.SubmitServiceClient.SubmitCode(context.Background(), request)
	if err != nil {
		return 0, fmt.Errorf("submit service err: %v", err)
	}
	if resp == nil {
		return 0, fmt.Errorf("submit service resp is nil")
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return 0, fmt.Errorf(resp.BaseResponse.Message)
	}
	return uint(resp.SubmissionId), nil
}

func (s *SubmitServiceImpl) ReSubmitCode(submissionId uint) error {
	request := &submitterpb.ReSubmitCodeRequest{
		SubmissionId: uint64(submissionId),
	}
	resp, err := rpc_submitter.SubmitServiceClient.ReSubmitCode(context.Background(), request)
	if err != nil {
		return err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return nil
}

func (s *SubmitServiceImpl) ListSubmission(pageNo int32, pageSize int32, condition *entity.SubmissionSearchCondition) ([]*entity.Submission, *entity.Page, error) {
	request := &submitterpb.ListSubmissionsRequest{
		Username:  condition.Username,
		ProblemId: uint64(condition.ProblemId),
		Result:    int32(condition.Status),
		Language:  int32(condition.Language),
		PageNo:    pageNo,
		PageSize:  pageSize,
	}
	resp, err := rpc_submitter.SubmitServiceClient.ListSubmissions(context.Background(), request)
	if err != nil {
		return nil, nil, err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return nil, nil, fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return adapter.RpcSubmissionsToEntitySubmissions(resp.Submissions, false), &entity.Page{
		TotalCount: resp.TotalCount,
		TotalPages: resp.TotalPages,
	}, nil
}

func (s *SubmitServiceImpl) GetSubmission(submissionId uint) (*entity.Submission, error) {
	request := &submitterpb.GetSubmissionRequest{
		SubmissionId: uint64(submissionId),
	}
	resp, err := rpc_submitter.SubmitServiceClient.GetSubmission(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return nil, fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return adapter.RpcSubmissionToEntitySubmission(resp.Submission, true), nil
}
