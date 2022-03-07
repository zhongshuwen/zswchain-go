package rex

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewRentCPU(
	from zsw.AccountName,
	receiver zsw.AccountName,
	loanPayment zsw.Asset,
	loanFund zsw.Asset,
) *zsw.Action {
	return &zsw.Action{
		Account: REXAN,
		Name:    ActN("rentcpu"),
		Authorization: []zsw.PermissionLevel{
			{Actor: from, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(RentCPU{
			From:        from,
			Receiver:    receiver,
			LoanPayment: loanPayment,
			LoanFund:    loanFund,
		}),
	}
}

type RentCPU struct {
	From        zsw.AccountName
	Receiver    zsw.AccountName
	LoanPayment zsw.Asset
	LoanFund    zsw.Asset
}
