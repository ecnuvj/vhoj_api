package router

import (
	"github.com/ecnuvj/vhoj_api/handler"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()
	router.POST("/problem/submit", handler.SubmitCode)
}
