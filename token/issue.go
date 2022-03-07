package token

import zsw "github.com/zhongshuwen/zswchain-go"

func NewIssue(to zsw.AccountName, quantity zsw.Asset, memo string) *zsw.Action {
	return &zsw.Action{
		Account: AN("zswhq.token"),
		Name:    ActN("issue"),
		Authorization: []zsw.PermissionLevel{
			{Actor: to, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(Issue{
			To:       to,
			Quantity: quantity,
			Memo:     memo,
		}),
	}
}

// Issue represents the `issue` struct on the `eosio.token` contract.
type Issue struct {
	To       zsw.AccountName `json:"to"`
	Quantity zsw.Asset       `json:"quantity"`
	Memo     string          `json:"memo"`
}
