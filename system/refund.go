package system

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewRefund returns a `refund` action that lives on the
// `zswhq.system` contract.
func NewRefund(owner zsw.AccountName) *zsw.Action {
	return &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("refund"),
		Authorization: []zsw.PermissionLevel{
			{Actor: owner, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(Refund{
			Owner: owner,
		}),
	}
}

// Refund represents the `zswhq.system::refund` action
type Refund struct {
	Owner zsw.AccountName `json:"owner"`
}
