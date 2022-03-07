package system

import (
	"github.com/zhongshuwen/zswchain-go"
)

func init() {
	zsw.RegisterAction(AN("zswhq"), ActN("setcode"), SetCode{})
	zsw.RegisterAction(AN("zswhq"), ActN("setabi"), SetABI{})
	zsw.RegisterAction(AN("zswhq"), ActN("newaccount"), NewAccount{})
	zsw.RegisterAction(AN("zswhq"), ActN("delegatebw"), DelegateBW{})
	zsw.RegisterAction(AN("zswhq"), ActN("undelegatebw"), UndelegateBW{})
	zsw.RegisterAction(AN("zswhq"), ActN("refund"), Refund{})
	zsw.RegisterAction(AN("zswhq"), ActN("regproducer"), RegProducer{})
	zsw.RegisterAction(AN("zswhq"), ActN("unregprod"), UnregProducer{})
	zsw.RegisterAction(AN("zswhq"), ActN("regproxy"), RegProxy{})
	zsw.RegisterAction(AN("zswhq"), ActN("voteproducer"), VoteProducer{})
	zsw.RegisterAction(AN("zswhq"), ActN("claimrewards"), ClaimRewards{})
	zsw.RegisterAction(AN("zswhq"), ActN("buyram"), BuyRAM{})
	zsw.RegisterAction(AN("zswhq"), ActN("buyrambytes"), BuyRAMBytes{})
	zsw.RegisterAction(AN("zswhq"), ActN("linkauth"), LinkAuth{})
	zsw.RegisterAction(AN("zswhq"), ActN("unlinkauth"), UnlinkAuth{})
	zsw.RegisterAction(AN("zswhq"), ActN("deleteauth"), DeleteAuth{})
	zsw.RegisterAction(AN("zswhq"), ActN("rmvproducer"), RemoveProducer{})
	zsw.RegisterAction(AN("zswhq"), ActN("setprods"), SetProds{})
	zsw.RegisterAction(AN("zswhq"), ActN("setpriv"), SetPriv{})
	zsw.RegisterAction(AN("zswhq"), ActN("canceldelay"), CancelDelay{})
	zsw.RegisterAction(AN("zswhq"), ActN("bidname"), Bidname{})
	// zsw.RegisterAction(AN("zswhq"), ActN("nonce"), &Nonce{})
	zsw.RegisterAction(AN("zswhq"), ActN("sellram"), SellRAM{})
	zsw.RegisterAction(AN("zswhq"), ActN("updateauth"), UpdateAuth{})
	zsw.RegisterAction(AN("zswhq"), ActN("setramrate"), SetRAMRate{})
	zsw.RegisterAction(AN("zswhq"), ActN("setalimits"), Setalimits{})
}

var AN = zsw.AN
var PN = zsw.PN
var ActN = zsw.ActN
