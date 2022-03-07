package system

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewDelegateBW returns a `delegatebw` action that lives on the
// `eosio.system` contract.
func NewDelegateBW(from, receiver zsw.AccountName, stakeCPU, stakeNet zsw.Asset, transfer bool) *zsw.Action {
	return &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("delegatebw"),
		Authorization: []zsw.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(DelegateBW{
			From:     from,
			Receiver: receiver,
			StakeNet: stakeNet,
			StakeCPU: stakeCPU,
			Transfer: zsw.Bool(transfer),
		}),
	}
}

// DelegateBW represents the `eosio.system::delegatebw` action.
type DelegateBW struct {
	From     zsw.AccountName `json:"from"`
	Receiver zsw.AccountName `json:"receiver"`
	StakeNet zsw.Asset       `json:"stake_net_quantity"`
	StakeCPU zsw.Asset       `json:"stake_cpu_quantity"`
	Transfer zsw.Bool        `json:"transfer"`
}
