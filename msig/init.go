package msig

import (
	"github.com/zhongshuwen/zswchain-go"
)

func init() {
	zsw.RegisterAction(AN("zswhq.msig"), ActN("propose"), &Propose{})
	zsw.RegisterAction(AN("zswhq.msig"), ActN("approve"), &Approve{})
	zsw.RegisterAction(AN("zswhq.msig"), ActN("unapprove"), &Unapprove{})
	zsw.RegisterAction(AN("zswhq.msig"), ActN("cancel"), &Cancel{})
	zsw.RegisterAction(AN("zswhq.msig"), ActN("exec"), &Exec{})
}

var AN = zsw.AN
var PN = zsw.PN
var ActN = zsw.ActN
