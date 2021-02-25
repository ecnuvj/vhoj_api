package router

import (
	"github.com/ecnuvj/vhoj_api/handler"
	"github.com/gin-gonic/gin"
	"log"
)

func Init() {
	router := gin.Default()
	router.POST("/problem/submit", handler.SubmitCode)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("router run error: %v", err)
	}
}
