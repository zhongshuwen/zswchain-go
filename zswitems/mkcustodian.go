package zswitems

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewMakeCustodian(creator zsw.AccountName, custodianName zsw.AccountName, zswId zsw.Uint128, altId zsw.Uint128, permissions zsw.Uint128, status uint32, incomingFreezePeriod uint32, notifyAccounts []zsw.AccountName) *zsw.Action {
	return &zsw.Action{
		Account: AN("zsw.items"),
		Name:    ActN("mkcustodian"),
		Authorization: []zsw.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(MakeCustodian{
			Creator: creator,
			CustodianName: custodianName,
			ZswId: zswId,
			AltId: altId,
			Permissions: permissions,
			Status: status,
			IncomingFreezePeriod: incomingFreezePeriod,
			NotifyAccounts: notifyAccounts,
		}),
	}
}


type MakeCustodian struct {
  Creator zsw.AccountName `json:"creator"`
  CustodianName zsw.AccountName `json:"custodian_name"`
  ZswId zsw.Uint128 `json:"zsw_id"`
  AltId zsw.Uint128 `json:"alt_id"`
  Permissions zsw.Uint128 `json:"permissions"`
  Status uint32 `json:"status"`
  IncomingFreezePeriod uint32 `json:"incoming_freeze_period"`
  NotifyAccounts []zsw.AccountName `json:"notify_accounts"`
}