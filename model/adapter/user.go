package adapter

import (
	"github.com/ecnuvj/vhoj_api/model/entity"
	"github.com/ecnuvj/vhoj_rpc/model/userpb"
)

//no password
func RpcUserToEntityUser(user *userpb.User) *entity.User {
	return &entity.User{
		UserId:     uint(user.UserId),
		Username:   user.Username,
		Email:      user.Email,
		School:     user.School,
		UserAuthId: uint(user.UserAuthId),
		Roles:      RpcRolesToEntityRoles(user.Roles),
		Accepted:   user.Accepted,
		Submitted:  user.Submitted,
	}
}

func EntityUserToRpcUser(user *entity.User) *userpb.User {
	return &userpb.User{
		UserId:     uint64(user.UserId),
		Username:   user.Username,
		UserAuthId: uint64(user.UserAuthId),
		Password:   user.Password,
		Email:      user.Email,
		School:     user.School,
		Roles:      EntityRolesToRpcRoles(user.Roles),
		Accepted:   user.Accepted,
		Submitted:  user.Submitted,
	}
}

func RpcUsersToEntityUsers(users []*userpb.User) []*entity.User {
	retUsers := make([]*entity.User, len(users))
	for i, u := range users {
		retUsers[i] = RpcUserToEntityUser(u)
	}
	return retUsers
}
