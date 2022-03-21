package zswitems

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewMakeSchema(authorizer zsw.AccountName, creator zsw.AccountName, schemaName zsw.AccountName, schemaFormat []zsw.ZswItemsFormat) *zsw.Action {
	return &zsw.Action{
		Account: AN("zsw.items"),
		Name:    ActN("mkschema"),
		Authorization: []zsw.PermissionLevel{
			{Actor: authorizer, Permission: PN("active")},
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(MakeSchema{
			Authorizer:   authorizer,
			Creator:      creator,
			SchemaName:   schemaName,
			SchemaFormat: schemaFormat,
		}),
	}
}

type MakeSchema struct {
	Authorizer   zsw.AccountName      `json:"authorizer"`
	Creator      zsw.AccountName      `json:"creator"`
	SchemaName   zsw.AccountName      `json:"schema_name"`
	SchemaFormat []zsw.ZswItemsFormat `json:"schema_format"`
}
