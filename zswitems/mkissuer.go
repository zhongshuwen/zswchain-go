package zswitems

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewMakeIssuer(authorizer zsw.AccountName, issuerName zsw.AccountName, zswId zsw.Uint128, altId zsw.Uint128, permissions zsw.Uint128, status uint32) *zsw.Action {
	return &zsw.Action{
		Account: AN("zsw.items"),
		Name:    ActN("mkissuer"),
		Authorization: []zsw.PermissionLevel{
			{Actor: authorizer, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(MakeIssuer{
			Authorizer: authorizer,
			IssuerName: issuerName,
			ZswId: zswId,
			AltId: altId,
			Permissions: permissions,
			Status: status,
		}),
	}
}


type MakeIssuer struct {
  Authorizer zsw.AccountName `json:"authorizer"`
  IssuerName zsw.AccountName `json:"issuer_name"`
  ZswId zsw.Uint128 `json:"zsw_id"`
  AltId zsw.Uint128 `json:"alt_id"`
  Permissions zsw.Uint128 `json:"permissions"`
  Status uint32 `json:"status"`
}