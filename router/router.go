package router

import (
	"github.com/ecnuvj/vhoj_api/handler"
	"github.com/gin-gonic/gin"
	"log"
)

func Init() {
	router := gin.Default()
	router.POST("/problem/submit", handler.SubmitCode)

	user := router.Group("/user")
	{
		user.POST("/login", handler.Login)
		user.POST("/register", handler.Register)
		user.POST("/update", handler.UpdateUserInfo)
		//todo auth
		user.POST("/auth", handler.AuthUser)
		user.GET("/list", handler.ListUsers)
		user.DELETE("/delete", handler.DeleteUser)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("router run error: %v", err)
	}
}
