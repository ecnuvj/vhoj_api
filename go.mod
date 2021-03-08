module github.com/ecnuvj/vhoj_api

go 1.14

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/ecnuvj/vhoj_common v0.0.0-00010101000000-000000000000
	github.com/ecnuvj/vhoj_rpc v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.3
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.5.1
)

replace github.com/ecnuvj/vhoj_rpc => ../vhoj_rpc

replace github.com/ecnuvj/vhoj_common => ../vhoj_common
