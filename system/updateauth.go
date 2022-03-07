package system

import "github.com/zhongshuwen/zswchain-go"

// NewUpdateAuth creates an action from the `eosio.system` contract
// called `updateauth`.
//
// usingPermission needs to be `owner` if you want to modify the
// `owner` authorization, otherwise `active` will do for the rest.
func NewUpdateAuth(account zsw.AccountName, permission, parent zsw.PermissionName, authority zsw.Authority, usingPermission zsw.PermissionName) *zsw.Action {
	a := &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("updateauth"),
		Authorization: []zsw.PermissionLevel{
			{
				Actor:      account,
				Permission: usingPermission,
			},
		},
		ActionData: zsw.NewActionData(UpdateAuth{
			Account:    account,
			Permission: permission,
			Parent:     parent,
			Auth:       authority,
		}),
	}

	return a
}

// UpdateAuth represents the hard-coded `updateauth` action.
//
// If you change the `active` permission, `owner` is the required parent.
//
// If you change the `owner` permission, there should be no parent.
type UpdateAuth struct {
	Account    zsw.AccountName    `json:"account"`
	Permission zsw.PermissionName `json:"permission"`
	Parent     zsw.PermissionName `json:"parent"`
	Auth       zsw.Authority      `json:"auth"`
}
