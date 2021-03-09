package main

import (
	"github.com/ecnuvj/vhoj_api/bootstrap"
	_ "github.com/ecnuvj/vhoj_common/pkg/common/constants/status_type"
	_ "github.com/ecnuvj/vhoj_common/pkg/common/constants/submission_status"
)

func main() {
	bootstrap.Init()
}
