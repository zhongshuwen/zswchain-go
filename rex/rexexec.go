package rex

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewREXExec(user zsw.AccountName, max uint16) *zsw.Action {
	return &zsw.Action{
		Account: REXAN,
		Name:    ActN("rexexec"),
		Authorization: []zsw.PermissionLevel{
			{Actor: user, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(REXExec{
			User: user,
			Max:  max,
		}),
	}
}

type REXExec struct {
	User zsw.AccountName
	Max  uint16
}
