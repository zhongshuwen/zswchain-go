package zswperms

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewSetZswPerms(sender, user, scope zsw.AccountName, permissions zsw.Uint128) *zsw.Action {
	return &zsw.Action{
		Account: AN("zsw.perms"),
		Name:    ActN("setperms"),
		Authorization: []zsw.PermissionLevel{
			{Actor: sender, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(SetZswPerms{
			Sender: sender,
			Scope: scope,
			User: user,
			Permissions: permissions,
		}),
	}
}


type SetZswPerms struct {
  Sender zsw.AccountName `json:"sender"`
  Scope zsw.AccountName `json:"scope"`
  User zsw.AccountName `json:"user"`
  Permissions zsw.Uint128 `json:"perm_bits"`
}