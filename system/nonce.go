package system

import "github.com/zhongshuwen/zswchain-go"

// NewNonce returns a `nonce` action that lives on the
// `zswhq.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `eos-bios` boot process by the
// `zswhq.system` contract.
func NewNonce(nonce string) *zsw.Action {
	a := &zsw.Action{
		Account:       AN("zswhq"),
		Name:          ActN("nonce"),
		Authorization: []zsw.PermissionLevel{
			//{Actor: AN("zswhq"), Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(Nonce{
			Value: nonce,
		}),
	}
	return a
}
