package system

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewBuyRAMBytes will buy at current market price a given number of
// bytes of RAM, and grant them to the `receiver` account.
func NewBuyRAMBytes(payer, receiver zsw.AccountName, bytes uint32) *zsw.Action {
	a := &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("buyrambytes"),
		Authorization: []zsw.PermissionLevel{
			{Actor: payer, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(BuyRAMBytes{
			Payer:    payer,
			Receiver: receiver,
			Bytes:    bytes,
		}),
	}
	return a
}

// BuyRAMBytes represents the `zswhq.system::buyrambytes` action.
type BuyRAMBytes struct {
	Payer    zsw.AccountName `json:"payer"`
	Receiver zsw.AccountName `json:"receiver"`
	Bytes    uint32          `json:"bytes"`
}
