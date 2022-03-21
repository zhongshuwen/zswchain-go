package zswitems

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewMakeItemTemplate(authorizer zsw.AccountName, creator zsw.AccountName, zswId zsw.Uint128, itemTemplateId uint64, collectionId uint64, itemType uint32, schemaName zsw.Name, immutableMetadata []zsw.ZswItemsMetadataKV, itemExternalMetadataUrlTemplate string) *zsw.Action {
	return &zsw.Action{
		Account: AN("zsw.items"),
		Name:    ActN("mkitemtpl"),
		Authorization: []zsw.PermissionLevel{
			{Actor: authorizer, Permission: PN("active")},
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(MakeItemTemplate{
			Authorizer:                      authorizer,
			Creator:                         creator,
			ZswId:                           zswId,
			ItemTemplateId:                  itemTemplateId,
			CollectionId:                    collectionId,
			ItemType:                        itemType,
			SchemaName:                      schemaName,
			ImmutableMetadata:               immutableMetadata,
			ItemExternalMetadataUrlTemplate: itemExternalMetadataUrlTemplate,
		}),
	}
}

type MakeItemTemplate struct {
	Authorizer                      zsw.AccountName          `json:"authorizer"`
	Creator                         zsw.AccountName          `json:"creator"`
	ZswId                           zsw.Uint128              `json:"zsw_id"`
	ItemTemplateId                  uint64                   `json:"item_template_id"`
	CollectionId                    uint64                   `json:"collection_id"`
	ItemType                        uint32                   `json:"item_type"`
	SchemaName                      zsw.Name                 `json:"schema_name"`
	ImmutableMetadata               []zsw.ZswItemsMetadataKV `json:"immutable_metadata"`
	ItemExternalMetadataUrlTemplate string                   `json:"item_external_metadata_url_template"`
}
