package msig

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewCancel returns a `cancel` action that lives on the
// `zswhq.msig` contract.
func NewCancel(proposer zsw.AccountName, proposalName zsw.Name, canceler zsw.AccountName) *zsw.Action {
	return &zsw.Action{
		Account: zsw.AccountName("zswhq.msig"),
		Name:    zsw.ActionName("cancel"),
		// TODO: double check in this package that the `Actor` is always the `proposer`..
		Authorization: []zsw.PermissionLevel{
			{Actor: canceler, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(Cancel{proposer, proposalName, canceler}),
	}
}

type Cancel struct {
	Proposer     zsw.AccountName `json:"proposer"`
	ProposalName zsw.Name        `json:"proposal_name"`
	Canceler     zsw.AccountName `json:"canceler"`
}
