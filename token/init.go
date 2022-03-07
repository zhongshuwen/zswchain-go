package token

import "github.com/zhongshuwen/zswchain-go"

func init() {
	zsw.RegisterAction(AN("zswhq.token"), ActN("transfer"), Transfer{})
	zsw.RegisterAction(AN("zswhq.token"), ActN("issue"), Issue{})
	zsw.RegisterAction(AN("zswhq.token"), ActN("create"), Create{})
}
