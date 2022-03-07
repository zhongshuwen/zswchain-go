package zswitems

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)
func NewMakeItem(authorizer zsw.AccountName, creator zsw.AccountName, authorizedMinter zsw.AccountName, itemId uint64, zswId zsw.Uint128, itemConfig uint32, collectionId uint64, maxSupply uint64, itemType uint32, externalMetadataUrl string, schemaName zsw.AccountName, metadata zsw.ZswItemsMetadata) *zsw.Action {
	return &zsw.Action{
		Account: AN("zsw.items"),
		Name:    ActN("mkitem"),
		Authorization: []zsw.PermissionLevel{
			{Actor: authorizer, Permission: PN("active")},
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(MakeItem{
			Authorizer: authorizer,
			Creator: creator,
			AuthorizedMinter: authorizedMinter,
			ItemId: itemId,
			ZswId: zswId,
			ItemConfig: itemConfig,
			CollectionId: collectionId,
			MaxSupply: maxSupply,
			ItemType: itemType,
			ExternalMetadataUrl: externalMetadataUrl,
			SchemaName: schemaName,
			Metadata: metadata,
		}),
	}
}


type MakeItem struct {
  Authorizer zsw.AccountName `json:"authorizer"`
  Creator zsw.AccountName `json:"creator"`
  AuthorizedMinter zsw.AccountName `json:"authorized_minter"`
  ItemId uint64 `json:"item_id"`
  ZswId zsw.Uint128 `json:"zsw_id"`
  ItemConfig uint32 `json:"item_config"`
  CollectionId uint64 `json:"collection_id"`
  MaxSupply uint64 `json:"max_supply"`
  ItemType uint32 `json:"item_type"`
  ExternalMetadataUrl string `json:"external_metadata_url"`
  SchemaName zsw.AccountName `json:"schema_name"`
  Metadata zsw.ZswItemsMetadata `json:"metadata"`
}