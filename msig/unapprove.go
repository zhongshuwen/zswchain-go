package msig

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewUnapprove returns a `unapprove` action that lives on the
// `eosio.msig` contract.
func NewUnapprove(proposer zsw.AccountName, proposalName zsw.Name, level zsw.PermissionLevel) *zsw.Action {
	return &zsw.Action{
		Account:       zsw.AccountName("zswhq.msig"),
		Name:          zsw.ActionName("unapprove"),
		Authorization: []zsw.PermissionLevel{level},
		ActionData:    zsw.NewActionData(Unapprove{proposer, proposalName, level}),
	}
}

type Unapprove struct {
	Proposer     zsw.AccountName     `json:"proposer"`
	ProposalName zsw.Name            `json:"proposal_name"`
	Level        zsw.PermissionLevel `json:"level"`
}
