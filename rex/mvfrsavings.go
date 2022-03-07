package rex

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewMoveFromSavings(owner zsw.AccountName, rex zsw.Asset) *zsw.Action {
	return &zsw.Action{
		Account: REXAN,
		Name:    ActN("mvfrsavings"),
		Authorization: []zsw.PermissionLevel{
			{Actor: owner, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(MoveFromSavings{
			Owner: owner,
			REX:   rex,
		}),
	}
}

type MoveFromSavings struct {
	Owner zsw.AccountName
	REX   zsw.Asset
}
