module github.com/ecnuvj/vhoj_api

go 1.14

require (
	github.com/ecnuvj/vhoj_rpc v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
)

replace github.com/ecnuvj/vhoj_rpc => ../vhoj_rpc

replace github.com/ecnuvj/vhoj_common => ../vhoj_common
