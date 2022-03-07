package rex

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewDefundCPULoan(from zsw.AccountName, loanNumber uint64, amount zsw.Asset) *zsw.Action {
	return &zsw.Action{
		Account: REXAN,
		Name:    ActN("defcpuloan"),
		Authorization: []zsw.PermissionLevel{
			{Actor: from, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(DefundCPULoan{
			From:       from,
			LoanNumber: loanNumber,
			Amount:     amount,
		}),
	}
}

type DefundCPULoan struct {
	From       zsw.AccountName
	LoanNumber uint64
	Amount     zsw.Asset
}
