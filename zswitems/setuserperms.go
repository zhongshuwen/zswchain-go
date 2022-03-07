package zswitems

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewSetUserPermissions is an action representing a setting a user's permissions to be broadcast
// through the chain network.
func NewSetUserPermissions(sender, user zsw.AccountName, permissions zsw.Uint128) *zsw.Action {
	a := &zsw.Action{
		Account: ZswItemsAN,
		Name:    ActN("setuserperms"),
		Authorization: []zsw.PermissionLevel{
			{Actor: sender, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(SetUserPermissions{
			Sender:      sender,
			User:        user,
			Permissions: permissions,
		}),
	}
	return a
}

// UnVote represents the `zsw.items::unvote` action.
type SetUserPermissions struct {
	Sender      zsw.AccountName `json:"sender"`
	User        zsw.AccountName `json:"user"`
	Permissions zsw.Uint128     `json:"permissions"`
}
