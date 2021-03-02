package handler

import (
	"github.com/ecnuvj/vhoj_api/model/contract"
	"github.com/ecnuvj/vhoj_api/service"
	"github.com/ecnuvj/vhoj_api/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUsers(c *gin.Context) {
	request := &contract.ListUsersRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListUsersResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	users, pageInfo, err := service.UserService.ListUsers(request.PageNo, request.PageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListUsersResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.ListUsersResponse{
		Users:        users,
		PageInfo:     pageInfo,
		BaseResponse: util.NewSuccessResponse("success"),
	})
}

func AuthUser(c *gin.Context) {
	request := &contract.AuthUserRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.AuthUserResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	util.CheckRequest(request)
	user, err := service.UserService.AuthUser(request.UserId, request.Roles)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.AuthUserResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.AuthUserResponse{
		User:         user,
		BaseResponse: util.NewSuccessResponse("auth success"),
	})
}

func DeleteUser(c *gin.Context) {
	request := &contract.DeleteUserRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.DeleteUserResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	err := service.UserService.DeleteUser(request.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.DeleteUserResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.DeleteUserResponse{
		BaseResponse: util.NewSuccessResponse("delete success"),
	})
}

func Login(c *gin.Context) {
	request := &contract.LoginRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.LoginResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	user, err := service.UserService.Login(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.LoginResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.LoginResponse{
		User:         user,
		BaseResponse: util.NewSuccessResponse("login success"),
	})
}

func Register(c *gin.Context) {
	request := &contract.RegisterRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.RegisterResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	user, err := service.UserService.Register(request.User)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.RegisterResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.RegisterResponse{
		User:         user,
		BaseResponse: util.NewSuccessResponse("register success"),
	})
}

func UpdateUserInfo(c *gin.Context) {
	request := &contract.UpdateUserInfoRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.UpdateUserInfoResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	user, err := service.UserService.UpdateUserInfo(request.UserId, request.User)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.UpdateUserInfoResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.UpdateUserInfoResponse{
		User:         user,
		BaseResponse: util.NewSuccessResponse("update success"),
	})
}
