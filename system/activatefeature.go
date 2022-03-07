package system

import (
	"github.com/zhongshuwen/zswchain-go"
)

func NewActivateFeature(featureDigest zsw.Checksum256) *zsw.Action {
	return &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("activate"),
		Authorization: []zsw.PermissionLevel{
			{Actor: AN("zswhq"), Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(Activate{
			FeatureDigest: featureDigest,
		}),
	}
}

// Activate represents a `activate` action on the `zswhq` contract.
type Activate struct {
	FeatureDigest zsw.Checksum256 `json:"feature_digest"`
}
