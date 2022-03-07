package forum

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewUnPost is an action undoing a post that is active
func NewUnPost(poster zsw.AccountName, postUUID string) *zsw.Action {
	a := &zsw.Action{
		Account: ForumAN,
		Name:    ActN("unpost"),
		Authorization: []zsw.PermissionLevel{
			{Actor: poster, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(UnPost{
			Poster:   poster,
			PostUUID: postUUID,
		}),
	}
	return a
}

// UnPost represents the `zswhq.forum::unpost` action.
type UnPost struct {
	Poster   zsw.AccountName `json:"poster"`
	PostUUID string          `json:"post_uuid"`
}
