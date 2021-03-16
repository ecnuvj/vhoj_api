package handler

import (
	"github.com/ecnuvj/vhoj_api/auth"
	"github.com/ecnuvj/vhoj_api/model/contract"
	"github.com/ecnuvj/vhoj_api/service"
	"github.com/ecnuvj/vhoj_api/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Tags user
// @Summary 用户列表
// @Description 分页用户
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.ListUsersRequest true "list user request"
// @Success 200 {object} contract.ListUsersResponse
// @Failure 400 {object} contract.ListUsersResponse
// @Router /user/list [get]
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

// @Tags user
// @Summary 授权用户
// @Description 授权用户
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.AuthUserRequest true "auth user request"
// @Success 200 {object} contract.AuthUserResponse
// @Failure 400 {object} contract.AuthUserResponse
// @Router /user/auth [post]
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

// @Tags user
// @Summary 删除用户
// @Description 删除用户
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.DeleteUserRequest true "delete user request"
// @Success 200 {object} contract.DeleteUserResponse
// @Failure 400 {object} contract.DeleteUserResponse
// @Router /user/delete [delete]
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

// @Tags user
// @Summary 登录
// @Description 登录
// @Accept  json
// @Produce json
// @Param   request body contract.LoginRequest true "request"
// @Success 200 {object} contract.LoginResponse
// @Failure 400 {object} contract.LoginResponse
// @Router /user/login [post]
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
	token, _ := auth.GenerateToken(user)
	user.Token = token
	c.JSON(http.StatusOK, &contract.LoginResponse{
		User:         user,
		BaseResponse: util.NewSuccessResponse("login success"),
	})
}

// @Tags user
// @Summary 注册
// @Description 注册
// @Accept  json
// @Produce json
// @Param   request body contract.RegisterRequest true "request"
// @Success 200 {object} contract.RegisterResponse
// @Failure 400 {object} contract.RegisterResponse
// @Router /user/register [post]
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

// @Tags user
// @Summary 更新信息
// @Description 更新信息
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.UpdateUserInfoRequest true "request"
// @Success 200 {object} contract.UpdateUserInfoResponse
// @Failure 400 {object} contract.UpdateUserInfoResponse
// @Router /user/register [post]
func UpdateUserInfo(c *gin.Context) {
	request := &contract.UpdateUserInfoRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.UpdateUserInfoResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	//str, _ := json.Marshal(request.User)
	//fmt.Println(string(str))
	//return
	token, exist := c.Get("auth")
	if !exist {
		c.JSON(http.StatusUnauthorized, &contract.UserInfoResponse{
			BaseResponse: util.NewFailureResponse("auth token not exist"),
		})
		return
	}
	claims := token.(*auth.Claims)
	if claims.UserId != strconv.Itoa(int(request.UserId)) {
		c.JSON(http.StatusUnauthorized, &contract.UserInfoResponse{
			BaseResponse: util.NewFailureResponse("token not match"),
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
	authToken, _ := auth.GenerateToken(user)
	user.Token = authToken
	c.JSON(http.StatusOK, &contract.UpdateUserInfoResponse{
		User:         user,
		BaseResponse: util.NewSuccessResponse("update success"),
	})
}

// @Tags user
// @Summary 用户信息
// @Description 用户信息
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.UserInfoRequest true "request"
// @Success 200 {object} contract.UserInfoResponse
// @Failure 400 {object} contract.UserInfoResponse
// @Router /user/info [post]
func UserInfo(c *gin.Context) {
	request := &contract.UserInfoRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.UserInfoResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	token, exist := c.Get("auth")
	if !exist {
		c.JSON(http.StatusUnauthorized, &contract.UserInfoResponse{
			BaseResponse: util.NewFailureResponse("auth token not exist"),
		})
		return
	}
	claims := token.(*auth.Claims)
	userId, _ := strconv.Atoi(claims.UserId)
	user, err := service.UserService.GetUserById(uint(userId))
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.UserInfoResponse{
			BaseResponse: util.NewFailureResponse("service error: %v", err),
		})
		return
	}
	authToken, _ := auth.GenerateToken(user)
	user.Token = authToken
	c.JSON(http.StatusOK, &contract.UserInfoResponse{
		User:         user,
		BaseResponse: util.NewSuccessResponse("success"),
	})
}
