package bootstrap

import (
	"github.com/ecnuvj/vhoj_api/bootstrap/rpc_client"
	"github.com/ecnuvj/vhoj_api/router"
)

func Init() {
	rpc_client.Init()
	router.Init()
}
