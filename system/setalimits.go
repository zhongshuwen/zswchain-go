package system

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewSetalimits sets the account limits. Requires signature from `eosio@active` account.
func NewSetalimits(account zsw.AccountName, ramBytes, netWeight, cpuWeight int64) *zsw.Action {
	a := &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("setalimit"),
		Authorization: []zsw.PermissionLevel{
			{Actor: zsw.AccountName("zswhq"), Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(Setalimits{
			Account:   account,
			RAMBytes:  ramBytes,
			NetWeight: netWeight,
			CPUWeight: cpuWeight,
		}),
	}
	return a
}

// Setalimits represents the `zswhq.system::setalimit` action.
type Setalimits struct {
	Account   zsw.AccountName `json:"account"`
	RAMBytes  int64           `json:"ram_bytes"`
	NetWeight int64           `json:"net_weight"`
	CPUWeight int64           `json:"cpu_weight"`
}
