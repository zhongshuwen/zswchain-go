package zswitems

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)
func NewMakeCollection(authorizer zsw.AccountName, creator zsw.AccountName, issuingPlatform zsw.AccountName, collectionId uint64, zswCode uint64, collectionType uint32, itemConfig uint32, secondaryMarketFee uint16, primaryMarketFee uint16, schemaName zsw.AccountName, externalMetadataUrl string, royaltyFeeCollector zsw.AccountName, notifyAccounts []zsw.AccountName, metadata zsw.ZswItemsMetadata) *zsw.Action {
	return &zsw.Action{
		Account: AN("zsw.items"),
		Name:    ActN("mkcollection"),
		Authorization: []zsw.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
			{Actor: authorizer, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(MakeCollection{
			Authorizer: authorizer,
			Creator: creator,
			IssuingPlatform: issuingPlatform,
			CollectionId: collectionId,
			ZswCode: zswCode,
			CollectionType: collectionType,
			ItemConfig: itemConfig,
			SecondaryMarketFee: secondaryMarketFee,
			PrimaryMarketFee: primaryMarketFee,
			SchemaName: schemaName,
			ExternalMetadataUrl: externalMetadataUrl,
			RoyaltyFeeCollector: royaltyFeeCollector,
			NotifyAccounts: notifyAccounts,
			Metadata: metadata,
		}),
	}
}


type MakeCollection struct {
  Authorizer zsw.AccountName `json:"authorizer"`
  Creator zsw.AccountName `json:"creator"`
  IssuingPlatform zsw.AccountName `json:"issuing_platform"`
  CollectionId uint64 `json:"collection_id"`
  ZswCode uint64 `json:"zsw_code"`
  CollectionType uint32 `json:"collection_type"`
  ItemConfig uint32 `json:"item_config"`
  SecondaryMarketFee uint16 `json:"secondary_market_fee"`
  PrimaryMarketFee uint16 `json:"primary_market_fee"`
  SchemaName zsw.AccountName `json:"schema_name"`
  ExternalMetadataUrl string `json:"external_metadata_url"`
  RoyaltyFeeCollector zsw.AccountName `json:"royalty_fee_collector"`
  NotifyAccounts []zsw.AccountName `json:"notify_accounts"`
  Metadata zsw.ZswItemsMetadata `json:"metadata"`
}