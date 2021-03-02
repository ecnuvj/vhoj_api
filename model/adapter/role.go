package adapter

import (
	"github.com/ecnuvj/vhoj_api/model/entity"
	"github.com/ecnuvj/vhoj_rpc/model/userpb"
)

func RpcRolesToEntityRoles(roles []*userpb.Role) []*entity.Role {
	retRoles := make([]*entity.Role, len(roles))
	for i, c := range roles {
		retRoles[i].RoleId = uint(c.RoleId)
		retRoles[i].RoleName = c.RoleName
	}
	return retRoles
}
