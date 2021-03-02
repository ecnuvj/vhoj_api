package adapter

import (
	"github.com/ecnuvj/vhoj_api/model/entity"
	"github.com/ecnuvj/vhoj_rpc/model/userpb"
)

func RpcRolesToEntityRoles(roles []*userpb.Role) []*entity.Role {
	retRoles := make([]*entity.Role, len(roles))
	for i, c := range roles {
		retRoles[i] = &entity.Role{
			RoleId:   uint(c.RoleId),
			RoleName: c.RoleName,
		}
	}
	return retRoles
}

func EntityRolesToRpcRoles(roles []*entity.Role) []*userpb.Role {
	retRoles := make([]*userpb.Role, len(roles))
	for i, c := range roles {
		retRoles[i] = &userpb.Role{
			RoleId:   uint64(c.RoleId),
			RoleName: c.RoleName,
		}
	}
	return retRoles
}
