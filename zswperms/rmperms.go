package zswperms

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewRemoveZswPerms(sender, user zsw.AccountName, permissions zsw.Uint128) *zsw.Action {
	return &zsw.Action{
		Account: AN("zsw.perms"),
		Name:    ActN("rmperms"),
		Authorization: []zsw.PermissionLevel{
			{Actor: sender, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(RemoveZswPerms{
			Sender: sender,
			User: user,
			Permissions: permissions,
		}),
	}
}


type RemoveZswPerms struct {
  Sender zsw.AccountName `json:"sender"`
  User zsw.AccountName `json:"user"`
  Permissions zsw.Uint128 `json:"permissions"`
}