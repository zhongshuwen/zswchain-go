package system

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewInitSystem returns a `init` action that lives on the
// `zswhq.system` contract.
func NewInitSystem(version zsw.Varuint32, core zsw.Symbol) *zsw.Action {
	return &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("init"),
		Authorization: []zsw.PermissionLevel{
			{
				Actor:      AN("zswhq"),
				Permission: zsw.PermissionName("active"),
			},
		},
		ActionData: zsw.NewActionData(Init{
			Version: version,
			Core:    core,
		}),
	}
}

// Init represents the `zswhq.system::init` action
type Init struct {
	Version zsw.Varuint32 `json:"version"`
	Core    zsw.Symbol    `json:"core"`
}
