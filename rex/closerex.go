package rex

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewCloseREX(owner zsw.AccountName) *zsw.Action {
	return &zsw.Action{
		Account: REXAN,
		Name:    ActN("closerex"),
		Authorization: []zsw.PermissionLevel{
			{Actor: owner, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(CloseREX{
			Ownwer: owner,
		}),
	}
}

type CloseREX struct {
	Ownwer zsw.AccountName
}
