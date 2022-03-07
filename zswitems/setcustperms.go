package zswitems

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewSetCustodianPermissions is an action representing a setting a custodian's permissions to be broadcast
// through the chain network.
func NewSetCustodianPermissions(sender, custodian zsw.AccountName, permissions zsw.Uint128) *zsw.Action {
	a := &zsw.Action{
		Account: ZswItemsAN,
		Name:    ActN("setcustperms"),
		Authorization: []zsw.PermissionLevel{
			{Actor: sender, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(SetCustodianPermissions{
			Sender:      sender,
			Custodian:   custodian,
			Permissions: permissions,
		}),
	}
	return a
}

// UnVote represents the `zsw.items::unvote` action.
type SetCustodianPermissions struct {
	Sender      zsw.AccountName `json:"sender"`
	Custodian   zsw.AccountName `json:"custodian"`
	Permissions zsw.Uint128     `json:"permissions"`
}
