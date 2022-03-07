package forum

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewUnVote is an action representing the action to undoing a current vote
func NewUnVote(voter zsw.AccountName, proposalName zsw.Name) *zsw.Action {
	a := &zsw.Action{
		Account: ForumAN,
		Name:    ActN("unvote"),
		Authorization: []zsw.PermissionLevel{
			{Actor: voter, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(UnVote{
			Voter:        voter,
			ProposalName: proposalName,
		}),
	}
	return a
}

// UnVote represents the `eosio.forum::unvote` action.
type UnVote struct {
	Voter        zsw.AccountName `json:"voter"`
	ProposalName zsw.Name        `json:"proposal_name"`
}
