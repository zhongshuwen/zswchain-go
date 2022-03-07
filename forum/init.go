package forum

import zsw "github.com/zhongshuwen/zswchain-go"

func init() {
	zsw.RegisterAction(ForumAN, ActN("clnproposal"), CleanProposal{})
	zsw.RegisterAction(ForumAN, ActN("expire"), Expire{})
	zsw.RegisterAction(ForumAN, ActN("post"), Post{})
	zsw.RegisterAction(ForumAN, ActN("propose"), Propose{})
	zsw.RegisterAction(ForumAN, ActN("status"), Status{})
	zsw.RegisterAction(ForumAN, ActN("unpost"), UnPost{})
	zsw.RegisterAction(ForumAN, ActN("unvote"), UnVote{})
	zsw.RegisterAction(ForumAN, ActN("vote"), Vote{})
}

var AN = zsw.AN
var PN = zsw.PN
var ActN = zsw.ActN

var ForumAN = AN("zswhq.forum")
