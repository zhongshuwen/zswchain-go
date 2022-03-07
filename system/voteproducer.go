package system

import "github.com/zhongshuwen/zswchain-go"

// NewNonce returns a `nonce` action that lives on the
// `eosio.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `eos-bios` boot process by the
// `eosio.system` contract.
func NewVoteProducer(voter zsw.AccountName, proxy zsw.AccountName, producers ...zsw.AccountName) *zsw.Action {
	a := &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("voteproducer"),
		Authorization: []zsw.PermissionLevel{
			{Actor: voter, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(
			VoteProducer{
				Voter:     voter,
				Proxy:     proxy,
				Producers: producers,
			},
		),
	}
	return a
}

// VoteProducer represents the `eosio.system::voteproducer` action
type VoteProducer struct {
	Voter     zsw.AccountName   `json:"voter"`
	Proxy     zsw.AccountName   `json:"proxy"`
	Producers []zsw.AccountName `json:"producers"`
}
