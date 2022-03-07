package forum

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewVote is an action representing a simple vote to be broadcast
// through the chain network.
func NewVote(voter zsw.AccountName, proposalName zsw.Name, voteValue uint8, voteJSON string) *zsw.Action {
	a := &zsw.Action{
		Account: ForumAN,
		Name:    ActN("vote"),
		Authorization: []zsw.PermissionLevel{
			{Actor: voter, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(Vote{
			Voter:        voter,
			ProposalName: proposalName,
			Vote:         voteValue,
			VoteJSON:     voteJSON,
		}),
	}
	return a
}

// Vote represents the `eosio.forum::vote` action.
type Vote struct {
	Voter        zsw.AccountName `json:"voter"`
	ProposalName zsw.Name        `json:"proposal_name"`
	Vote         uint8           `json:"vote"`
	VoteJSON     string          `json:"vote_json"`
}
