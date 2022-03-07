package rex

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewFundCPULoan(from zsw.AccountName, loanNumber uint64, payment zsw.Asset) *zsw.Action {
	return &zsw.Action{
		Account: REXAN,
		Name:    ActN("fundcpuloan"),
		Authorization: []zsw.PermissionLevel{
			{Actor: from, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(FundCPULoan{
			From:       from,
			LoanNumber: loanNumber,
			Payment:    payment,
		}),
	}
}

type FundCPULoan struct {
	From       zsw.AccountName
	LoanNumber uint64
	Payment    zsw.Asset
}
