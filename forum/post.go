package forum

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// NewPost is an action representing a simple message to be posted
// through the chain network.
func NewPost(poster zsw.AccountName, postUUID, content string, replyToPoster zsw.AccountName, replyToPostUUID string, certify bool, jsonMetadata string) *zsw.Action {
	a := &zsw.Action{
		Account: ForumAN,
		Name:    ActN("post"),
		Authorization: []zsw.PermissionLevel{
			{Actor: poster, Permission: zsw.PermissionName("active")},
		},
		ActionData: zsw.NewActionData(Post{
			Poster:          poster,
			PostUUID:        postUUID,
			Content:         content,
			ReplyToPoster:   replyToPoster,
			ReplyToPostUUID: replyToPostUUID,
			Certify:         certify,
			JSONMetadata:    jsonMetadata,
		}),
	}
	return a
}

// Post represents the `eosio.forum::post` action.
type Post struct {
	Poster          zsw.AccountName `json:"poster"`
	PostUUID        string          `json:"post_uuid"`
	Content         string          `json:"content"`
	ReplyToPoster   zsw.AccountName `json:"reply_to_poster"`
	ReplyToPostUUID string          `json:"reply_to_post_uuid"`
	Certify         bool            `json:"certify"`
	JSONMetadata    string          `json:"json_metadata"`
}
