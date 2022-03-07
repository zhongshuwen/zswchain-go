package forum

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewPropose is an action to submit a proposal for vote.
func NewPropose(proposer zsw.AccountName, proposalName zsw.Name, title string, proposalJSON string, expiresAt zsw.JSONTime) *zsw.Action {
	a := &zsw.Action{
		Account: ForumAN,
		Name:    ActN("propose"),
		Authorization: []zsw.PermissionLevel{
			{Actor: proposer, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(Propose{
			Proposer:     proposer,
			ProposalName: proposalName,
			Title:        title,
			ProposalJSON: proposalJSON,
			ExpiresAt:    expiresAt,
		}),
	}
	return a
}

// Propose represents the `eosio.forum::propose` action.
type Propose struct {
	Proposer     zsw.AccountName `json:"proposer"`
	ProposalName zsw.Name        `json:"proposal_name"`
	Title        string          `json:"title"`
	ProposalJSON string          `json:"proposal_json"`
	ExpiresAt    zsw.JSONTime    `json:"expires_at"`
}
