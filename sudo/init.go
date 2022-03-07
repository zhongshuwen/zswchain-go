package sudo

import zsw "github.com/zhongshuwen/zswchain-go"

func init() {
	zsw.RegisterAction(AN("zswhq.wrap"), ActN("exec"), Exec{})
}

var AN = zsw.AN
var ActN = zsw.ActN
