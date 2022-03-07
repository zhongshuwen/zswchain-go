package zswitems

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewItemMint is an action representing a simple item minting to be broadcast
// through the chain network.

func NewItemMint(minter, to, toCustodian zsw.AccountName, freezeTime uint32, itemIds []uint64, amounts []uint64, memo string) *zsw.Action {
	return &zsw.Action{
		Account: AN("zsw.items"),
		Name:    ActN("transfer"),
		Authorization: []zsw.PermissionLevel{
			{Actor: minter, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(ItemMint{
			Minter:      minter,
			To:          to,
			ToCustodian: toCustodian,
			ItemIds:     itemIds,
			Amounts:     amounts,
			Memo:        memo,
			FreezeTime:  freezeTime,
		}),
	}
}

// ItemTransfer represents the `transfer` struct on `zsw.items` contract.
type ItemMint struct {
	Minter      zsw.AccountName `json:"minter"`
	To          zsw.AccountName `json:"to"`
	ToCustodian zsw.AccountName `json:"to_custodian"`
	ItemIds     []uint64        `json:"item_ids"`
	Amounts     []uint64        `json:"amounts"`
	Memo        string          `json:"memo"`
	FreezeTime  uint32          `json:"freeze_time"`
}
