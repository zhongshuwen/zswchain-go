package forum

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// CleanProposal is an action to flush proposal and allow RAM used by it.
func NewCleanProposal(cleaner zsw.AccountName, proposalName zsw.Name, maxCount uint64) *zsw.Action {
	a := &zsw.Action{
		Account: ForumAN,
		Name:    ActN("clnproposal"),
		Authorization: []zsw.PermissionLevel{
			{Actor: cleaner, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(CleanProposal{
			ProposalName: proposalName,
			MaxCount:     maxCount,
		}),
	}
	return a
}

// CleanProposal represents the `eosio.forum::clnproposal` action.
type CleanProposal struct {
	ProposalName zsw.Name `json:"proposal_name"`
	MaxCount     uint64   `json:"max_count"`
}
