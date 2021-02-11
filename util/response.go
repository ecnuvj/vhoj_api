package util

import (
	"fmt"
	"github.com/ecnuvj/vhoj_api/model/contract"
	"time"
)

func NewSuccessResponse(format string, args ...interface{}) *contract.BaseResponse {
	statusInfo := &contract.StatusInfo{
		Time:    time.Now().Unix(),
		Message: fmt.Sprintf(format, args...),
	}
	return &contract.BaseResponse{
		StatusCode: contract.SUCCESS,
		StatusInfo: statusInfo,
	}
}

func NewFailureResponse(format string, args ...interface{}) *contract.BaseResponse {
	statusInfo := &contract.StatusInfo{
		Time:    time.Now().Unix(),
		Message: fmt.Sprintf(format, args...),
	}
	return &contract.BaseResponse{
		StatusCode: contract.FAILURE,
		StatusInfo: statusInfo,
	}
}
