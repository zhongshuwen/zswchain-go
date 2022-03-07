package sudo

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewExec creates an `exec` action, found in the `zswhq.wrap`
// contract.
//
// Given an `zsw.Transaction`, call `zsw.MarshalBinary` on it first,
// pass the resulting bytes as `zsw.HexBytes` here.
func NewExec(executer zsw.AccountName, transaction zsw.Transaction) *zsw.Action {
	a := &zsw.Action{
		Account: zsw.AccountName("zswhq.wrap"),
		Name:    zsw.ActionName("exec"),
		Authorization: []zsw.PermissionLevel{
			{Actor: executer, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(Exec{
			Executer:    executer,
			Transaction: transaction,
		}),
	}
	return a
}

// Exec represents the `zswhq.system::exec` action.
type Exec struct {
	Executer    zsw.AccountName `json:"executer"`
	Transaction zsw.Transaction `json:"trx"`
}
