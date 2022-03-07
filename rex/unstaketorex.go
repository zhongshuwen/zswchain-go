package rex

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewUnstakeToREX(
	owner zsw.AccountName,
	receiver zsw.AccountName,
	fromNet zsw.Asset,
	fromCPU zsw.Asset,
) *zsw.Action {
	return &zsw.Action{
		Account: REXAN,
		Name:    ActN("unstaketorex"),
		Authorization: []zsw.PermissionLevel{
			{Actor: owner, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(UnstakeToREX{
			Owner:    owner,
			Receiver: receiver,
			FromNet:  fromNet,
			FromCPU:  fromCPU,
		}),
	}
}

type UnstakeToREX struct {
	Owner    zsw.AccountName
	Receiver zsw.AccountName
	FromNet  zsw.Asset
	FromCPU  zsw.Asset
}
