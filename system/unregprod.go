package system

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewUnregProducer returns a `unregprod` action that lives on the
// `zswhq.system` contract.
func NewUnregProducer(producer zsw.AccountName) *zsw.Action {
	return &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("unregprod"),
		Authorization: []zsw.PermissionLevel{
			{Actor: producer, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(UnregProducer{
			Producer: producer,
		}),
	}
}

// UnregProducer represents the `zswhq.system::unregprod` action
type UnregProducer struct {
	Producer zsw.AccountName `json:"producer"`
}
