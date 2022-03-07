package rex

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewMoveToSavings(owner zsw.AccountName, rex zsw.Asset) *zsw.Action {
	return &zsw.Action{
		Account: REXAN,
		Name:    ActN("mvtosavings"),
		Authorization: []zsw.PermissionLevel{
			{Actor: owner, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(MoveToSavings{
			Owner: owner,
			REX:   rex,
		}),
	}
}

type MoveToSavings struct {
	Owner zsw.AccountName
	REX   zsw.Asset
}
