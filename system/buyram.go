package system

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewBuyRAM(payer, receiver zsw.AccountName, eosQuantity uint64) *zsw.Action {
	a := &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("buyram"),
		Authorization: []zsw.PermissionLevel{
			{Actor: payer, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(BuyRAM{
			Payer:    payer,
			Receiver: receiver,
			Quantity: zsw.NewEOSAsset(int64(eosQuantity)),
		}),
	}
	return a
}

// BuyRAM represents the `eosio.system::buyram` action.
type BuyRAM struct {
	Payer    zsw.AccountName `json:"payer"`
	Receiver zsw.AccountName `json:"receiver"`
	Quantity zsw.Asset       `json:"quant"` // specified in EOS
}
