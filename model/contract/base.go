package contract

type BaseRequest struct {
}

type BaseResponse struct {
	StatusCode StatusCode  `json:"status_code"`
	StatusInfo *StatusInfo `json:"status_info"`
}

type StatusInfo struct {
	Time    int64  `json:"time"`
	Message string `json:"message"`
}

type StatusCode int32

const (
	SUCCESS StatusCode = 1
	FAILURE StatusCode = 2
	UNKNOWN StatusCode = 3
)
