package forum

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewExpire is an action to expire a proposal ahead of its natural death.
func NewExpire(proposer zsw.AccountName, proposalName zsw.Name) *zsw.Action {
	a := &zsw.Action{
		Account: ForumAN,
		Name:    ActN("expire"),
		Authorization: []zsw.PermissionLevel{
			{Actor: proposer, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(Expire{
			ProposalName: proposalName,
		}),
	}
	return a
}

// Expire represents the `eosio.forum::propose` action.
type Expire struct {
	ProposalName zsw.Name `json:"proposal_name"`
}
