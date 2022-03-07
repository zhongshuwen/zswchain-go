package msig

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewExec returns a `exec` action that lives on the
// `zswhq.msig` contract.
func NewExec(proposer zsw.AccountName, proposalName zsw.Name, executer zsw.AccountName) *zsw.Action {
	return &zsw.Action{
		Account: zsw.AccountName("zswhq.msig"),
		Name:    zsw.ActionName("exec"),
		// TODO: double check in this package that the `Actor` is always the `proposer`..
		Authorization: []zsw.PermissionLevel{
			{Actor: executer, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(Exec{proposer, proposalName, executer}),
	}
}

type Exec struct {
	Proposer     zsw.AccountName `json:"proposer"`
	ProposalName zsw.Name        `json:"proposal_name"`
	Executer     zsw.AccountName `json:"executer"`
}
