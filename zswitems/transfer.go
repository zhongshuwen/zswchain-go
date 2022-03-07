package zswitems

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewItemTransfer is an action representing a simple item transfer to be broadcast
// through the chain network.

func NewItemTransfer(authorizer, from, to, fromCustodian, toCustodian zsw.AccountName, freezeTime uint32, useLiquidBalance bool, maxUnfreezeIterations uint32, itemIds []uint64, amounts []uint64, memo string) *zsw.Action {
	var authorizers []zsw.PermissionLevel
	if useLiquidBalance == true && fromCustodian != "nullnullnull" {
		authorizers = []zsw.PermissionLevel{
			{Actor: authorizer, Permission: PN("active")},
			{Actor: fromCustodian, Permission: PN("active")},
		}
		
	}else{
		authorizers = []zsw.PermissionLevel{
			{Actor: authorizer, Permission: PN("active")},
		}
	}
	return &zsw.Action{
		Account: AN("zsw.items"),
		Name:    ActN("transfer"),
		Authorization: authorizers,
		ActionData: zsw.NewActionData(ItemTransfer{
			Authorizer:            authorizer,
			From:                  from,
			To:                    to,
			FromCustodian:         fromCustodian,
			ToCustodian:           toCustodian,
			FreezeTime:            freezeTime,
			UseLiquidBalance:      useLiquidBalance,
			MaxUnfreezeIterations: maxUnfreezeIterations,
			ItemIds:               itemIds,
			Amounts:               amounts,
			Memo:                  memo,
		}),
	}
}

// ItemTransfer represents the `transfer` struct on `zsw.items` contract.
type ItemTransfer struct {
	Authorizer            zsw.AccountName `json:"authorizer"`
	From                  zsw.AccountName `json:"from"`
	To                    zsw.AccountName `json:"to"`
	FromCustodian         zsw.AccountName `json:"from_custodian"`
	ToCustodian           zsw.AccountName `json:"to_custodian"`
	FreezeTime            uint32          `json:"freeze_time"`
	UseLiquidBalance      bool            `json:"use_liquid_backup"`
	MaxUnfreezeIterations uint32          `json:"max_unfreeze_iterations"`
	ItemIds               []uint64        `json:"item_ids"`
	Amounts               []uint64        `json:"amounts"`
	Memo                  string          `json:"memo"`
}
