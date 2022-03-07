package rex

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewFundNetLoan(from zsw.AccountName, loanNumber uint64, payment zsw.Asset) *zsw.Action {
	return &zsw.Action{
		Account: REXAN,
		Name:    ActN("fundnetloan"),
		Authorization: []zsw.PermissionLevel{
			{Actor: from, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(FundNetLoan{
			From:       from,
			LoanNumber: loanNumber,
			Payment:    payment,
		}),
	}
}

type FundNetLoan struct {
	From       zsw.AccountName
	LoanNumber uint64
	Payment    zsw.Asset
}
