package zsw

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
	"github.com/zhongshuwen/zswchain-go/ecc"
)

var abiString = `
{
	"version": "eosio::abi/1.0",
	"types": [{
		"new_type_name": "new_type_name_1",
		"type": "name"
	}],
	"structs": [
	{
		"name": "struct_name_1",
		"base": "struct_name_2",
		"fields": [
			{"name":"struct_1_field_1", "type":"new_type_name_1"},
			{"name":"struct_1_field_2", "type":"struct_name_3"},
			{"name":"struct_1_field_3", "type":"string?"},
			{"name":"struct_1_field_4", "type":"string?"},
			{"name":"struct_1_field_5", "type":"struct_name_4[]"}
		]
   },{
		"name": "struct_name_2",
		"base": "",
		"fields": [
			{"name":"struct_2_field_1", "type":"string"}
		]
   },{
		"name": "struct_name_3",
		"base": "",
		"fields": [
			{"name":"struct_3_field_1", "type":"string"}
		]
   },{
		"name": "struct_name_4",
		"base": "",
		"fields": [
			{"name":"struct_4_field_1", "type":"string"}
		]
   }
	],
  "actions": [{
		"name": "action_name_1",
		"type": "struct_name_1",
		"ricardian_contract": ""
  }],
  "tables": [{
      "name": "table_name_1",
      "index_type": "i64",
      "key_names": [
        "key_name_1"
      ],
      "key_types": [
        "string"
      ],
      "type": "struct_name_1"
    }
  ]
}
`

var abiData = []byte(`{
	"struct_2_field_1": "struct_2_field_1_value",
	"struct_1_field_1": Name("zhongshuwen1"),
	"struct_1_field_2": M{
		"struct_3_field_1": "struct_3_field_1_value",
	},
	"struct_1_field_3": "struct_1_field_3_value",
	//"struct_1_field_4": "struct_1_field_4_value",
	"struct_1_field_5": ["struct_1_field_5_value_1","struct_1_field_5_value_2"],
}`)
var cangPinAbi = (`{
	"version": "eosio::abi/1.2",
	"types": [
		 {
				"new_type_name": "ATTRIBUTE_MAP",
				"type": "pair_string_SERDATA_ATTRIBUTE[]"
		 },
		 {
				"new_type_name": "DOUBLE_VEC",
				"type": "float64[]"
		 },
		 {
				"new_type_name": "FLOAT_VEC",
				"type": "float32[]"
		 },
		 {
				"new_type_name": "INT16_VEC",
				"type": "int16[]"
		 },
		 {
				"new_type_name": "INT32_VEC",
				"type": "int32[]"
		 },
		 {
				"new_type_name": "INT64_VEC",
				"type": "int64[]"
		 },
		 {
				"new_type_name": "INT8_VEC",
				"type": "bytes"
		 },
		 {
				"new_type_name": "SERDATA_ATTRIBUTE",
				"type": "variant_int8_int16_int32_int64_uint8_uint16_uint32_uint64_float32_float64_string_INT8_VEC_INT16_VEC_INT32_VEC_INT64_VEC_UINT8_VEC_UINT16_VEC_UINT32_VEC_UINT64_VEC_FLOAT_VEC_DOUBLE_VEC_STRING_VEC"
		 },
		 {
				"new_type_name": "STRING_VEC",
				"type": "string[]"
		 },
		 {
				"new_type_name": "UINT16_VEC",
				"type": "uint16[]"
		 },
		 {
				"new_type_name": "UINT32_VEC",
				"type": "uint32[]"
		 },
		 {
				"new_type_name": "UINT64_VEC",
				"type": "uint64[]"
		 },
		 {
				"new_type_name": "UINT8_VEC",
				"type": "bytes"
		 }
	],
	"structs": [
		 {
				"name": "FORMAT",
				"base": "",
				"fields": [
					 {
							"name": "name",
							"type": "string"
					 },
					 {
							"name": "type",
							"type": "string"
					 }
				]
		 },
		 {
				"name": "init",
				"base": "",
				"fields": [
					 {
							"name": "initializer",
							"type": "name"
					 }
				]
		 },
		 {
				"name": "logmint",
				"base": "",
				"fields": [
					 {
							"name": "minter",
							"type": "name"
					 },
					 {
							"name": "collection_id",
							"type": "uint64"
					 },
					 {
							"name": "to",
							"type": "name"
					 },
					 {
							"name": "to_custodian",
							"type": "name"
					 },
					 {
							"name": "item_ids",
							"type": "uint64[]"
					 },
					 {
							"name": "amounts",
							"type": "uint64[]"
					 },
					 {
							"name": "memo",
							"type": "string"
					 }
				]
		 },
		 {
				"name": "logmkitem",
				"base": "",
				"fields": [
					 {
							"name": "authorizer",
							"type": "name"
					 },
					 {
							"name": "authorized_minter",
							"type": "name"
					 },
					 {
							"name": "item_id",
							"type": "uint64"
					 },
					 {
							"name": "zsw_id",
							"type": "uint128"
					 },
					 {
							"name": "item_config",
							"type": "uint32"
					 },
					 {
							"name": "item_template_id",
							"type": "uint64"
					 },
					 {
							"name": "max_supply",
							"type": "uint64"
					 },
					 {
							"name": "schema_name",
							"type": "name"
					 },
					 {
							"name": "collection_id",
							"type": "uint64"
					 },
					 {
							"name": "immutable_metadata",
							"type": "ATTRIBUTE_MAP"
					 },
					 {
							"name": "mutable_metadata",
							"type": "ATTRIBUTE_MAP"
					 },
					 {
							"name": "immutable_template_data",
							"type": "ATTRIBUTE_MAP"
					 }
				]
		 },
		 {
				"name": "logmkitemtpl",
				"base": "",
				"fields": [
					 {
							"name": "authorizer",
							"type": "name"
					 },
					 {
							"name": "creator",
							"type": "name"
					 },
					 {
							"name": "zsw_id",
							"type": "uint128"
					 },
					 {
							"name": "item_template_id",
							"type": "uint64"
					 },
					 {
							"name": "collection_id",
							"type": "uint64"
					 },
					 {
							"name": "item_type",
							"type": "uint32"
					 },
					 {
							"name": "schema_name",
							"type": "name"
					 },
					 {
							"name": "immutable_metadata",
							"type": "ATTRIBUTE_MAP"
					 },
					 {
							"name": "item_external_metadata_url_template",
							"type": "string"
					 }
				]
		 },
		 {
				"name": "logsetdata",
				"base": "",
				"fields": [
					 {
							"name": "authorized_editor",
							"type": "name"
					 },
					 {
							"name": "item_id",
							"type": "uint64"
					 },
					 {
							"name": "collection_id",
							"type": "uint64"
					 },
					 {
							"name": "old_data",
							"type": "ATTRIBUTE_MAP"
					 },
					 {
							"name": "new_data",
							"type": "ATTRIBUTE_MAP"
					 }
				]
		 },
		 {
				"name": "logtransfer",
				"base": "",
				"fields": [
					 {
							"name": "authorizer",
							"type": "name"
					 },
					 {
							"name": "collection_id",
							"type": "uint64"
					 },
					 {
							"name": "from",
							"type": "name"
					 },
					 {
							"name": "to",
							"type": "name"
					 },
					 {
							"name": "from_custodian",
							"type": "name"
					 },
					 {
							"name": "to_custodian",
							"type": "name"
					 },
					 {
							"name": "item_ids",
							"type": "uint64[]"
					 },
					 {
							"name": "amounts",
							"type": "uint64[]"
					 },
					 {
							"name": "memo",
							"type": "string"
					 }
				]
		 },
		 {
				"name": "mint",
				"base": "",
				"fields": [
					 {
							"name": "minter",
							"type": "name"
					 },
					 {
							"name": "to",
							"type": "name"
					 },
					 {
							"name": "to_custodian",
							"type": "name"
					 },
					 {
							"name": "item_ids",
							"type": "uint64[]"
					 },
					 {
							"name": "amounts",
							"type": "uint64[]"
					 },
					 {
							"name": "memo",
							"type": "string"
					 },
					 {
							"name": "freeze_time",
							"type": "uint32"
					 }
				]
		 },
		 {
				"name": "mkcollection",
				"base": "",
				"fields": [
					 {
							"name": "authorizer",
							"type": "name"
					 },
					 {
							"name": "zsw_id",
							"type": "uint128"
					 },
					 {
							"name": "collection_id",
							"type": "uint64"
					 },
					 {
							"name": "collection_type",
							"type": "uint32"
					 },
					 {
							"name": "creator",
							"type": "name"
					 },
					 {
							"name": "issuing_platform",
							"type": "name"
					 },
					 {
							"name": "item_config",
							"type": "uint32"
					 },
					 {
							"name": "secondary_market_fee",
							"type": "uint16"
					 },
					 {
							"name": "primary_market_fee",
							"type": "uint16"
					 },
					 {
							"name": "royalty_fee_collector",
							"type": "name"
					 },
					 {
							"name": "max_supply",
							"type": "uint64"
					 },
					 {
							"name": "max_items",
							"type": "uint64"
					 },
					 {
							"name": "max_supply_per_item",
							"type": "uint64"
					 },
					 {
							"name": "schema_name",
							"type": "name"
					 },
					 {
							"name": "authorized_minters",
							"type": "name[]"
					 },
					 {
							"name": "notify_accounts",
							"type": "name[]"
					 },
					 {
							"name": "authorized_mutable_data_editors",
							"type": "name[]"
					 },
					 {
							"name": "metadata",
							"type": "ATTRIBUTE_MAP"
					 },
					 {
							"name": "external_metadata_url",
							"type": "string"
					 }
				]
		 },
		 {
				"name": "mkcustodian",
				"base": "",
				"fields": [
					 {
							"name": "creator",
							"type": "name"
					 },
					 {
							"name": "custodian_name",
							"type": "name"
					 },
					 {
							"name": "zsw_id",
							"type": "uint128"
					 },
					 {
							"name": "alt_id",
							"type": "uint128"
					 },
					 {
							"name": "permissions",
							"type": "uint128"
					 },
					 {
							"name": "status",
							"type": "uint32"
					 },
					 {
							"name": "incoming_freeze_period",
							"type": "uint32"
					 },
					 {
							"name": "notify_accounts",
							"type": "name[]"
					 }
				]
		 },
		 {
				"name": "mkissuer",
				"base": "",
				"fields": [
					 {
							"name": "authorizer",
							"type": "name"
					 },
					 {
							"name": "issuer_name",
							"type": "name"
					 },
					 {
							"name": "zsw_id",
							"type": "uint128"
					 },
					 {
							"name": "alt_id",
							"type": "uint128"
					 },
					 {
							"name": "permissions",
							"type": "uint128"
					 },
					 {
							"name": "status",
							"type": "uint32"
					 }
				]
		 },
		 {
				"name": "mkitem",
				"base": "",
				"fields": [
					 {
							"name": "authorizer",
							"type": "name"
					 },
					 {
							"name": "authorized_minter",
							"type": "name"
					 },
					 {
							"name": "item_id",
							"type": "uint64"
					 },
					 {
							"name": "zsw_id",
							"type": "uint128"
					 },
					 {
							"name": "item_config",
							"type": "uint32"
					 },
					 {
							"name": "item_template_id",
							"type": "uint64"
					 },
					 {
							"name": "max_supply",
							"type": "uint64"
					 },
					 {
							"name": "schema_name",
							"type": "name"
					 },
					 {
							"name": "immutable_metadata",
							"type": "ATTRIBUTE_MAP"
					 },
					 {
							"name": "mutable_metadata",
							"type": "ATTRIBUTE_MAP"
					 }
				]
		 },
		 {
				"name": "mkitemtpl",
				"base": "",
				"fields": [
					 {
							"name": "authorizer",
							"type": "name"
					 },
					 {
							"name": "creator",
							"type": "name"
					 },
					 {
							"name": "zsw_id",
							"type": "uint128"
					 },
					 {
							"name": "item_template_id",
							"type": "uint64"
					 },
					 {
							"name": "collection_id",
							"type": "uint64"
					 },
					 {
							"name": "item_type",
							"type": "uint32"
					 },
					 {
							"name": "schema_name",
							"type": "name"
					 },
					 {
							"name": "immutable_metadata",
							"type": "ATTRIBUTE_MAP"
					 },
					 {
							"name": "item_external_metadata_url_template",
							"type": "string"
					 }
				]
		 },
		 {
				"name": "mkroyaltyusr",
				"base": "",
				"fields": [
					 {
							"name": "authorizer",
							"type": "name"
					 },
					 {
							"name": "newroyaltyusr",
							"type": "name"
					 },
					 {
							"name": "zsw_id",
							"type": "uint128"
					 },
					 {
							"name": "alt_id",
							"type": "uint128"
					 },
					 {
							"name": "status",
							"type": "uint32"
					 }
				]
		 },
		 {
				"name": "mkschema",
				"base": "",
				"fields": [
					 {
							"name": "authorizer",
							"type": "name"
					 },
					 {
							"name": "creator",
							"type": "name"
					 },
					 {
							"name": "schema_name",
							"type": "name"
					 },
					 {
							"name": "schema_format",
							"type": "FORMAT[]"
					 }
				]
		 },
		 {
				"name": "pair_string_SERDATA_ATTRIBUTE",
				"base": "",
				"fields": [
					 {
							"name": "key",
							"type": "string"
					 },
					 {
							"name": "value",
							"type": "SERDATA_ATTRIBUTE"
					 }
				]
		 },
		 {
				"name": "s_collections",
				"base": "",
				"fields": [
					 {
							"name": "zsw_id",
							"type": "uint128"
					 },
					 {
							"name": "collection_type",
							"type": "uint32"
					 },
					 {
							"name": "creator",
							"type": "name"
					 },
					 {
							"name": "issuing_platform",
							"type": "name"
					 },
					 {
							"name": "item_config",
							"type": "uint32"
					 },
					 {
							"name": "secondary_market_fee",
							"type": "uint16"
					 },
					 {
							"name": "primary_market_fee",
							"type": "uint16"
					 },
					 {
							"name": "royalty_fee_collector",
							"type": "name"
					 },
					 {
							"name": "issued_supply",
							"type": "uint64"
					 },
					 {
							"name": "max_supply",
							"type": "uint64"
					 },
					 {
							"name": "items_count",
							"type": "uint64"
					 },
					 {
							"name": "max_items",
							"type": "uint64"
					 },
					 {
							"name": "max_supply_per_item",
							"type": "uint64"
					 },
					 {
							"name": "schema_name",
							"type": "name"
					 },
					 {
							"name": "authorized_minters",
							"type": "name[]"
					 },
					 {
							"name": "notify_accounts",
							"type": "name[]"
					 },
					 {
							"name": "authorized_mutable_data_editors",
							"type": "name[]"
					 },
					 {
							"name": "serialized_metadata",
							"type": "bytes"
					 },
					 {
							"name": "external_metadata_url",
							"type": "string"
					 }
				]
		 },
		 {
				"name": "s_custodian_user_pairs",
				"base": "",
				"fields": [
					 {
							"name": "custodian_user_pair_id",
							"type": "uint64"
					 },
					 {
							"name": "user",
							"type": "name"
					 },
					 {
							"name": "custodian",
							"type": "name"
					 }
				]
		 },
		 {
				"name": "s_custodians",
				"base": "",
				"fields": [
					 {
							"name": "custodian_id",
							"type": "uint64"
					 },
					 {
							"name": "custodian_name",
							"type": "name"
					 },
					 {
							"name": "zsw_id",
							"type": "uint128"
					 },
					 {
							"name": "alt_id",
							"type": "uint128"
					 },
					 {
							"name": "permissions",
							"type": "uint128"
					 },
					 {
							"name": "status",
							"type": "uint32"
					 },
					 {
							"name": "incoming_freeze_period",
							"type": "uint32"
					 },
					 {
							"name": "notify_accounts",
							"type": "name[]"
					 }
				]
		 },
		 {
				"name": "s_frozen_balances",
				"base": "",
				"fields": [
					 {
							"name": "frozen_balance_id",
							"type": "uint64"
					 },
					 {
							"name": "balance",
							"type": "uint64"
					 },
					 {
							"name": "status",
							"type": "uint32"
					 }
				]
		 },
		 {
				"name": "s_issuerstatus",
				"base": "",
				"fields": [
					 {
							"name": "issuer_name",
							"type": "name"
					 },
					 {
							"name": "zsw_id",
							"type": "uint128"
					 },
					 {
							"name": "alt_id",
							"type": "uint128"
					 },
					 {
							"name": "permissions",
							"type": "uint128"
					 },
					 {
							"name": "status",
							"type": "uint32"
					 }
				]
		 },
		 {
				"name": "s_item_templates",
				"base": "",
				"fields": [
					 {
							"name": "zsw_id",
							"type": "uint128"
					 },
					 {
							"name": "collection_id",
							"type": "uint64"
					 },
					 {
							"name": "item_type",
							"type": "uint32"
					 },
					 {
							"name": "schema_name",
							"type": "name"
					 },
					 {
							"name": "serialized_immutable_metadata",
							"type": "bytes"
					 },
					 {
							"name": "item_external_metadata_url_template",
							"type": "string"
					 }
				]
		 },
		 {
				"name": "s_itembalances",
				"base": "",
				"fields": [
					 {
							"name": "item_id",
							"type": "uint64"
					 },
					 {
							"name": "status",
							"type": "uint32"
					 },
					 {
							"name": "balance_normal_liquid",
							"type": "uint64"
					 },
					 {
							"name": "balance_frozen",
							"type": "uint64"
					 },
					 {
							"name": "balance_in_custody_liquid",
							"type": "uint64"
					 },
					 {
							"name": "active_custodian_pairs",
							"type": "uint64[]"
					 }
				]
		 },
		 {
				"name": "s_items",
				"base": "",
				"fields": [
					 {
							"name": "zsw_id",
							"type": "uint128"
					 },
					 {
							"name": "item_config",
							"type": "uint32"
					 },
					 {
							"name": "item_template_id",
							"type": "uint64"
					 },
					 {
							"name": "collection_id",
							"type": "uint64"
					 },
					 {
							"name": "issued_supply",
							"type": "uint64"
					 },
					 {
							"name": "max_supply",
							"type": "uint64"
					 },
					 {
							"name": "schema_name",
							"type": "name"
					 },
					 {
							"name": "serialized_immutable_metadata",
							"type": "bytes"
					 },
					 {
							"name": "serialized_mutable_metadata",
							"type": "bytes"
					 }
				]
		 },
		 {
				"name": "s_royaltyusers",
				"base": "",
				"fields": [
					 {
							"name": "user_name",
							"type": "name"
					 },
					 {
							"name": "zsw_id",
							"type": "uint128"
					 },
					 {
							"name": "alt_id",
							"type": "uint128"
					 },
					 {
							"name": "status",
							"type": "uint32"
					 },
					 {
							"name": "reported_fees_3rd",
							"type": "uint64"
					 },
					 {
							"name": "reported_fees_zsw",
							"type": "uint64"
					 },
					 {
							"name": "uncollected_fees_3rd",
							"type": "uint64"
					 },
					 {
							"name": "uncollected_fees_zsw",
							"type": "uint64"
					 }
				]
		 },
		 {
				"name": "s_schemas",
				"base": "",
				"fields": [
					 {
							"name": "schema_name",
							"type": "name"
					 },
					 {
							"name": "format",
							"type": "FORMAT[]"
					 }
				]
		 },
		 {
				"name": "setcustperms",
				"base": "",
				"fields": [
					 {
							"name": "sender",
							"type": "name"
					 },
					 {
							"name": "custodian",
							"type": "name"
					 },
					 {
							"name": "permissions",
							"type": "uint128"
					 }
				]
		 },
		 {
				"name": "setitemdata",
				"base": "",
				"fields": [
					 {
							"name": "authorized_editor",
							"type": "name"
					 },
					 {
							"name": "item_id",
							"type": "uint64"
					 },
					 {
							"name": "new_mutable_data",
							"type": "ATTRIBUTE_MAP"
					 }
				]
		 },
		 {
				"name": "setuserperms",
				"base": "",
				"fields": [
					 {
							"name": "sender",
							"type": "name"
					 },
					 {
							"name": "user",
							"type": "name"
					 },
					 {
							"name": "permissions",
							"type": "uint128"
					 }
				]
		 },
		 {
				"name": "transfer",
				"base": "",
				"fields": [
					 {
							"name": "authorizer",
							"type": "name"
					 },
					 {
							"name": "from",
							"type": "name"
					 },
					 {
							"name": "to",
							"type": "name"
					 },
					 {
							"name": "from_custodian",
							"type": "name"
					 },
					 {
							"name": "to_custodian",
							"type": "name"
					 },
					 {
							"name": "freeze_time",
							"type": "uint32"
					 },
					 {
							"name": "use_liquid_backup",
							"type": "bool"
					 },
					 {
							"name": "max_unfreeze_iterations",
							"type": "uint32"
					 },
					 {
							"name": "item_ids",
							"type": "uint64[]"
					 },
					 {
							"name": "amounts",
							"type": "uint64[]"
					 },
					 {
							"name": "memo",
							"type": "string"
					 }
				]
		 }
	],
	"actions": [
		 {
				"name": "init",
				"type": "init",
				"ricardian_contract": ""
		 },
		 {
				"name": "logmint",
				"type": "logmint",
				"ricardian_contract": ""
		 },
		 {
				"name": "logmkitem",
				"type": "logmkitem",
				"ricardian_contract": ""
		 },
		 {
				"name": "logmkitemtpl",
				"type": "logmkitemtpl",
				"ricardian_contract": ""
		 },
		 {
				"name": "logsetdata",
				"type": "logsetdata",
				"ricardian_contract": ""
		 },
		 {
				"name": "logtransfer",
				"type": "logtransfer",
				"ricardian_contract": ""
		 },
		 {
				"name": "mint",
				"type": "mint",
				"ricardian_contract": ""
		 },
		 {
				"name": "mkcollection",
				"type": "mkcollection",
				"ricardian_contract": ""
		 },
		 {
				"name": "mkcustodian",
				"type": "mkcustodian",
				"ricardian_contract": ""
		 },
		 {
				"name": "mkissuer",
				"type": "mkissuer",
				"ricardian_contract": ""
		 },
		 {
				"name": "mkitem",
				"type": "mkitem",
				"ricardian_contract": ""
		 },
		 {
				"name": "mkitemtpl",
				"type": "mkitemtpl",
				"ricardian_contract": ""
		 },
		 {
				"name": "mkroyaltyusr",
				"type": "mkroyaltyusr",
				"ricardian_contract": ""
		 },
		 {
				"name": "mkschema",
				"type": "mkschema",
				"ricardian_contract": ""
		 },
		 {
				"name": "setcustperms",
				"type": "setcustperms",
				"ricardian_contract": ""
		 },
		 {
				"name": "setitemdata",
				"type": "setitemdata",
				"ricardian_contract": ""
		 },
		 {
				"name": "setuserperms",
				"type": "setuserperms",
				"ricardian_contract": ""
		 },
		 {
				"name": "transfer",
				"type": "transfer",
				"ricardian_contract": ""
		 }
	],
	"tables": [
		 {
				"name": "collections",
				"index_type": "i64",
				"type": "s_collections"
		 },
		 {
				"name": "custodians",
				"index_type": "i64",
				"type": "s_custodians"
		 },
		 {
				"name": "custodianups",
				"index_type": "i64",
				"type": "s_custodian_user_pairs"
		 },
		 {
				"name": "frozenbals",
				"index_type": "i64",
				"type": "s_frozen_balances"
		 },
		 {
				"name": "issuerstatus",
				"index_type": "i64",
				"type": "s_issuerstatus"
		 },
		 {
				"name": "itembalances",
				"index_type": "i64",
				"type": "s_itembalances"
		 },
		 {
				"name": "items",
				"index_type": "i64",
				"type": "s_items"
		 },
		 {
				"name": "itemtemplate",
				"index_type": "i64",
				"type": "s_item_templates"
		 },
		 {
				"name": "royaltyusers",
				"index_type": "i64",
				"type": "s_royaltyusers"
		 },
		 {
				"name": "schemas",
				"index_type": "i64",
				"type": "s_schemas"
		 }
	],
	"variants": [
		 {
				"name": "variant_int8_int16_int32_int64_uint8_uint16_uint32_uint64_float32_float64_string_INT8_VEC_INT16_VEC_INT32_VEC_INT64_VEC_UINT8_VEC_UINT16_VEC_UINT32_VEC_UINT64_VEC_FLOAT_VEC_DOUBLE_VEC_STRING_VEC",
				"types": [
					 "int8",
					 "int16",
					 "int32",
					 "int64",
					 "uint8",
					 "uint16",
					 "uint32",
					 "uint64",
					 "float32",
					 "float64",
					 "string",
					 "INT8_VEC",
					 "INT16_VEC",
					 "INT32_VEC",
					 "INT64_VEC",
					 "UINT8_VEC",
					 "UINT16_VEC",
					 "UINT32_VEC",
					 "UINT64_VEC",
					 "FLOAT_VEC",
					 "DOUBLE_VEC",
					 "STRING_VEC"
				]
		 }
	]
}`)

func TestABIEncoder_Encode(t *testing.T) {

	testCases := []map[string]interface{}{
		{"caseName": "sunny path", "actionName": "action_name_1", "expectedError": nil, "abi": abiString},
		{"caseName": "missing action", "actionName": "bad_action_name", "expectedError": fmt.Errorf("encode action: action bad_action_name not found in abi"), "abi": abiString},
	}

	for _, c := range testCases {
		caseName := c["caseName"].(string)
		t.Run(caseName, func(t *testing.T) {

			abi, err := NewABI(strings.NewReader(c["abi"].(string)))
			require.NoError(t, err)

			_, err = abi.EncodeAction(ActionName(c["actionName"].(string)), abiData)
			assert.Equal(t, c["expectedError"], err)

			if c["expectedError"] != nil {
				return
			}

			//decoder := NewABIDecoder(buf.Bytes(), strings.NewReader(abiString))
			//result := make(M)
			//err = decoder.Decode(result, ActionName(c["actionName"].(string)))
			//require.NoError(t, err)

			//assert.Equal(t, abiData, result)
			//fmt.Println(result)
		})
	}
}

func TestABIEncoder_encodeMissingActionStruct(t *testing.T) {

	abiString := `
{
	"version": "eosio::abi/1.0",
	"types": [{
		"new_type_name": "new.type.name.1",
		"type": "name"
	}],
	"structs": [
	],
  "actions": [{
		"name": "action.name.1",
		"type": "struct.name.1",
		"ricardian_contract": ""
  }]
}
`

	abi, err := NewABI(strings.NewReader(abiString))
	require.NoError(t, err)

	_, err = abi.EncodeAction(ActionName("action.name.1"), abiData)
	assert.Equal(t, "encode action: encode struct [struct.name.1] not found in abi", err.Error())
}

/*
func TestABIEncoder_encodeMissingActionStructCangPin(t *testing.T) {

	actionData := `{
		"authorizer": "zsw.admin",
		"authorized_minter": "kxjdtestnew1",
		"item_id": 55,
		"zsw_id": "0x00000055",
		"item_config": 11,
		"item_template_id": 33,
		"max_supply": 9000000,
		"schema_name": "coolnametest",
		"immutable_metadata": [
			{"key": "name", "value": ["string", "My first item"]},
			{"key": "image_url", "value": ["string", "https://cangpin.test.chao7.cn/f/images/shanghai.png"]},
			{"key": "rarity", "value": ["uint32", 10]}
		],
		"mutable_metadata": []
	}`

	abi, err := NewABI(strings.NewReader(cangPinAbi))
	require.NoError(t, err)

	action, err := abi.EncodeAction(ActionName("mkitem"), []byte(actionData))

	assert.NoError(t, err, "error encoding action")
	decoded, err := abi.DecodeAction(action, ActionName("mkitem"))

	assert.NoError(t, err, "error encoding action")
	fmt.Println(string(decoded))
}
*/
func TestABIEncoder_encodeErrorInBase(t *testing.T) {

	abiString := `
{
	"version": "eosio::abi/1.0",
	"types": [{
		"new_type_name": "new.type.name.1",
		"type": "name"
	}],
	"structs": [
	{
		"name": "struct.name.1",
		"base": "struct.name.2",
		"fields": [
			{"name":"struct.1.field.1", "type":"new.type.name.1"}
		]
   }
	],
  "actions": [{
		"name": "action.name.1",
		"type": "struct.name.1",
		"ricardian_contract": ""
  }]
}
`

	abi, err := NewABI(strings.NewReader(abiString))
	require.NoError(t, err)

	_, err = abi.EncodeAction(ActionName("action.name.1"), abiData)
	assert.Equal(t, "encode action: encode base [struct.name.1]: encode struct [struct.name.2] not found in abi", err.Error())
}

func TestABIEncoder_encodeField(t *testing.T) {

	testCases := []map[string]interface{}{
		{"caseName": "sunny path", "fieldName": "field_name", "fieldType": "string", "expectedValue": "0f6669656c642e312e76616c75652e31", "json": "{\"field_name\": \"field.1.value.1\"}", "isOptional": false, "isArray": false, "expectedError": nil, "writer": new(bytes.Buffer)},
		{"caseName": "optional present", "fieldName": "field_name", "fieldType": "string", "expectedValue": "010f6669656c642e312e76616c75652e31", "json": "{\"field_name\": \"field.1.value.1\"}", "isOptional": true, "isArray": false, "expectedError": nil, "writer": new(bytes.Buffer)},
		{"caseName": "optional not present", "fieldName": "field_name", "fieldType": "string", "expectedValue": "00", "json": "{\"field_name_other\": \"field.1.value.2\"}", "isOptional": true, "isArray": false, "expectedError": nil, "writer": new(bytes.Buffer)},
		{"caseName": "optional present write flag err", "fieldName": "field_name", "fieldType": "string", "expectedValue": "010f6669656c642e312e76616c75652e31", "json": "{\"field_name\": \"field.1.value.1\"}", "isOptional": true, "isArray": false, "expectedError": fmt.Errorf("error.1"), "writer": mockWriter{err: fmt.Errorf("error.1")}},
		{"caseName": "not optional not present", "fieldName": "field_name", "fieldType": "string", "expectedValue": "00", "json": "{\"field_name_other\": \"field.1.value.2\"}", "isOptional": false, "isArray": false, "expectedError": fmt.Errorf("encode field: none optional field [field_name] as a nil value"), "writer": new(bytes.Buffer)},
		{"caseName": "array", "fieldName": "field_name", "fieldType": "string", "expectedValue": "020f6669656c642e312e76616c75652e310f6669656c642e312e76616c75652e32", "json": "{\"field_name\": [\"field.1.value.1\",\"field.1.value.2\"]}", "isOptional": false, "isArray": true, "expectedError": nil, "writer": new(bytes.Buffer)},
		{"caseName": "expected array got string", "fieldName": "field_name", "fieldType": "string", "expectedValue": "", "json": "{\"field_name\": \"field.1.value.1\"}", "isOptional": false, "isArray": true, "expectedError": fmt.Errorf("encode field: expected array for field [field_name] got [String]"), "writer": new(bytes.Buffer)},
	}

	for _, c := range testCases {
		caseName := c["caseName"].(string)
		t.Run(caseName, func(t *testing.T) {
			buf := c["writer"].(mockWriterable)
			encoder := NewEncoder(buf)

			abi := ABI{}

			json := c["json"].(string)
			fieldName := c["fieldName"].(string)
			fieldType := c["fieldType"].(string)
			isOptional := c["isOptional"].(bool)
			isArray := c["isArray"].(bool)

			err := abi.encodeField(encoder, fieldName, fieldType, isOptional, isArray, []byte(json))

			if c["expectedError"] == nil {
				require.Nil(t, err)
				assert.Equal(t, c["expectedValue"], hex.EncodeToString(buf.Bytes()), c["caseName"])
			} else {
				expectedError := c["expectedError"].(error)
				require.Equal(t, expectedError.Error(), err.Error(), c["caseName"])
			}
		})

	}
}

func TestABI_Write(t *testing.T) {
	testCases := []map[string]interface{}{
		{"caseName": "string", "typeName": "string", "expectedValue": "0e746869732e69732e612e74657374", "json": "{\"testField\":\"this.is.a.test\""},
		{"caseName": "min int8", "typeName": "int8", "expectedValue": "80", "json": "{\"testField\":-128}"},
		{"caseName": "max int8", "typeName": "int8", "expectedValue": "7f", "json": "{\"testField\":127}", "expectedError": nil},
		{"caseName": "out of range int8", "typeName": "int8", "expectedValue": "", "json": "{\"testField\":128}", "expectedError": fmt.Errorf("writing field: [test_field_name] type int8 : strconv.ParseInt: parsing \"128\": value out of range")},
		{"caseName": "out of range int8", "typeName": "int8", "expectedValue": "", "json": "{\"testField\":-129}", "expectedError": fmt.Errorf("writing field: [test_field_name] type int8 : strconv.ParseInt: parsing \"-129\": value out of range")},
		{"caseName": "min uint8", "typeName": "uint8", "expectedValue": "00", "json": "{\"testField\":0}", "expectedError": nil},
		{"caseName": "max uint8", "typeName": "uint8", "expectedValue": "ff", "json": "{\"testField\":255}", "expectedError": nil},
		{"caseName": "out of range uint8", "typeName": "uint8", "expectedValue": "", "json": "{\"testField\":-1}", "expectedError": fmt.Errorf("writing field: [test_field_name] type uint8 : strconv.ParseUint: parsing \"-1\": invalid syntax")},
		{"caseName": "out of range uint8", "typeName": "uint8", "expectedValue": "", "json": "{\"testField\":256}", "expectedError": fmt.Errorf("writing field: [test_field_name] type uint8 : strconv.ParseUint: parsing \"256\": value out of range")},
		{"caseName": "min int16", "typeName": "int16", "expectedValue": "0080", "json": "{\"testField\":-32768}", "expectedError": nil, "isOptional": false, "isArray": false, "fieldName": "testedField"},
		{"caseName": "max int16", "typeName": "int16", "expectedValue": "ff7f", "json": "{\"testField\":32767}", "expectedError": nil, "isOptional": false, "isArray": false, "fieldName": "testedField"},
		{"caseName": "out of range int16", "typeName": "int16", "expectedValue": "", "json": "{\"testField\":-32769}", "expectedError": fmt.Errorf("writing field: [test_field_name] type int16 : strconv.ParseInt: parsing \"-32769\": value out of range")},
		{"caseName": "out of range int16", "typeName": "int16", "expectedValue": "", "json": "{\"testField\":32768}", "expectedError": fmt.Errorf("writing field: [test_field_name] type int16 : strconv.ParseInt: parsing \"32768\": value out of range")},
		{"caseName": "min uint16", "typeName": "uint16", "expectedValue": "0000", "json": "{\"testField\":0}", "expectedError": nil, "isOptional": false, "isArray": false, "fieldName": "testedField"},
		{"caseName": "max uint16", "typeName": "uint16", "expectedValue": "ffff", "json": "{\"testField\":65535}", "expectedError": nil, "isOptional": false, "isArray": false, "fieldName": "testedField"},
		{"caseName": "out of range uint16", "typeName": "uint16", "expectedValue": "", "json": "{\"testField\":-1}", "expectedError": fmt.Errorf("writing field: [test_field_name] type uint16 : strconv.ParseUint: parsing \"-1\": invalid syntax")},
		{"caseName": "out of range uint16", "typeName": "uint16", "expectedValue": "", "json": "{\"testField\":65536}", "expectedError": fmt.Errorf("writing field: [test_field_name] type uint16 : strconv.ParseUint: parsing \"65536\": value out of range")},
		{"caseName": "min int32", "typeName": "int32", "expectedValue": "00000080", "json": "{\"testField\":-2147483648}", "expectedError": nil, "isOptional": false, "isArray": false, "fieldName": "testedField"},
		{"caseName": "max int32", "typeName": "int32", "expectedValue": "ffffff7f", "json": "{\"testField\":2147483647}", "expectedError": nil, "isOptional": false, "isArray": false, "fieldName": "testedField"},
		{"caseName": "out of range int32", "typeName": "int32", "expectedValue": "", "json": "{\"testField\":-2147483649}", "expectedError": fmt.Errorf("writing field: [test_field_name] type int32 : strconv.ParseInt: parsing \"-2147483649\": value out of range")},
		{"caseName": "out of range int32", "typeName": "int32", "expectedValue": "", "json": "{\"testField\":2147483648}", "expectedError": fmt.Errorf("writing field: [test_field_name] type int32 : strconv.ParseInt: parsing \"2147483648\": value out of range")},
		{"caseName": "min uint32", "typeName": "uint32", "expectedValue": "00000000", "json": "{\"testField\":0}", "expectedError": nil, "isOptional": false, "isArray": false, "fieldName": "testedField"},
		{"caseName": "max uint32", "typeName": "uint32", "expectedValue": "ffffffff", "json": "{\"testField\":4294967295}", "expectedError": nil, "isOptional": false, "isArray": false, "fieldName": "testedField"},
		{"caseName": "out of range uint32", "typeName": "uint32", "expectedValue": "", "json": "{\"testField\":-1}", "expectedError": fmt.Errorf("writing field: [test_field_name] type uint32 : strconv.ParseUint: parsing \"-1\": invalid syntax")},
		{"caseName": "out of range uint32", "typeName": "uint32", "expectedValue": "", "json": "{\"testField\":4294967296}", "expectedError": fmt.Errorf("writing field: [test_field_name] type uint32 : strconv.ParseUint: parsing \"4294967296\": value out of range")},
		{"caseName": "min int64", "typeName": "int64", "expectedValue": "0000000000000080", "json": "{\"testField\":\"-9223372036854775808\"}", "expectedError": nil, "isOptional": false, "isArray": false, "fieldName": "testedField"},
		{"caseName": "mid int64", "typeName": "int64", "expectedValue": "00f0ffffffffffff", "json": "{\"testField\":-4096}", "expectedError": nil, "isOptional": false, "isArray": false, "fieldName": "testedField"},
		{"caseName": "max int64", "typeName": "int64", "expectedValue": "ffffffffffffff7f", "json": "{\"testField\":\"9223372036854775807\"}", "expectedError": nil, "isOptional": false, "isArray": false, "fieldName": "testedField"},
		{"caseName": "out of range int64 lower", "typeName": "int64", "expectedValue": "", "json": "{\"testField\":-9223372036854775809}", "expectedError": fmt.Errorf("encoding int64: json: cannot unmarshal number -9223372036854775809 into Go value of type int64")},
		{"caseName": "out of range int64 upper", "typeName": "int64", "expectedValue": "", "json": "{\"testField\":9223372036854775808}", "expectedError": fmt.Errorf("encoding int64: json: cannot unmarshal number 9223372036854775808 into Go value of type int64")},
		{"caseName": "min uint64", "typeName": "uint64", "expectedValue": "0000000000000000", "json": "{\"testField\":0}", "expectedError": nil, "isOptional": false, "isArray": false, "fieldName": "testedField"},
		{"caseName": "mid uint64", "typeName": "uint64", "expectedValue": "c06ddb095f285813", "json": "{\"testField\":\"1393908473323548096\"}", "expectedError": nil, "isOptional": false, "isArray": false, "fieldName": "testedField"},
		{"caseName": "max uint64", "typeName": "uint64", "expectedValue": "ffffffffffffffff", "json": "{\"testField\":\"18446744073709551615\"}", "expectedError": nil, "isOptional": false, "isArray": false, "fieldName": "testedField"},
		{"caseName": "out of range uint64 lower", "typeName": "uint64", "expectedValue": "", "json": "{\"testField\":-1}", "expectedError": fmt.Errorf("encoding uint64: json: cannot unmarshal number -1 into Go value of type uint64")},
		{"caseName": "out of range uint64 upper", "typeName": "uint64", "expectedValue": "", "json": "{\"testField\":18446744073709551616}", "expectedError": fmt.Errorf("encoding uint64: json: cannot unmarshal number 18446744073709551616 into Go value of type uint64")},
		{"caseName": "int128", "typeName": "int128", "expectedValue": "01020000000000000200000000000000", "json": "{\"testField\":\"0x01020000000000000200000000000000\"}"},
		{"caseName": "uint128", "typeName": "uint128", "expectedValue": "01000000000000000200000000000000", "json": "{\"testField\":\"0x01000000000000000200000000000000\"}"},
		{"caseName": "varint32", "typeName": "varint32", "expectedValue": "ffffffff0f", "json": "{\"testField\":-2147483648}", "expectedError": nil, "isOptional": false, "isArray": false, "fieldName": "testedField"},
		{"caseName": "varuint32", "typeName": "varuint32", "expectedValue": "ffffffff0f", "json": "{\"testField\":4294967295}", "expectedError": nil, "isOptional": false, "isArray": false, "fieldName": "testedField"}, //{"caseName": "min varuint32", "typeName": "varuint32", "expectedValue": "0", "json": Varuint32(0), "expectedError": nil, "isOptional": false, "isArray": false, "fieldName": "testedField"},
		{"caseName": "min float32", "typeName": "float32", "expectedValue": "01000000", "json": "{\"testField\":0.000000000000000000000000000000000000000000001401298464324817}", "expectedError": nil},
		{"caseName": "max float32", "typeName": "float32", "expectedValue": "ffff7f7f", "json": "{\"testField\":340282346638528860000000000000000000000}", "expectedError": nil},
		{"caseName": "err float32", "typeName": "float32", "expectedValue": "ffff7f7f", "json": "{\"testField\":440282346638528860000000000000000000000}", "expectedError": fmt.Errorf("writing field: [test_field_name] type float32 : strconv.ParseFloat: parsing \"440282346638528860000000000000000000000\": value out of range")},
		{"caseName": "min float64", "typeName": "float64", "expectedValue": "0100000000000000", "json": "{\"testField\":0.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005}", "expectedError": nil},
		{"caseName": "max float64", "typeName": "float64", "expectedValue": "ffffffffffffef7f", "json": "{\"testField\":179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000}", "expectedError": nil},
		{"caseName": "err float64", "typeName": "float64", "expectedValue": "ffffffffffffef7f", "json": "{\"testField\":279769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000}", "expectedError": fmt.Errorf("writing field: [test_field_name] type float64 : strconv.ParseFloat: parsing \"279769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000\": value out of range")},
		{"caseName": "float128", "typeName": "float128", "expectedValue": "ffffffffffffef7fffffffffffffef7f", "json": "{\"testField\":\"0xffffffffffffef7fffffffffffffef7f\"}"},
		{"caseName": "bool true", "typeName": "bool", "expectedValue": "01", "json": "{\"testField\":true}", "expectedError": nil},
		{"caseName": "bool false", "typeName": "bool", "expectedValue": "00", "json": "{\"testField\":false}", "expectedError": nil},
		{"caseName": "time_point", "typeName": "time_point", "expectedValue": "0100000000000000", "json": "{\"testField\":\"1970-01-01T00:00:00.001\"", "expectedError": nil},
		{"caseName": "time_point err", "typeName": "time_point", "expectedValue": "0100000000000000", "json": "{\"testField\":\"bad.date\"", "expectedError": fmt.Errorf("writing field: time_point: parsing time \"bad.date\" as \"2006-01-02T15:04:05.999\": cannot parse \"bad.date\" as \"2006\"")},
		{"caseName": "time_point_sec", "typeName": "time_point_sec", "expectedValue": "01000000", "json": "{\"testField\":\"1970-01-01T00:00:01\"", "expectedError": nil},
		{"caseName": "time_point_sec err", "typeName": "time_point_sec", "expectedValue": "01000000g", "json": "{\"testField\":\"bad date\"", "expectedError": fmt.Errorf("writing field: time_point_sec: parsing time \"bad date\" as \"2006-01-02T15:04:05\": cannot parse \"bad date\" as \"2006\"")},
		{"caseName": "block_timestamp_type", "typeName": "block_timestamp_type", "expectedValue": "ec8a4546", "json": "{\"testField\":\"2018-09-05T12:48:54-04:00\"}", "expectedError": nil},
		{"caseName": "block_timestamp_type err", "typeName": "block_timestamp_type", "expectedValue": "ec8a4546", "json": "{\"testField\":\"this is not a date\"}", "expectedError": fmt.Errorf("writing field: block_timestamp_type: parsing time \"this is not a date\" as \"2006-01-02T15:04:05.999999-07:00\": cannot parse \"this is not a date\" as \"2006\"")},
		{"caseName": "Name", "typeName": "name", "expectedValue": "0000000000ea3055", "json": "{\"testField\":\"zswhq\"}", "expectedError": nil},
		{"caseName": "Name", "typeName": "name", "expectedValue": "", "json": "{\"testField\":\"waytolongnametomakethetestcrash\"}", "expectedError": fmt.Errorf("writing field: name: waytolongnametomakethetestcrash is to long. expected length of max 12 characters")},
		{"caseName": "bytes", "typeName": "bytes", "expectedValue": "0e746869732e69732e612e74657374", "json": "{\"testField\":\"746869732e69732e612e74657374\"}", "expectedError": nil},
		{"caseName": "bytes err", "typeName": "bytes", "expectedValue": "0e746869732e69732e612e74657374", "json": "{\"testField\":\"those are not bytes\"}", "expectedError": fmt.Errorf("writing field: bytes: encoding/hex: invalid byte: U+0074 't'")},
		{"caseName": "checksum160", "typeName": "checksum160", "expectedValue": "0000000000000000000000000000000000000000", "json": "{\"testField\":\"0000000000000000000000000000000000000000\"}", "expectedError": nil},
		{"caseName": "checksum256", "typeName": "checksum256", "expectedValue": "0000000000000000000000000000000000000000000000000000000000000000", "json": "{\"testField\":\"0000000000000000000000000000000000000000000000000000000000000000\"}", "expectedError": nil},
		{"caseName": "checksum512", "typeName": "checksum512", "expectedValue": "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", "json": "{\"testField\":\"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000\"}", "expectedError": nil},
		{"caseName": "checksum160 to long", "typeName": "checksum160", "expectedValue": "", "json": "{\"testField\":\"10000000000000000000000000000000000000000\"}", "expectedError": fmt.Errorf("writing field: checksum160: expected length of 40 got 41 for value 10000000000000000000000000000000000000000")},
		{"caseName": "checksum256 to long", "typeName": "checksum256", "expectedValue": "", "json": "{\"testField\":\"10000000000000000000000000000000000000000000000000000000000000000\"}", "expectedError": fmt.Errorf("writing field: checksum256: expected length of 64 got 65 for value 10000000000000000000000000000000000000000000000000000000000000000")},
		{"caseName": "checksum512 to long", "typeName": "checksum512", "expectedValue": "", "json": "{\"testField\":\"100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000\"}", "expectedError": fmt.Errorf("writing field: checksum512: expected length of 128 got 129 for value 100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")},
		{"caseName": "checksum160 hex err", "typeName": "checksum160", "expectedValue": "", "json": "{\"testField\":\"BADX000000000000000000000000000000000000\"}", "expectedError": fmt.Errorf("writing field: checksum160: encoding/hex: invalid byte: U+0058 'X'")},
		{"caseName": "checksum256 hex err", "typeName": "checksum256", "expectedValue": "", "json": "{\"testField\":\"BADX000000000000000000000000000000000000000000000000000000000000\"}", "expectedError": fmt.Errorf("writing field: checksum256: encoding/hex: invalid byte: U+0058 'X'")},
		{"caseName": "checksum512 hex err", "typeName": "checksum512", "expectedValue": "", "json": "{\"testField\":\"BADX0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000\"}", "expectedError": fmt.Errorf("writing field: checksum512: encoding/hex: invalid byte: U+0058 'X'")},
		{"caseName": "public_key", "typeName": "public_key", "expectedValue": "00000000000000000000000000000000000000000000000000000000000000000000", "json": "{\"testField\":\"" + ecc.PublicKeyPrefixCompat + "1111111111111111111111111111111114T1Anm\"}", "expectedError": nil},
		{"caseName": "public_key err", "typeName": "public_key", "expectedValue": "", "json": "{\"testField\":\"" + ecc.PublicKeyPrefixCompat + "1111111111111111111111114T1Anm\"}", "expectedError": fmt.Errorf("writing field: public_key: public key checksum failed, found f9246dd2 but expected 86e7b522")},
		{"caseName": "signature", "typeName": "signature", "expectedValue": "000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", "json": "{\"testField\":\"SIG_K1_111111111111111111111111111111111111111111111111111111111111111116uk5ne\"}", "expectedError": nil},
		{"caseName": "signature err", "typeName": "signature", "expectedValue": "", "json": "{\"testField\":\"SIG_K1_BADX11111111111111111111111111111111111111111111111111111111111116uk5ne\"}", "expectedError": fmt.Errorf("writing field: public_key: signature checksum failed, found 3aea1e96 but expected e72f76ff")},
		{"caseName": "symbol", "typeName": "symbol", "expectedValue": "04454f5300000000", "json": "{\"testField\":\"4,EOS\"}", "expectedError": nil},
		{"caseName": "symbol format error", "typeName": "symbol", "expectedValue": "", "json": "{\"testField\":\"4EOS\"}", "expectedError": fmt.Errorf("writing field: symbol: symbol should be of format '4,EOS'")},
		{"caseName": "symbol format error", "typeName": "symbol", "expectedValue": "", "json": "{\"testField\":\"abc,EOS\"}", "expectedError": fmt.Errorf("writing field: symbol: strconv.ParseUint: parsing \"abc\": invalid syntax")},
		{"caseName": "symbol_code", "typeName": "symbol_code", "expectedValue": "ffffffffffffffff", "json": "{\"testField\":18446744073709551615}", "expectedError": nil},
		{"caseName": "asset", "typeName": "asset", "expectedValue": "a08601000000000004454f5300000000", "json": "{\"testField\":\"10.0000 EOS\"}", "expectedError": nil},
		{"caseName": "asset err", "typeName": "asset", "expectedValue": "", "json": "{\"testField\":\"AA.0000 EOS\"}", "expectedError": fmt.Errorf("writing field: asset: strconv.ParseInt: parsing \"AA0000\": invalid syntax")},
		{"caseName": "extended_asset", "typeName": "extended_asset", "expectedValue": "0a0000000000000004454f5300000000202932c94c833055", "json": "{\"testField\":{\"quantity\":\"0.0010 EOS\",\"contract\":\"zhongshuwen1\"}}", "expectedError": nil},
		{"caseName": "extended_asset err", "typeName": "extended_asset", "expectedValue": "", "json": "{\"testField\":{\"quantity\":\"abc.0010 EOS\",\"contract\":\"zhongshuwen1\"}}", "expectedError": fmt.Errorf("writing field: extended_asset: strconv.ParseInt: parsing \"abc0010\": invalid syntax")},
		{"caseName": "bad type", "typeName": "bad.type.1", "expectedValue": nil, "json": "{\"testField\":0}", "expectedError": fmt.Errorf("writing field of type [bad.type.1]: unknown type")},
		{"caseName": "optional present", "typeName": "string", "expectedValue": "0776616c75652e31", "json": "{\"testField\":\"value.1\"}", "expectedError": nil},
		{"caseName": "struct", "typeName": "struct_name_1", "expectedValue": "0e746869732e69732e612e74657374", "json": "{\"testField\": {\"field_name_1\":\"this.is.a.test\"}}", "expectedError": nil},
		{"caseName": "struct err", "typeName": "struct_name_1", "expectedValue": "0e746869732e69732e612e74657374", "json": "{\"testField\": {}", "expectedError": fmt.Errorf("encoding fields: encode field: none optional field [field_name_1] as a nil value")},
	}

	for _, c := range testCases {

		t.Run(c["caseName"].(string), func(t *testing.T) {
			var buffer bytes.Buffer
			encoder := NewEncoder(&buffer)

			abi := ABI{
				Structs: []StructDef{
					{
						Name:   "struct_name_1",
						Base:   "",
						Fields: []FieldDef{{Name: "field_name_1", Type: "string"}},
					},
				},
			}
			fieldName := "test_field_name"
			result := gjson.Get(c["json"].(string), "testField")
			err := abi.writeField(encoder, fieldName, c["typeName"].(string), result)

			if c["expectedError"] == nil {
				require.Nil(t, err, c["caseName"])
				assert.Equal(t, c["expectedValue"], hex.EncodeToString(buffer.Bytes()), c["caseName"])
			} else {
				expectedError := c["expectedError"].(error)
				require.Equal(t, expectedError.Error(), err.Error(), c["caseName"])
			}
		})
	}
}

type mockWriterable interface {
	Write(p []byte) (n int, err error)
	Bytes() []byte
}
type mockWriter struct {
	length int
	err    error
}

func (w mockWriter) Write(p []byte) (n int, err error) {
	return w.length, w.err
}

func (w mockWriter) Bytes() []byte {
	return []byte{}
}
