package rex

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewSellREX(from zsw.AccountName, rex zsw.Asset) *zsw.Action {
	return &zsw.Action{
		Account: REXAN,
		Name:    ActN("sellrex"),
		Authorization: []zsw.PermissionLevel{
			{Actor: from, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(SellREX{
			From: from,
			REX:  rex,
		}),
	}
}

type SellREX struct {
	From zsw.AccountName
	REX  zsw.Asset
}
