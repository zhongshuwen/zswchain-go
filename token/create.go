package token

import zsw "github.com/zhongshuwen/zswchain-go"

func NewCreate(issuer zsw.AccountName, maxSupply zsw.Asset) *zsw.Action {
	return &zsw.Action{
		Account: AN("zswhq.token"),
		Name:    ActN("create"),
		Authorization: []zsw.PermissionLevel{
			{Actor: AN("zswhq.token"), Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(Create{
			Issuer:        issuer,
			MaximumSupply: maxSupply,
		}),
	}
}

// Create represents the `create` struct on the `zswhq.token` contract.
type Create struct {
	Issuer        zsw.AccountName `json:"issuer"`
	MaximumSupply zsw.Asset       `json:"maximum_supply"`
}
