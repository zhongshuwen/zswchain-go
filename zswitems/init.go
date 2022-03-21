package zswitems

import zsw "github.com/zhongshuwen/zswchain-go"

func init() {
	zsw.RegisterAction(ZswItemsAN, ActN("mint"), ItemMint{})
	zsw.RegisterAction(ZswItemsAN, ActN("setcustperms"), SetCustodianPermissions{})
	zsw.RegisterAction(ZswItemsAN, ActN("setuserperms"), SetUserPermissions{})
	zsw.RegisterAction(ZswItemsAN, ActN("mkcollection"), MakeCollection{})
	zsw.RegisterAction(ZswItemsAN, ActN("mkcustodian"), MakeCustodian{})
	zsw.RegisterAction(ZswItemsAN, ActN("mkissuer"), MakeIssuer{})
	zsw.RegisterAction(ZswItemsAN, ActN("mkitem"), MakeItem{})
	zsw.RegisterAction(ZswItemsAN, ActN("mkschema"), MakeSchema{})
	zsw.RegisterAction(ZswItemsAN, ActN("mkitemtpl"), MakeItemTemplate{})
	zsw.RegisterAction(ZswItemsAN, ActN("mkroyaltyusr"), MakeRoyaltyUser{})
	zsw.RegisterAction(ZswItemsAN, ActN("setcustperms"), SetCustodianPermissions{})
	zsw.RegisterAction(ZswItemsAN, ActN("setuserperms"), SetUserPermissions{})
	zsw.RegisterAction(ZswItemsAN, ActN("transfer"), ItemTransfer{})
}

var AN = zsw.AN
var PN = zsw.PN
var ActN = zsw.ActN

var ZswItemsAN = AN("zsw.items")
