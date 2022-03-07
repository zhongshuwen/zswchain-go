package rex

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewUpdateREX(owner zsw.AccountName) *zsw.Action {
	return &zsw.Action{
		Account: REXAN,
		Name:    ActN("updaterex"),
		Authorization: []zsw.PermissionLevel{
			{Actor: owner, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(UpdateREX{
			Owner: owner,
		}),
	}
}

type UpdateREX struct {
	Owner zsw.AccountName
}
