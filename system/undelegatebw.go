package system

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewUndelegateBW returns a `undelegatebw` action that lives on the
// `zswhq.system` contract.
func NewUndelegateBW(from, receiver zsw.AccountName, unstakeCPU, unstakeNet zsw.Asset) *zsw.Action {
	return &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("undelegatebw"),
		Authorization: []zsw.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(UndelegateBW{
			From:       from,
			Receiver:   receiver,
			UnstakeNet: unstakeNet,
			UnstakeCPU: unstakeCPU,
		}),
	}
}

// UndelegateBW represents the `zswhq.system::undelegatebw` action.
type UndelegateBW struct {
	From       zsw.AccountName `json:"from"`
	Receiver   zsw.AccountName `json:"receiver"`
	UnstakeNet zsw.Asset       `json:"unstake_net_quantity"`
	UnstakeCPU zsw.Asset       `json:"unstake_cpu_quantity"`
}
