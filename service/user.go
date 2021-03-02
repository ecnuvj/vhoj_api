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
	Register(*entity.User) (*entity.User, error)
	UpdateUserInfo(*entity.User) (*entity.User, error)
	AuthUser(uint, []*entity.Role) (*entity.User, error)
}

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
	return adapter.RpcUserToEntityUser(resp.User), nil
}

func (u *UserServiceImpl) Register(user *entity.User) (*entity.User, error) {
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
		return nil, fmt.Errorf("resp error: %v", err)
	}
	return adapter.RpcUserToEntityUser(resp.User), nil
}

func (u *UserServiceImpl) UpdateUserInfo(user *entity.User) (*entity.User, error) {
	request := &userpb
}

func (u *UserServiceImpl) AuthUser(u2 uint, roles []*entity.Role) (*entity.User, error) {
	panic("implement me")
}
