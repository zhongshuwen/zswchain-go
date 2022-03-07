package system

import (
	"github.com/zhongshuwen/zswchain-go"
	"github.com/zhongshuwen/zswchain-go/ecc"
)

// NewNewAccount returns a `newaccount` action that lives on the
// `zswhq.system` contract.
func NewNewAccount(creator, newAccount zsw.AccountName, publicKey ecc.PublicKey) *zsw.Action {
	return &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("newaccount"),
		Authorization: []zsw.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner: zsw.Authority{
				Threshold: 1,
				Keys: []zsw.KeyWeight{
					{
						PublicKey: publicKey,
						Weight:    1,
					},
				},
				Accounts: []zsw.PermissionLevelWeight{},
			},
			Active: zsw.Authority{
				Threshold: 1,
				Keys: []zsw.KeyWeight{
					{
						PublicKey: publicKey,
						Weight:    1,
					},
				},
				Accounts: []zsw.PermissionLevelWeight{},
			},
		}),
	}
}

// NewDelegatedNewAccount returns a `newaccount` action that lives on the
// `zswhq.system` contract. It is filled with an authority structure that
// delegates full control of the new account to an already existing account.
func NewDelegatedNewAccount(creator, newAccount zsw.AccountName, delegatedTo zsw.AccountName) *zsw.Action {
	return &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("newaccount"),
		Authorization: []zsw.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner: zsw.Authority{
				Threshold: 1,
				Keys:      []zsw.KeyWeight{},
				Accounts: []zsw.PermissionLevelWeight{
					zsw.PermissionLevelWeight{
						Permission: zsw.PermissionLevel{
							Actor:      delegatedTo,
							Permission: PN("active"),
						},
						Weight: 1,
					},
				},
			},
			Active: zsw.Authority{
				Threshold: 1,
				Keys:      []zsw.KeyWeight{},
				Accounts: []zsw.PermissionLevelWeight{
					zsw.PermissionLevelWeight{
						Permission: zsw.PermissionLevel{
							Actor:      delegatedTo,
							Permission: PN("active"),
						},
						Weight: 1,
					},
				},
			},
		}),
	}
}

// NewCustomNewAccount returns a `newaccount` action that lives on the
// `zswhq.system` contract. You can specify your own `owner` and
// `active` permissions.
func NewCustomNewAccount(creator, newAccount zsw.AccountName, owner, active zsw.Authority) *zsw.Action {
	return &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("newaccount"),
		Authorization: []zsw.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner:   owner,
			Active:  active,
		}),
	}
}

// NewAccount represents a `newaccount` action on the `zswhq.system`
// contract. It is one of the rare ones to be hard-coded into the
// blockchain.
type NewAccount struct {
	Creator zsw.AccountName `json:"creator"`
	Name    zsw.AccountName `json:"name"`
	Owner   zsw.Authority   `json:"owner"`
	Active  zsw.Authority   `json:"active"`
}
