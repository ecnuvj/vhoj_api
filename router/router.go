package router

import (
	_ "github.com/ecnuvj/vhoj_api/docs"
	"github.com/ecnuvj/vhoj_api/handler"
	"github.com/ecnuvj/vhoj_api/router/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func Init() {
	router := gin.Default()
	router.Use(middleware.Cors())
	user := router.Group("/user")
	{
		user.POST("/login", handler.Login)
		user.POST("/register", handler.Register)
		user.POST("/update", middleware.Auth(), handler.UpdateUserInfo)
		user.POST("/info", middleware.Auth(), handler.UserInfo)
		user.POST("/auth", middleware.Auth(), middleware.RoleCheck("admin"), handler.AuthUser)
		user.POST("/list", middleware.Auth(), middleware.RoleCheck("admin"), handler.ListUsers)
		user.DELETE("/delete", middleware.Auth(), middleware.RoleCheck("admin"), handler.DeleteUser)
		user.GET("/roles", handler.RoleList)
	}

	problem := router.Group("/problem")
	{
		problem.POST("/list", middleware.Auth(), handler.ListProblems)
		problem.GET("/show", handler.ShowProblem)
		problem.POST("/submit", middleware.Auth(), handler.SubmitCode)
		problem.POST("/search", middleware.Auth(), handler.SearchProblem)
	}

	contest := router.Group("/contest")
	{
		contest.POST("/create", middleware.Auth(), middleware.RoleCheck("creator"), handler.CreateContest)
		contest.GET("/list", handler.ListContests)
		contest.GET("/show", middleware.Auth(), handler.ShowContest)
		contest.POST("/search", handler.SearchContest)
		contest.POST("/problem/add", middleware.Auth(), handler.AddContestProblem)
		contest.DELETE("/problem/delete", middleware.Auth(), handler.DeleteContestProblem)
		contest.POST("/admin/add", middleware.Auth(), handler.AddContestAdmin)
		contest.DELETE("/admin/delete", middleware.Auth(), handler.DeleteContestAdmin)
		contest.POST("/join", middleware.Auth(), handler.JoinContest)
		contest.POST("/update", middleware.Auth(), handler.UpdateContest)
		contest.POST("/generate", middleware.Auth(), handler.GenerateParticipants)
		contest.POST("/rank", middleware.Auth(), handler.ContestRank)
	}

	status := router.Group("/status")
	{
		status.POST("/list", handler.ListSubmissions)
		status.POST("/showCode", middleware.Auth(), middleware.RoleCheck("reviewer"), handler.ShowSubmissionCode)
		status.POST("/reSubmit", middleware.Auth(), handler.ReSubmitCode)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("router run error: %v", err)
	}
}
