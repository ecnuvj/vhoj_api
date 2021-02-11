package service

import (
	"context"
	"fmt"
	"github.com/ecnuvj/vhoj_api/model/contract"
	"github.com/ecnuvj/vhoj_rpc/client/rpc_submitter"
	"github.com/ecnuvj/vhoj_rpc/model/base"
	"github.com/ecnuvj/vhoj_rpc/model/submitterpb"
)

type ISubmitService interface {
	SubmitCode(*contract.SubmitCodeRequest) (uint, error)
	ReSubmitCode(uint) error
}

var SubmitService ISubmitService = &SubmitServiceImpl{}

type SubmitServiceImpl struct{}

func (s *SubmitServiceImpl) SubmitCode(req *contract.SubmitCodeRequest) (uint, error) {
	request := &submitterpb.SubmitCodeRequest{
		ProblemId:  uint64(req.ProblemId),
		UserId:     uint64(req.UserId),
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

func (s *SubmitServiceImpl) ReSubmitCode(u uint) error {
	panic("implement me")
}
