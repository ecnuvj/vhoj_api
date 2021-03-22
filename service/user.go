package service

import (
	"context"
	"fmt"
	"github.com/ecnuvj/vhoj_api/model/adapter"
	"github.com/ecnuvj/vhoj_api/model/entity"
	"github.com/ecnuvj/vhoj_rpc/client/rpc_user"
	"github.com/ecnuvj/vhoj_rpc/model/base"
	"github.com/ecnuvj/vhoj_rpc/model/userpb"
)

type IUserService interface {
	Login(string, string) (*entity.User, error)
	Register(info *entity.RegisterUserInfo) (*entity.User, error)
	UpdateUserInfo(uint, *entity.User) (*entity.User, error)
	AuthUser(uint, []*entity.Role) (*entity.User, error)
	ListUsers(int32, int32) ([]*entity.User, *entity.Page, error)
	DeleteUser(uint) error
	GetUserById(uint) (*entity.User, error)
	GetUserByUsername(string) (*entity.User, error)
	GetRoleList() ([]*entity.Role, error)
}

var UserService IUserService = &UserServiceImpl{}

type UserServiceImpl struct{}

func (u *UserServiceImpl) Login(username string, password string) (*entity.User, error) {
	request := &userpb.LoginRequest{
		Nickname: username,
		Password: password,
	}
	resp, err := rpc_user.UserServiceClient.Login(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return nil, fmt.Errorf("resp error: %v", err)
	}
	return adapter.RpcUserToEntityUser(resp.User, false), nil
}

func (u *UserServiceImpl) Register(user *entity.RegisterUserInfo) (*entity.User, error) {
	request := &userpb.RegisterRequest{
		User: &userpb.User{
			Username: user.Username,
			Password: user.Password,
			Email:    user.Email,
			School:   user.School,
		},
	}
	resp, err := rpc_user.UserServiceClient.Register(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return nil, fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return adapter.RpcUserToEntityUser(resp.User, false), nil
}

//user id must have
func (u *UserServiceImpl) UpdateUserInfo(userId uint, user *entity.User) (*entity.User, error) {
	request := &userpb.UpdateUserInfoRequest{
		UserId: uint64(userId),
		User:   adapter.EntityUserToRpcUser(user),
	}
	resp, err := rpc_user.UserServiceClient.UpdateUserInfo(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return nil, fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return adapter.RpcUserToEntityUser(resp.User, false), nil
}

func (u *UserServiceImpl) AuthUser(userId uint, roles []*entity.Role) (*entity.User, error) {
	request := &userpb.UpdateUserRolesRequest{
		Roles:  adapter.EntityRolesToRpcRoles(roles),
		UserId: uint64(userId),
	}
	resp, err := rpc_user.UserServiceClient.UpdateUserRoles(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return nil, fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return adapter.RpcUserToEntityUser(resp.User, false), nil
}

func (u *UserServiceImpl) ListUsers(pageNo int32, pageSize int32) ([]*entity.User, *entity.Page, error) {
	request := &userpb.GetAllUsersRequest{
		PageNo:   pageNo,
		PageSize: pageSize,
	}
	resp, err := rpc_user.UserServiceClient.GetAllUsers(context.Background(), request)
	if err != nil {
		return nil, nil, err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return nil, nil, fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return adapter.RpcUsersToEntityUsers(resp.Users, false), &entity.Page{
		TotalCount: resp.TotalCount,
		TotalPages: resp.TotalPages,
	}, nil
}

func (u *UserServiceImpl) DeleteUser(userId uint) error {
	request := &userpb.DeleteUsersRequest{
		UserIds: []uint64{uint64(userId)},
	}
	resp, err := rpc_user.UserServiceClient.DeleteUsers(context.Background(), request)
	if err != nil {
		return err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return nil
}

func (u *UserServiceImpl) GetUserById(userId uint) (*entity.User, error) {
	request := &userpb.GetUserByIdRequest{
		UserId: uint64(userId),
	}
	resp, err := rpc_user.UserServiceClient.GetUserById(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return nil, fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return adapter.RpcUserToEntityUser(resp.User, false), nil
}

func (u *UserServiceImpl) GetUserByUsername(username string) (*entity.User, error) {
	request := &userpb.GetUserByUsernameRequest{
		Username: username,
	}
	resp, err := rpc_user.UserServiceClient.GetUserByUsername(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return nil, fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return adapter.RpcUserToEntityUser(resp.User, false), nil
}

func (u *UserServiceImpl) GetRoleList() ([]*entity.Role, error) {
	request := &userpb.GetRoleListRequest{}
	resp, err := rpc_user.UserServiceClient.GetRoleList(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.BaseResponse.Status != base.REPLY_STATUS_SUCCESS {
		return nil, fmt.Errorf("resp error: %v", resp.BaseResponse.Message)
	}
	return adapter.RpcRolesToEntityRoles(resp.Roles), nil
}
