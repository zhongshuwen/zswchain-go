package zswperms

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewAddZswPerms(sender, user zsw.AccountName, permissions zsw.Uint128) *zsw.Action {
	return &zsw.Action{
		Account: AN("zsw.perms"),
		Name:    ActN("addperms"),
		Authorization: []zsw.PermissionLevel{
			{Actor: sender, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(AddZswPerms{
			Sender: sender,
			User: user,
			Permissions: permissions,
		}),
	}
}


type AddZswPerms struct {
  Sender zsw.AccountName `json:"sender"`
  User zsw.AccountName `json:"user"`
  Permissions zsw.Uint128 `json:"permissions"`
}