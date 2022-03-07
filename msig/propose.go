package msig

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewPropose returns a `propose` action that lives on the
// `eosio.msig` contract.
func NewPropose(proposer zsw.AccountName, proposalName zsw.Name, requested []zsw.PermissionLevel, transaction *zsw.Transaction) *zsw.Action {
	return &zsw.Action{
		Account: zsw.AccountName("zswhq.msig"),
		Name:    zsw.ActionName("propose"),
		Authorization: []zsw.PermissionLevel{
			{Actor: proposer, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(Propose{proposer, proposalName, requested, transaction}),
	}
}

type Propose struct {
	Proposer     zsw.AccountName       `json:"proposer"`
	ProposalName zsw.Name              `json:"proposal_name"`
	Requested    []zsw.PermissionLevel `json:"requested"`
	Transaction  *zsw.Transaction      `json:"trx"`
}
