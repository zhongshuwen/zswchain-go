package system

import "github.com/zhongshuwen/zswchain-go"

// NewUnlinkAuth creates an action from the `zswhq.system` contract
// called `unlinkauth`.
//
// `unlinkauth` detaches a previously set permission from a
// `code::actionName`. See `linkauth`.
func NewUnlinkAuth(account, code zsw.AccountName, actionName zsw.ActionName) *zsw.Action {
	a := &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("unlinkauth"),
		Authorization: []zsw.PermissionLevel{
			{
				Actor:      account,
				Permission: zsw.PermissionName("active"),
			},
		},
		ActionData: zsw.NewActionData(UnlinkAuth{
			Account: account,
			Code:    code,
			Type:    actionName,
		}),
	}

	return a
}

// UnlinkAuth represents the native `unlinkauth` action, through the
// system contract.
type UnlinkAuth struct {
	Account zsw.AccountName `json:"account"`
	Code    zsw.AccountName `json:"code"`
	Type    zsw.ActionName  `json:"type"`
}
