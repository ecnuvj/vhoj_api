package service

import (
	"context"
	"fmt"
	"github.com/ecnuvj/vhoj_api/model/adapter"
	"github.com/ecnuvj/vhoj_api/model/contract"
	"github.com/ecnuvj/vhoj_api/model/entity"
	"github.com/ecnuvj/vhoj_rpc/client/rpc_submitter"
	"github.com/ecnuvj/vhoj_rpc/model/base"
	"github.com/ecnuvj/vhoj_rpc/model/submitterpb"
)

type ISubmitService interface {
	SubmitCode(*contract.SubmitCodeRequest) (uint, error)
	ReSubmitCode(uint) error
	ListSubmission(int32, int32, *entity.SubmissionSearchCondition) ([]*entity.Submission, *entity.Page, error)
	GetSubmissionCode(uint) (string, error)
}

var SubmitService ISubmitService = &SubmitServiceImpl{}

type SubmitServiceImpl struct{}

func (s *SubmitServiceImpl) SubmitCode(req *contract.SubmitCodeRequest) (uint, error) {
	request := &submitterpb.SubmitCodeRequest{
		ProblemId: uint64(req.ProblemId),
		UserId:    uint64(req.UserId),
		//todo add username
		Language:   req.Language,
		ContestId:  uint64(req.ContestId),
		SourceCode: req.SourceCode,
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
	return adapter.RpcSubmissionsToEntitySubmissions(resp.Submissions), &entity.Page{
		TotalCount: resp.TotalCount,
		TotalPages: resp.TotalPages,
	}, nil
}

func (s *SubmitServiceImpl) GetSubmissionCode(submissionId uint) (string, error) {
	request := &submitterpb.GetSubmissionRequest{
		SubmissionId: uint64(submissionId),
	}
	resp, err := rpc_submitter.SubmitServiceClient.GetSubmission(context.Background(), request)
	if err != nil {
		return "", err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return "", fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return resp.Submission.Code, nil
}
