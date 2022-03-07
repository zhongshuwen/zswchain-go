package system

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewSellRAM will sell at current market price a given number of
// bytes of RAM.
func NewSellRAM(account zsw.AccountName, bytes uint64) *zsw.Action {
	a := &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("sellram"),
		Authorization: []zsw.PermissionLevel{
			{Actor: account, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(SellRAM{
			Account: account,
			Bytes:   bytes,
		}),
	}
	return a
}

// SellRAM represents the `eosio.system::sellram` action.
type SellRAM struct {
	Account zsw.AccountName `json:"account"`
	Bytes   uint64          `json:"bytes"`
}
