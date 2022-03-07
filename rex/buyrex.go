package rex

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewBuyREX(from zsw.AccountName, amount zsw.Asset) *zsw.Action {
	return &zsw.Action{
		Account: REXAN,
		Name:    ActN("buyrex"),
		Authorization: []zsw.PermissionLevel{
			{Actor: from, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(BuyREX{
			From:   from,
			Amount: amount,
		}),
	}
}

type BuyREX struct {
	From   zsw.AccountName
	Amount zsw.Asset
}
