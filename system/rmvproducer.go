package system

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewRemoveProducer returns a `rmvproducer` action that lives on the
// `eosio.system` contract.  This is to be called by the consortium of
// BPs, to oust a BP from its place.  If you want to unregister
// yourself as a BP, use `unregprod`.
func NewRemoveProducer(producer zsw.AccountName) *zsw.Action {
	return &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("rmvproducer"),
		Authorization: []zsw.PermissionLevel{
			{Actor: AN("zswhq"), Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(RemoveProducer{
			Producer: producer,
		}),
	}
}

// RemoveProducer represents the `eosio.system::rmvproducer` action
type RemoveProducer struct {
	Producer zsw.AccountName `json:"producer"`
}
