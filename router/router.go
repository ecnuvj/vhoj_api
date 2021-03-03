package router

import (
	"github.com/ecnuvj/vhoj_api/handler"
	"github.com/gin-gonic/gin"
	"log"
)

func Init() {
	router := gin.Default()

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

	problem := router.Group("/problem")
	{
		problem.GET("/list", handler.ListProblems)
		problem.GET("/show", handler.ShowProblem)
		problem.POST("/submit", handler.SubmitCode)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("router run error: %v", err)
	}
}
