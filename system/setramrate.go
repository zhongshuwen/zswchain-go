package system

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewSetRAMRate(bytesPerBlock uint16) *zsw.Action {
	a := &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("setram"),
		Authorization: []zsw.PermissionLevel{
			{
				Actor:      AN("zswhq"),
				Permission: zsw.PermissionName("active"),
			},
		},
		ActionData: zsw.NewActionData(SetRAMRate{
			BytesPerBlock: bytesPerBlock,
		}),
	}
	return a
}

// SetRAMRate represents the system contract's `setramrate` action.
type SetRAMRate struct {
	BytesPerBlock uint16 `json:"bytes_per_block"`
}
