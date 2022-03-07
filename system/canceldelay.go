package system

import "github.com/zhongshuwen/zswchain-go"

// NewCancelDelay creates an action from the `zswhq.system` contract
// called `canceldelay`.
//
// `canceldelay` allows you to cancel a deferred transaction,
// previously sent to the chain with a `delay_sec` larger than 0.  You
// need to sign with cancelingAuth, to cancel a transaction signed
// with that same authority.
func NewCancelDelay(cancelingAuth zsw.PermissionLevel, transactionID zsw.Checksum256) *zsw.Action {
	a := &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("canceldelay"),
		Authorization: []zsw.PermissionLevel{
			cancelingAuth,
		},
		ActionData: zsw.NewActionData(CancelDelay{
			CancelingAuth: cancelingAuth,
			TransactionID: transactionID,
		}),
	}

	return a
}

// CancelDelay represents the native `canceldelay` action, through the
// system contract.
type CancelDelay struct {
	CancelingAuth zsw.PermissionLevel `json:"canceling_auth"`
	TransactionID zsw.Checksum256     `json:"trx_id"`
}
