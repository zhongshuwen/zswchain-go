package token

import zsw "github.com/zhongshuwen/zswchain-go"

func NewTransfer(from, to zsw.AccountName, quantity zsw.Asset, memo string) *zsw.Action {
	return &zsw.Action{
		Account: AN("zswhq.token"),
		Name:    ActN("transfer"),
		Authorization: []zsw.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(Transfer{
			From:     from,
			To:       to,
			Quantity: quantity,
			Memo:     memo,
		}),
	}
}

// Transfer represents the `transfer` struct on `zswhq.token` contract.
type Transfer struct {
	From     zsw.AccountName `json:"from"`
	To       zsw.AccountName `json:"to"`
	Quantity zsw.Asset       `json:"quantity"`
	Memo     string          `json:"memo"`
}
