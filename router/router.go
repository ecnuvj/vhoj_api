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
		problem.POST("/search", handler.SearchProblem)
	}

	contest := router.Group("/contest")
	{
		contest.POST("/create", handler.CreateContest)
		contest.GET("/list", handler.ListContests)
		contest.GET("/show", handler.ShowContest)
		contest.POST("/search", handler.SearchContest)
		contest.POST("/problem/add", handler.AddContestProblem)
		contest.DELETE("/problem/delete", handler.DeleteContestProblem)
		contest.POST("/admin/add", handler.AddContestAdmin)
		contest.DELETE("/admin/delete", handler.DeleteContestAdmin)
		contest.POST("/join", handler.JoinContest)
		contest.POST("/update", handler.UpdateContest)
		contest.POST("/generate", handler.GenerateParticipants)
		contest.POST("/rank", handler.ContestRank)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("router run error: %v", err)
	}
}
