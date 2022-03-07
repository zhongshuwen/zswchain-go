package zswitems

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewMakeRoyaltyUser(authorizer zsw.AccountName, newRoyaltyUser zsw.AccountName, zswId zsw.Uint128, altId zsw.Uint128, status uint32) *zsw.Action {
	return &zsw.Action{
		Account: AN("zsw.items"),
		Name:    ActN("mkroyaltyusr"),
		Authorization: []zsw.PermissionLevel{
			{Actor: authorizer, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(MakeRoyaltyUser{
			Authorizer: authorizer,
			NewRoyaltyUser: newRoyaltyUser,
			ZswId: zswId,
			AltId: altId,
			Status: status,
		}),
	}
}


type MakeRoyaltyUser struct {
  Authorizer zsw.AccountName `json:"authorizer"`
  NewRoyaltyUser zsw.AccountName `json:"newroyaltyusr"`
  ZswId zsw.Uint128 `json:"zsw_id"`
  AltId zsw.Uint128 `json:"alt_id"`
  Status uint32 `json:"status"`
}