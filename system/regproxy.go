package system

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewRegProxy returns a `regproxy` action that lives on the
// `eosio.system` contract.
func NewRegProxy(proxy zsw.AccountName, isProxy bool) *zsw.Action {
	return &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("regproxy"),
		Authorization: []zsw.PermissionLevel{
			{Actor: proxy, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(RegProxy{
			Proxy:   proxy,
			IsProxy: isProxy,
		}),
	}
}

// RegProxy represents the `eosio.system::regproxy` action
type RegProxy struct {
	Proxy   zsw.AccountName `json:"proxy"`
	IsProxy bool            `json:"isproxy"`
}
