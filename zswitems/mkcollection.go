package zswitems

import (
	zsw "github.com/zhongshuwen/zswchain-go"
	"github.com/zhongshuwen/zswchain-go/zswattributes"
)

func NewMakeCollection(authorizer zsw.AccountName, zswId zsw.Uint128, collectionId uint64, collectionType uint32, creator zsw.AccountName, issuingPlatform zsw.AccountName, itemConfig uint32, secondaryMarketFee uint16, primaryMarketFee uint16, royaltyFeeCollector zsw.AccountName, maxSupply uint64, maxItems uint64, maxSupplyPerItem uint64, schemaName zsw.AccountName, authorizedMinters []zsw.AccountName, notifyAccounts []zsw.AccountName, authorizedMutableDataEditors []zsw.AccountName, metadata zswattributes.AttributeMap, externalMetadataUrl string) *zsw.Action {
	return &zsw.Action{
		Account: AN("zsw.items"),
		Name:    ActN("mkcollection"),
		Authorization: []zsw.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
			{Actor: authorizer, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(MakeCollection{
			Authorizer:                   authorizer,
			ZswId:                        zswId,
			CollectionId:                 collectionId,
			CollectionType:               collectionType,
			Creator:                      creator,
			IssuingPlatform:              issuingPlatform,
			ItemConfig:                   itemConfig,
			SecondaryMarketFee:           secondaryMarketFee,
			PrimaryMarketFee:             primaryMarketFee,
			RoyaltyFeeCollector:          royaltyFeeCollector,
			MaxSupply:                    maxSupply,
			MaxItems:                     maxItems,
			MaxSupplyPerItem:             maxSupplyPerItem,
			SchemaName:                   schemaName,
			AuthorizedMinters:            authorizedMinters,
			NotifyAccounts:               notifyAccounts,
			AuthorizedMutableDataEditors: authorizedMutableDataEditors,
			Metadata:                     metadata,
			ExternalMetadataUrl:          externalMetadataUrl,
		}),
	}
}

type MakeCollection struct {
	Authorizer                   zsw.AccountName            `json:"authorizer"`
	ZswId                        zsw.Uint128                `json:"zsw_id"`
	CollectionId                 uint64                     `json:"collection_id"`
	CollectionType               uint32                     `json:"collection_type"`
	Creator                      zsw.AccountName            `json:"creator"`
	IssuingPlatform              zsw.AccountName            `json:"issuing_platform"`
	ItemConfig                   uint32                     `json:"item_config"`
	SecondaryMarketFee           uint16                     `json:"secondary_market_fee"`
	PrimaryMarketFee             uint16                     `json:"primary_market_fee"`
	RoyaltyFeeCollector          zsw.AccountName            `json:"royalty_fee_collector"`
	MaxSupply                    uint64                     `json:"max_supply"`
	MaxItems                     uint64                     `json:"max_items"`
	MaxSupplyPerItem             uint64                     `json:"max_supply_per_item"`
	SchemaName                   zsw.Name                   `json:"schema_name"`
	AuthorizedMinters            []zsw.AccountName          `json:"authorized_minters"`
	NotifyAccounts               []zsw.AccountName          `json:"notify_accounts"`
	AuthorizedMutableDataEditors []zsw.AccountName          `json:"authorized_mutable_data_editors"`
	Metadata                     zswattributes.AttributeMap `json:"metadata"`
	ExternalMetadataUrl          string                     `json:"external_metadata_url"`
}
