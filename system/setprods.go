package system

import (
	zsw "github.com/zhongshuwen/zswchain-go"
	"github.com/zhongshuwen/zswchain-go/ecc"
)

// NewSetPriv returns a `setpriv` action that lives on the
// `eosio.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `eos-bios` boot process by the
// `eosio.system` contract.
func NewSetProds(producers []ProducerKey) *zsw.Action {
	a := &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("setprods"),
		Authorization: []zsw.PermissionLevel{
			{Actor: AN("zswhq"), Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(SetProds{
			Schedule: producers,
		}),
	}
	return a
}

// SetProds is present in `eosio.bios` contract. Used only at boot time.
type SetProds struct {
	Schedule []ProducerKey `json:"schedule"`
}

type ProducerKey struct {
	ProducerName    zsw.AccountName `json:"producer_name"`
	BlockSigningKey ecc.PublicKey   `json:"block_signing_key"`
}
