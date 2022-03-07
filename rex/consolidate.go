package rex

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewConsolidate(owner zsw.AccountName) *zsw.Action {
	return &zsw.Action{
		Account: REXAN,
		Name:    ActN("consolidate"),
		Authorization: []zsw.PermissionLevel{
			{Actor: owner, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(Consolidate{
			Owner: owner,
		}),
	}
}

type Consolidate struct {
	Owner zsw.AccountName
}
