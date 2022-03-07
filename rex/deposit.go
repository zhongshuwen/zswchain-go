package rex

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewDeposit(owner zsw.AccountName, amount zsw.Asset) *zsw.Action {
	return &zsw.Action{
		Account: REXAN,
		Name:    ActN("deposit"),
		Authorization: []zsw.PermissionLevel{
			{Actor: owner, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(Deposit{
			Owner:  owner,
			Amount: amount,
		}),
	}
}

type Deposit struct {
	Owner  zsw.AccountName
	Amount zsw.Asset
}
