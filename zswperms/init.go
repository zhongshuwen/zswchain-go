package zswperms

import zsw "github.com/zhongshuwen/zswchain-go"

func init() {
	zsw.RegisterAction(ZswPermsAN, ActN("addperms"), AddZswPerms{})
	zsw.RegisterAction(ZswPermsAN, ActN("rmperms"), RemoveZswPerms{})
	zsw.RegisterAction(ZswPermsAN, ActN("setperms"), SetZswPerms{})
}

var AN = zsw.AN
var PN = zsw.PN
var ActN = zsw.ActN

var ZswPermsAN = AN("zsw.perms")
