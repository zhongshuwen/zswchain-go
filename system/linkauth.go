package system

import "github.com/zhongshuwen/zswchain-go"

// NewLinkAuth creates an action from the `zswhq.system` contract
// called `linkauth`.
//
// `linkauth` allows you to attach certain permission to the given
// `code::actionName`. With this set on-chain, you can use the
// `requiredPermission` to sign transactions for `code::actionName`
// and not rely on your `active` (which might be more sensitive as it
// can sign anything) for the given operation.
func NewLinkAuth(account, code zsw.AccountName, actionName zsw.ActionName, requiredPermission zsw.PermissionName) *zsw.Action {
	a := &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("linkauth"),
		Authorization: []zsw.PermissionLevel{
			{
				Actor:      account,
				Permission: zsw.PermissionName("active"),
			},
		},
		ActionData: zsw.NewActionData(LinkAuth{
			Account:     account,
			Code:        code,
			Type:        actionName,
			Requirement: requiredPermission,
		}),
	}

	return a
}

// LinkAuth represents the native `linkauth` action, through the
// system contract.
type LinkAuth struct {
	Account     zsw.AccountName    `json:"account"`
	Code        zsw.AccountName    `json:"code"`
	Type        zsw.ActionName     `json:"type"`
	Requirement zsw.PermissionName `json:"requirement"`
}
