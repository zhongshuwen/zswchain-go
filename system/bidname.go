package system

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewBidname(bidder, newname zsw.AccountName, bid zsw.Asset) *zsw.Action {
	a := &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("bidname"),
		Authorization: []zsw.PermissionLevel{
			{Actor: bidder, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(Bidname{
			Bidder:  bidder,
			Newname: newname,
			Bid:     bid,
		}),
	}
	return a
}

// Bidname represents the `zswhq.system_contract::bidname` action.
type Bidname struct {
	Bidder  zsw.AccountName `json:"bidder"`
	Newname zsw.AccountName `json:"newname"`
	Bid     zsw.Asset       `json:"bid"` // specified in EOS
}
