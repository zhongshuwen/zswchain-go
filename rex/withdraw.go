package rex

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewWithdraw(owner zsw.AccountName, amount zsw.Asset) *zsw.Action {
	return &zsw.Action{
		Account: REXAN,
		Name:    ActN("withdraw"),
		Authorization: []zsw.PermissionLevel{
			{Actor: owner, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(Withdraw{
			Owner:  owner,
			Amount: amount,
		}),
	}
}

type Withdraw struct {
	Owner  zsw.AccountName
	Amount zsw.Asset
}
