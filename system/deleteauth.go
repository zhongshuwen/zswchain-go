package system

import "github.com/zhongshuwen/zswchain-go"

// NewDeleteAuth creates an action from the `eosio.system` contract
// called `deleteauth`.
//
// You cannot delete the `owner` or `active` permissions.  Also, if a
// permission is still linked through a previous `updatelink` action,
// you will need to `unlinkauth` first.
func NewDeleteAuth(account zsw.AccountName, permission zsw.PermissionName) *zsw.Action {
	a := &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("deleteauth"),
		Authorization: []zsw.PermissionLevel{
			{Actor: account, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(DeleteAuth{
			Account:    account,
			Permission: permission,
		}),
	}

	return a
}

// DeleteAuth represents the native `deleteauth` action, reachable
// through the `eosio.system` contract.
type DeleteAuth struct {
	Account    zsw.AccountName    `json:"account"`
	Permission zsw.PermissionName `json:"permission"`
}
