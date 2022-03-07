package rex

import zsw "github.com/zhongshuwen/zswchain-go"

func init() {
	zsw.RegisterAction(REXAN, ActN("buyrex"), BuyREX{})
	zsw.RegisterAction(REXAN, ActN("closerex"), CloseREX{})
	zsw.RegisterAction(REXAN, ActN("cnclrexorder"), CancelREXOrder{})
	zsw.RegisterAction(REXAN, ActN("consolidate"), Consolidate{})
	zsw.RegisterAction(REXAN, ActN("defcpuloan"), DefundCPULoan{})
	zsw.RegisterAction(REXAN, ActN("defnetloan"), DefundNetLoan{})
	zsw.RegisterAction(REXAN, ActN("deposit"), Deposit{})
	zsw.RegisterAction(REXAN, ActN("fundcpuloan"), FundCPULoan{})
	zsw.RegisterAction(REXAN, ActN("fundnetloan"), FundNetLoan{})
	zsw.RegisterAction(REXAN, ActN("mvfrsavings"), MoveFromSavings{})
	zsw.RegisterAction(REXAN, ActN("mvtosavings"), MoveToSavings{})
	zsw.RegisterAction(REXAN, ActN("rentcpu"), RentCPU{})
	zsw.RegisterAction(REXAN, ActN("rentnet"), RentNet{})
	zsw.RegisterAction(REXAN, ActN("rexexec"), REXExec{})
	zsw.RegisterAction(REXAN, ActN("sellrex"), SellREX{})
	zsw.RegisterAction(REXAN, ActN("unstaketorex"), UnstakeToREX{})
	zsw.RegisterAction(REXAN, ActN("updaterex"), UpdateREX{})
	zsw.RegisterAction(REXAN, ActN("withdraw"), Withdraw{})
}

var AN = zsw.AN
var PN = zsw.PN
var ActN = zsw.ActN

var REXAN = AN("zswhq")
