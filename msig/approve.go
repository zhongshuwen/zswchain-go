package msig

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewApprove returns a `approve` action that lives on the
// `eosio.msig` contract.
func NewApprove(proposer zsw.AccountName, proposalName zsw.Name, level zsw.PermissionLevel) *zsw.Action {
	return &zsw.Action{
		Account:       zsw.AccountName("zswhq.msig"),
		Name:          zsw.ActionName("approve"),
		Authorization: []zsw.PermissionLevel{level},
		ActionData:    zsw.NewActionData(Approve{proposer, proposalName, level}),
	}
}

type Approve struct {
	Proposer     zsw.AccountName     `json:"proposer"`
	ProposalName zsw.Name            `json:"proposal_name"`
	Level        zsw.PermissionLevel `json:"level"`
}
