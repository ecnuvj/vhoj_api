package contract

import "github.com/ecnuvj/vhoj_api/model/entity"

type ListUsersRequest struct {
	PageNo   int32 `json:"page_no" form:"page_no" binding:"required"`
	PageSize int32 `json:"page_size" form:"page_size" binding:"required"`
}

type ListUsersResponse struct {
	Users        []*entity.User `json:"users"`
	PageInfo     *entity.Page   `json:"page_info"`
	BaseResponse *BaseResponse  `json:"base_response"`
}

type AuthUserRequest struct {
	UserId uint           `json:"user_id" form:"user_id" binding:"required"`
	Roles  []*entity.Role `json:"roles" form:"roles" binding:"required"`
}

type AuthUserResponse struct {
	User         *entity.User  `json:"user"`
	BaseResponse *BaseResponse `json:"base_response"`
}

type DeleteUserRequest struct {
	UserId uint `json:"user_id" form:"user_id" binding:"required"`
}

type DeleteUserResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}

type LoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginResponse struct {
	User         *entity.User  `json:"user"`
	BaseResponse *BaseResponse `json:"base_response"`
}

type RegisterRequest struct {
	User *entity.RegisterUserInfo `json:"user" form:"user" binding:"required"`
}

type RegisterResponse struct {
	User         *entity.User  `json:"user"`
	BaseResponse *BaseResponse `json:"base_response"`
}

type UpdateUserInfoRequest struct {
	UserId uint         `json:"user_id" form:"user_id" binding:"required"`
	User   *entity.User `json:"user" form:"user" binding:"required"`
}

type UpdateUserInfoResponse struct {
	User         *entity.User  `json:"user"`
	BaseResponse *BaseResponse `json:"base_response"`
}

type UserInfoRequest struct {
}

type UserInfoResponse struct {
	User         *entity.User  `json:"user"`
	BaseResponse *BaseResponse `json:"base_response"`
}
