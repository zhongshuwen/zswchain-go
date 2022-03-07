package system

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewSetRAM(maxRAMSize uint64) *zsw.Action {
	a := &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("setram"),
		Authorization: []zsw.PermissionLevel{
			{
				Actor:      AN("zswhq"),
				Permission: zsw.PermissionName("active"),
			},
		},
		ActionData: zsw.NewActionData(SetRAM{
			MaxRAMSize: zsw.Uint64(maxRAMSize),
		}),
	}
	return a
}

// SetRAM represents the hard-coded `setram` action.
type SetRAM struct {
	MaxRAMSize zsw.Uint64 `json:"max_ram_size"`
}
