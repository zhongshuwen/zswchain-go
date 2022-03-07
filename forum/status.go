package forum

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// Status is an action to set a status update for a given account on the forum contract.
func NewStatus(account zsw.AccountName, content string) *zsw.Action {
	a := &zsw.Action{
		Account: ForumAN,
		Name:    ActN("status"),
		Authorization: []zsw.PermissionLevel{
			{Actor: account, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(Status{
			Account: account,
			Content: content,
		}),
	}
	return a
}

// Status represents the `zswhq.forum::status` action.
type Status struct {
	Account zsw.AccountName `json:"account_name"`
	Content string          `json:"content"`
}
