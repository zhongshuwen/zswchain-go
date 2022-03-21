package zswitems

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewMakeItem(authorizer zsw.AccountName, authorizedMinter zsw.AccountName, itemId uint64, zswId zsw.Uint128, itemConfig uint32, itemTemplateId uint64, maxSupply uint64, schemaName zsw.Name, immutableMetadata []zsw.ZswItemsMetadataKV, mutableMetadata []zsw.ZswItemsMetadataKV) *zsw.Action {
	return &zsw.Action{
		Account: AN("zsw.items"),
		Name:    ActN("mkitem"),
		Authorization: []zsw.PermissionLevel{
			{Actor: authorizer, Permission: PN("active")},
			{Actor: authorizedMinter, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(MakeItem{
			Authorizer:        authorizer,
			AuthorizedMinter:  authorizedMinter,
			ItemId:            itemId,
			ZswId:             zswId,
			ItemConfig:        itemConfig,
			ItemTemplateId:    itemTemplateId,
			MaxSupply:         maxSupply,
			SchemaName:        schemaName,
			ImmutableMetadata: immutableMetadata,
			MutableMetadata:   mutableMetadata,
		}),
	}
}

type MakeItem struct {
	Authorizer        zsw.AccountName          `json:"authorizer"`
	AuthorizedMinter  zsw.AccountName          `json:"authorized_minter"`
	ItemId            uint64                   `json:"item_id"`
	ZswId             zsw.Uint128              `json:"zsw_id"`
	ItemConfig        uint32                   `json:"item_config"`
	ItemTemplateId    uint64                   `json:"item_template_id"`
	MaxSupply         uint64                   `json:"max_supply"`
	SchemaName        zsw.Name                 `json:"schema_name"`
	ImmutableMetadata []zsw.ZswItemsMetadataKV `json:"immutable_metadata"`
	MutableMetadata   []zsw.ZswItemsMetadataKV `json:"mutable_metadata"`
}
