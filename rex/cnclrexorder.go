package rex

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewCancelREXOrder(owner zsw.AccountName) *zsw.Action {
	return &zsw.Action{
		Account: REXAN,
		Name:    ActN("cnclrexorder"),
		Authorization: []zsw.PermissionLevel{
			{Actor: owner, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(CancelREXOrder{
			Owner: owner,
		}),
	}
}

type CancelREXOrder struct {
	Owner zsw.AccountName
}
