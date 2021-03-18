package rpc_client

import (
	"github.com/ecnuvj/vhoj_rpc/client/rpc_problem"
	"github.com/ecnuvj/vhoj_rpc/client/rpc_submitter"
	"github.com/ecnuvj/vhoj_rpc/client/rpc_user"
)

func Init() {
	rpc_submitter.Init()
	rpc_user.Init()
	rpc_problem.Init()
}
