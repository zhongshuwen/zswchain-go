package rex

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewDefundNetLoan(from zsw.AccountName, loanNumber uint64, amount zsw.Asset) *zsw.Action {
	return &zsw.Action{
		Account: REXAN,
		Name:    ActN("defnetloan"),
		Authorization: []zsw.PermissionLevel{
			{Actor: from, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(DefundNetLoan{
			From:       from,
			LoanNumber: loanNumber,
			Amount:     amount,
		}),
	}
}

type DefundNetLoan struct {
	From       zsw.AccountName
	LoanNumber uint64
	Amount     zsw.Asset
}
