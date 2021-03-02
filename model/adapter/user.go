package adapter

import (
	"github.com/ecnuvj/vhoj_api/model/entity"
	"github.com/ecnuvj/vhoj_rpc/model/userpb"
)

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
