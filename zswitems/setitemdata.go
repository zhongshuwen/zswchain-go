package zswitems

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewSetItemData(authorizedEditor zsw.AccountName, itemId uint64, newMutableData zsw.AttributeMap) *zsw.Action {
	return &zsw.Action{
		Account: AN("zsw.items"),
		Name:    ActN("setitemdata"),
		Authorization: []zsw.PermissionLevel{
			{Actor: authorizedEditor, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(SetItemData{
			AuthorizedEditor: authorizedEditor,
			ItemId:           itemId,
			NewMutableData:   newMutableData,
		}),
	}
}

type SetItemData struct {
	AuthorizedEditor zsw.AccountName  `json:"authorized_editor"`
	ItemId           uint64           `json:"item_id"`
	NewMutableData   zsw.AttributeMap `json:"new_mutable_data"`
}
