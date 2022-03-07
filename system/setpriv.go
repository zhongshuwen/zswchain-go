package system

import zsw "github.com/zhongshuwen/zswchain-go"

// NewSetPriv returns a `setpriv` action that lives on the
// `eosio.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `eos-bios` boot process by the
// `eosio.system` contract.
func NewSetPriv(account zsw.AccountName) *zsw.Action {
	a := &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("setpriv"),
		Authorization: []zsw.PermissionLevel{
			{Actor: AN("zswhq"), Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(SetPriv{
			Account: account,
			IsPriv:  zsw.Bool(true),
		}),
	}
	return a
}

// SetPriv sets privileged account status. Used in the bios boot mechanism.
type SetPriv struct {
	Account zsw.AccountName `json:"account"`
	IsPriv  zsw.Bool        `json:"is_priv"`
}
