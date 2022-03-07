package system

import (
	zsw "github.com/zhongshuwen/zswchain-go"
	"github.com/zhongshuwen/zswchain-go/ecc"
)

// NewRegProducer returns a `regproducer` action that lives on the
// `eosio.system` contract.
func NewRegProducer(producer zsw.AccountName, producerKey ecc.PublicKey, url string, location uint16) *zsw.Action {
	return &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("regproducer"),
		Authorization: []zsw.PermissionLevel{
			{Actor: producer, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(RegProducer{
			Producer:    producer,
			ProducerKey: producerKey,
			URL:         url,
			Location:    location,
		}),
	}
}

// RegProducer represents the `eosio.system::regproducer` action
type RegProducer struct {
	Producer    zsw.AccountName `json:"producer"`
	ProducerKey ecc.PublicKey   `json:"producer_key"`
	URL         string          `json:"url"`
	Location    uint16          `json:"location"` // what,s the meaning of that anyway ?
}
