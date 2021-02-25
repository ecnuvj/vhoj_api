package handler

import (
	"github.com/ecnuvj/vhoj_api/model/contract"
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
}

func AuthUser(c *gin.Context) {
	request := &contract.AuthUserRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.AuthUserResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
}

func DeleteUser(c *gin.Context) {
	request := &contract.DeleteUserRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.DeleteUserResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
}

func Login(c *gin.Context) {
	request := &contract.LoginRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.LoginResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
}

func Register(c *gin.Context) {
	request := &contract.RegisterRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.RegisterResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
}

func UpdateUserInfo(c *gin.Context) {
	request := &contract.UpdateUserInfoRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.UpdateUserInfoResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
}
