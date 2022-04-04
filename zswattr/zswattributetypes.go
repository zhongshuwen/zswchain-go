package zswattr

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/gjson"
	zsw "github.com/zhongshuwen/zswchain-go"
	"github.com/zhongshuwen/zswchain-go/ecc"
)

var symbolRegex = regexp.MustCompile("^[0-9]{1,2},[A-Z]{1,7}$")
var symbolCodeRegex = regexp.MustCompile("^[A-Z]{1,7}$")

// For reference:

type Name string
type AccountName Name
type PermissionName Name
type ActionName Name
type TableName Name
type ScopeName Name

func AN(in string) AccountName    { return AccountName(in) }
func ActN(in string) ActionName   { return ActionName(in) }
func PN(in string) PermissionName { return PermissionName(in) }

func (n Name) String() string           { return string(n) }
func (n AccountName) String() string    { return string(n) }
func (n PermissionName) String() string { return string(n) }
func (n ActionName) String() string     { return string(n) }
func (n TableName) String() string      { return string(n) }
func (n ScopeName) String() string      { return string(n) }

// start core permissions flags
type ZswCorePermissions uint64

const (
	ZSW_CORE_PERMS_ADMIN ZswCorePermissions = 1 << iota
	ZSW_CORE_PERMS_SETCODE
	ZSW_CORE_PERMS_SETABI
	ZSW_CORE_PERMS_TRANSFER_TOKEN_TO_ANYONE
	ZSW_CORE_PERMS_RECEIVE_TOKEN_FROM_ANYONE
	ZSW_CORE_PERMS_TRANSFER_TOKEN_TO_CORE_CONTRACTS
	ZSW_CORE_PERMS_RECEIVE_TOKEN_AS_CORE_CONTRACTS
	ZSW_CORE_PERMS_CREATE_USER
	ZSW_CORE_PERMS_UPDATE_AUTH
	ZSW_CORE_PERMS_DELETE_AUTH
	ZSW_CORE_PERMS_LINK_AUTH
	ZSW_CORE_PERMS_UNLINK_AUTH
	ZSW_CORE_PERMS_CANCEL_DELAY
	ZSW_CORE_PERMS_CONFIRM_AUTHORIZE_USER_TX
	ZSW_CORE_PERMS_CONFIRM_AUTHORIZE_USER_TRANSFER_ITEM
	ZSW_CORE_PERMS_GENERAL_RESOURCES
	ZSW_CORE_PERMS_REGISTER_PRODUCER
	ZSW_CORE_PERMS_VOTE_PRODUCER
	ZSW_CORE_PERMS_MISC_FUNCTIONS
)

// end core permission flags

// start items permissions flags
type ZswItemsPermissions uint64

const (
	ZSW_ITEMS_PERMS_AUTHORIZE_CREATE_COLLECTION ZswItemsPermissions = 1 << iota
	ZSW_ITEMS_PERMS_AUTHORIZE_MODIFY_COLLECTION
	ZSW_ITEMS_PERMS_AUTHORIZE_CREATE_ITEM
	ZSW_ITEMS_PERMS_AUTHORIZE_MODIFY_ITEM
	ZSW_ITEMS_PERMS_AUTHORIZE_MINT_ITEM
	ZSW_ITEMS_PERMS_AUTHORIZE_CREATE_ISSUER
	ZSW_ITEMS_PERMS_AUTHORIZE_MODIFY_ISSUER
	ZSW_ITEMS_PERMS_AUTHORIZE_CREATE_ROYALTY_USER
	ZSW_ITEMS_PERMS_AUTHORIZE_MODIFY_ROYALTY_USER
	ZSW_ITEMS_PERMS_AUTHORIZE_CREATE_SCHEMA
	ZSW_ITEMS_PERMS_AUTHORIZE_MODIFY_SCHEMA
	ZSW_ITEMS_PERMS_ADMIN
	ZSW_ITEMS_PERMS_AUTHORIZE_CREATE_CUSTODIAN
	ZSW_ITEMS_PERMS_AUTHORIZE_MODIFY_CUSTODIAN
	ZSW_ITEMS_PERMS_AUTHORIZE_MINT_TO_OTHER_CUSTODIANS
	ZSW_ITEMS_PERMS_AUTHORIZE_MINT_TO_NULL_CUSTODIAN
	ZSW_ITEMS_PERMS_AUTHORIZE_CREATE_ITEM_TEMPLATE
	ZSW_ITEMS_PERMS_AUTHORIZE_MODIFY_ITEM_TEMPLATE
	ZSW_ITEMS_PERMS_AUTHORIZE_MODIFY_ITEM_METADATA
)

// end items permission flags

// start items custodian permissions flags
type ZswItemsCustodianPermissions uint64

const (
	CUSTODIAN_PERMS_ENABLED ZswItemsCustodianPermissions = 1 << iota
	CUSTODIAN_PERMS_TX_TO_SELF_CUSTODIAN
	CUSTODIAN_PERMS_RECEIVE_FROM_NULL_CUSTODIAN
	CUSTODIAN_PERMS_RECEIVE_FROM_ANY_CUSTODIAN
	CUSTODIAN_PERMS_RECEIVE_FROM_ZSW_CUSTODIAN
	CUSTODIAN_PERMS_SEND_TO_NULL_CUSTODIAN
	CUSTODIAN_PERMS_SEND_TO_ANY_CUSTODIAN
	CUSTODIAN_PERMS_SEND_TO_ZSW_CUSTODIAN
)

// end items custodian permission flags
// start items config
type ZswItemConfigFlags uint64

const (
	ITEM_CONFIG_TRANSFERABLE ZswItemConfigFlags = 1 << iota
	ITEM_CONFIG_BURNABLE
	ITEM_CONFIG_FROZEN
	ITEM_CONFIG_ALLOW_NOTIFY
	ITEM_CONFIG_ALLOW_MUTABLE_DATA
)

//end items config
type SafeString string

func (ss *SafeString) UnmarshalBinary(d *zsw.Decoder) error {
	s, e := d.SafeReadUTF8String()
	if e != nil {
		return e
	}

	*ss = SafeString(s)
	return nil
}

type AccountResourceLimit struct {
	Used      Int64 `json:"used"`
	Available Int64 `json:"available"`
	Max       Int64 `json:"max"`
}

type DelegatedBandwidth struct {
	From      AccountName `json:"from"`
	To        AccountName `json:"to"`
	NetWeight Asset       `json:"net_weight"`
	CPUWeight Asset       `json:"cpu_weight"`
}

type TotalResources struct {
	Owner     AccountName `json:"owner"`
	NetWeight Asset       `json:"net_weight"`
	CPUWeight Asset       `json:"cpu_weight"`
	RAMBytes  Int64       `json:"ram_bytes"`
}

type VoterInfo struct {
	Owner             AccountName   `json:"owner"`
	Proxy             AccountName   `json:"proxy"`
	Producers         []AccountName `json:"producers"`
	Staked            Int64         `json:"staked"`
	LastVoteWeight    Float64       `json:"last_vote_weight"`
	ProxiedVoteWeight Float64       `json:"proxied_vote_weight"`
	IsProxy           byte          `json:"is_proxy"`
}

type RefundRequest struct {
	Owner       AccountName `json:"owner"`
	RequestTime JSONTime    `json:"request_time"` //         {"name":"request_time", "type":"time_point_sec"},
	NetAmount   Asset       `json:"net_amount"`
	CPUAmount   Asset       `json:"cpu_amount"`
}

type CompressionType uint8

const (
	CompressionNone = CompressionType(iota)
	CompressionZlib
)

func (c CompressionType) String() string {
	switch c {
	case CompressionNone:
		return "none"
	case CompressionZlib:
		return "zlib"
	default:
		return ""
	}
}

func (c CompressionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *CompressionType) UnmarshalJSON(data []byte) error {
	tryNext, err := c.tryUnmarshalJSONAsBool(data)
	if err != nil && !tryNext {
		return err
	}
	if !tryNext {
		return nil
	}

	tryNext, err = c.tryUnmarshalJSONAsString(data)
	if err != nil && !tryNext {
		return err
	}
	if !tryNext {
		return nil
	}

	_, err = c.tryUnmarshalJSONAsUint8(data)
	return err
}

func (c *CompressionType) tryUnmarshalJSONAsString(data []byte) (tryNext bool, err error) {
	var s string
	err = json.Unmarshal(data, &s)
	if err != nil {
		_, isTypeError := err.(*json.UnmarshalTypeError)

		// Let's continue with next handler is we hit a type error, might be an integer...
		return isTypeError, err
	}

	switch s {
	case "none":
		*c = CompressionNone
	case "zlib":
		*c = CompressionZlib
	default:
		return false, fmt.Errorf("unknown compression type %s", s)
	}

	return false, nil
}

func (c *CompressionType) tryUnmarshalJSONAsBool(data []byte) (tryNext bool, err error) {
	var b bool
	err = json.Unmarshal(data, &b)
	if err != nil {
		_, isTypeError := err.(*json.UnmarshalTypeError)

		// Let's continue with next handler is we hit a type error, might be an integer...
		return isTypeError, err
	}

	if b {
		*c = CompressionZlib
	} else {
		*c = CompressionNone
	}
	return false, nil
}

func (c *CompressionType) tryUnmarshalJSONAsUint8(data []byte) (tryNext bool, err error) {
	var s uint8
	err = json.Unmarshal(data, &s)
	if err != nil {
		return false, err
	}

	switch s {
	case 0:
		*c = CompressionNone
	case 1:
		*c = CompressionZlib
	default:
		return false, fmt.Errorf("unknown compression type %d", s)
	}

	return false, nil
}

// CurrencyName

type CurrencyName string

type Bool bool

func (b *Bool) UnmarshalJSON(data []byte) error {
	var num int
	err := json.Unmarshal(data, &num)
	if err == nil {
		*b = Bool(num != 0)
		return nil
	}

	var boolVal bool
	if err := json.Unmarshal(data, &boolVal); err != nil {
		return fmt.Errorf("couldn't unmarshal bool as int or true/false: %w", err)
	}

	*b = Bool(boolVal)
	return nil
}

// Asset

// NOTE: there's also ExtendedAsset which is a quantity with the attached contract (AccountName)
type Asset struct {
	Amount Int64
	Symbol
}

func (a Asset) Add(other Asset) Asset {
	if a.Symbol != other.Symbol {
		panic("Add applies only to assets with the same symbol")
	}
	return Asset{Amount: a.Amount + other.Amount, Symbol: a.Symbol}
}

func (a Asset) Sub(other Asset) Asset {
	if a.Symbol != other.Symbol {
		panic("Sub applies only to assets with the same symbol")
	}
	return Asset{Amount: a.Amount - other.Amount, Symbol: a.Symbol}
}

func (a Asset) String() string {
	amt := a.Amount
	if amt < 0 {
		amt = -amt
	}

	precisionDigitCount := int(a.Symbol.Precision)
	dotAndPrecisionDigitCount := precisionDigitCount + 1

	strInt := strconv.FormatInt(int64(amt), 10)
	if len(strInt) < dotAndPrecisionDigitCount {
		// prepend `0` for the difference:
		strInt = strings.Repeat("0", dotAndPrecisionDigitCount-len(strInt)) + strInt
	}

	result := strInt
	if a.Symbol.Precision > 0 {
		result = strInt[:len(strInt)-precisionDigitCount] + "." + strInt[len(strInt)-precisionDigitCount:]
	}

	if a.Amount < 0 {
		result = "-" + result
	}

	return fmt.Sprintf("%s %s", result, a.Symbol.Symbol)
}

type ExtendedAsset struct {
	Asset    Asset       `json:"quantity"`
	Contract AccountName `json:"contract"`
}

// NOTE: there's also a new ExtendedSymbol (which includes the contract (as AccountName) on which it is)
type Symbol struct {
	Precision uint8
	Symbol    string

	// Caching of symbol code if it was computed once
	symbolCode uint64
}

func NewSymbolFromUint64(value uint64) (out Symbol) {
	out.Precision = uint8(value & 0xFF)
	out.symbolCode = value >> 8
	out.Symbol = SymbolCode(out.symbolCode).String()

	return
}

func NameToSymbol(name Name) (Symbol, error) {
	symbol := Symbol{}
	value, err := zsw.StringToName(string(name))
	if err != nil {
		return symbol, fmt.Errorf("name %s is invalid: %w", name, err)
	}

	symbol.Precision = uint8(value & 0xFF)
	symbol.symbolCode = value >> 8
	symbol.Symbol = SymbolCode(symbol.symbolCode).String()

	return symbol, nil
}

func StringToSymbol(str string) (Symbol, error) {
	symbol := Symbol{}
	if !symbolRegex.MatchString(str) {
		return symbol, fmt.Errorf("%s is not a valid symbol", str)
	}
	arrs := strings.Split(str, ",")
	precision, _ := strconv.ParseUint(string(arrs[0]), 10, 8)

	symbol.Precision = uint8(precision)
	symbol.Symbol = arrs[1]

	return symbol, nil
}

func MustStringToSymbol(str string) Symbol {
	symbol, err := StringToSymbol(str)
	if err != nil {
		panic(fmt.Errorf("invalid symbol %q: %w", str, err))
	}

	return symbol
}

func (s Symbol) SymbolCode() (SymbolCode, error) {
	if s.symbolCode != 0 {
		return SymbolCode(s.symbolCode), nil
	}

	symbolCode, err := StringToSymbolCode(s.Symbol)
	if err != nil {
		return 0, err
	}

	return SymbolCode(symbolCode), nil
}

func (s Symbol) MustSymbolCode() SymbolCode {
	symbolCode, err := StringToSymbolCode(s.Symbol)
	if err != nil {
		panic("invalid symbol code " + s.Symbol)
	}

	return symbolCode
}

func (s Symbol) ToUint64() (uint64, error) {
	symbolCode, err := s.SymbolCode()
	if err != nil {
		return 0, fmt.Errorf("symbol %s is not a valid symbol code: %w", s.Symbol, err)
	}

	return uint64(symbolCode)<<8 | uint64(s.Precision), nil
}

func (s Symbol) ToName() (string, error) {
	u, err := s.ToUint64()
	if err != nil {
		return "", err
	}
	return zsw.NameToString(u), nil
}

func (s Symbol) String() string {
	return fmt.Sprintf("%d,%s", s.Precision, s.Symbol)
}

type SymbolCode uint64

func NameToSymbolCode(name Name) (SymbolCode, error) {
	value, err := zsw.StringToName(string(name))
	if err != nil {
		return 0, fmt.Errorf("name %s is invalid: %w", name, err)
	}

	return SymbolCode(value), nil
}

func StringToSymbolCode(str string) (SymbolCode, error) {
	if len(str) > 7 {
		return 0, fmt.Errorf("string is too long to be a valid symbol_code")
	}

	var symbolCode uint64
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] < 'A' || str[i] > 'Z' {
			return 0, fmt.Errorf("only uppercase letters allowed in symbol_code string")
		}

		symbolCode <<= 8
		symbolCode = symbolCode | uint64(str[i])
	}

	return SymbolCode(symbolCode), nil
}

func (sc SymbolCode) ToName() string {
	return zsw.NameToString(uint64(sc))
}

func (sc SymbolCode) String() string {
	builder := strings.Builder{}

	symbolCode := uint64(sc)
	for i := 0; i < 7; i++ {
		if symbolCode == 0 {
			return builder.String()
		}

		builder.WriteByte(byte(symbolCode & 0xFF))
		symbolCode >>= 8
	}

	return builder.String()
}

func (sc SymbolCode) MarshalJSON() (data []byte, err error) {
	return []byte(`"` + sc.String() + `"`), nil
}

// ZSWCCSymbol represents the standard ZSWCC symbol on the chain.  It's
// here just to speed up things.
var ZSWCCSymbol = Symbol{Precision: 4, Symbol: "ZSWCC"}

// REXSymbol represents the standard REX symbol on the chain.  It's
// here just to speed up things.
var REXSymbol = Symbol{Precision: 4, Symbol: "REX"}

// TNTSymbol represents the standard EOSIO Testnet symbol on the testnet chain.
// Temporary Network Token (TNT) is the native token of the EOSIO Testnet.
// It's here just to speed up things.
var TNTSymbol = Symbol{Precision: 4, Symbol: "TNT"}

func NewZSWAsset(amount int64) Asset {
	return Asset{Amount: Int64(amount), Symbol: ZSWCCSymbol}
}

// NewAsset reads from a string an EOS asset.
//
// Deprecated: Use `NewAssetFromString` instead
func NewAsset(in string) (out Asset, err error) {
	return NewAssetFromString(in)
}

// NewAssetFromString reads a string an decode it to an zsw.Asset
// structure if possible. The input must contains an amount and
// a symbol. The precision is inferred based on the actual number
// of decimals present.
func NewAssetFromString(in string) (out Asset, err error) {
	out, err = newAssetFromString(in)
	if err != nil {
		return out, err
	}

	if out.Symbol.Symbol == "" {
		return out, fmt.Errorf("invalid format %q, expected an amount and a currency symbol", in)
	}

	return
}

func NewZSWAssetFromString(input string) (Asset, error) {
	return NewFixedSymbolAssetFromString(ZSWCCSymbol, input)
}

func NewREXAssetFromString(input string) (Asset, error) {
	return NewFixedSymbolAssetFromString(REXSymbol, input)
}

func NewTNTAssetFromString(input string) (Asset, error) {
	return NewFixedSymbolAssetFromString(TNTSymbol, input)
}

func NewFixedSymbolAssetFromString(symbol Symbol, input string) (out Asset, err error) {
	integralPart, decimalPart, symbolPart, err := splitAsset(input)
	if err != nil {
		return out, err
	}

	symbolCode := symbol.MustSymbolCode().String()
	precision := symbol.Precision

	if len(decimalPart) > int(precision) {
		return out, fmt.Errorf("symbol %s precision mismatch: expected %d, got %d", symbol, precision, len(decimalPart))
	}

	if symbolPart != "" && symbolPart != symbolCode {
		return out, fmt.Errorf("symbol %s code mismatch: expected %s, got %s", symbol, symbolCode, symbolPart)
	}

	if len(decimalPart) < int(precision) {
		decimalPart += strings.Repeat("0", int(precision)-len(decimalPart))
	}

	val, err := strconv.ParseInt(integralPart+decimalPart, 10, 64)
	if err != nil {
		return out, err
	}

	return Asset{
		Amount: Int64(val),
		Symbol: Symbol{Precision: precision, Symbol: symbolCode},
	}, nil
}

func newAssetFromString(in string) (out Asset, err error) {
	integralPart, decimalPart, symbolPart, err := splitAsset(in)
	if err != nil {
		return out, err
	}

	val, err := strconv.ParseInt(integralPart+decimalPart, 10, 64)
	if err != nil {
		return out, err
	}

	out.Amount = Int64(val)
	out.Symbol.Precision = uint8(len(decimalPart))
	out.Symbol.Symbol = symbolPart

	return
}

func splitAsset(input string) (integralPart, decimalPart, symbolPart string, err error) {
	input = strings.Trim(input, " ")
	if len(input) == 0 {
		return "", "", "", fmt.Errorf("input cannot be empty")
	}

	parts := strings.Split(input, " ")
	if len(parts) >= 1 {
		integralPart, decimalPart, err = splitAssetAmount(parts[0])
		if err != nil {
			return
		}
	}

	if len(parts) == 2 {
		symbolPart = parts[1]
		if len(symbolPart) > 7 {
			return "", "", "", fmt.Errorf("invalid asset %q, symbol should have less than 7 characters", input)
		}
	}

	if len(parts) > 2 {
		return "", "", "", fmt.Errorf("invalid asset %q, expecting an amount alone or an amount and a currency symbol", input)
	}

	return
}

func splitAssetAmount(input string) (integralPart, decimalPart string, err error) {
	parts := strings.Split(input, ".")
	switch len(parts) {
	case 1:
		integralPart = parts[0]
	case 2:
		integralPart = parts[0]
		decimalPart = parts[1]

		if len(decimalPart) > math.MaxUint8 {
			err = fmt.Errorf("invalid asset amount precision %q, should have less than %d characters", input, math.MaxUint8)

		}
	default:
		return "", "", fmt.Errorf("invalid asset amount %q, expected amount to have at most a single dot", input)
	}

	return
}

func (a *Asset) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	asset, err := NewAsset(s)
	if err != nil {
		return err
	}

	*a = asset

	return nil
}

func (a Asset) MarshalJSON() (data []byte, err error) {
	return json.Marshal(a.String())
}

type Permission struct {
	PermName     string    `json:"perm_name"`
	Parent       string    `json:"parent"`
	RequiredAuth Authority `json:"required_auth"`
}

type PermissionLevel struct {
	Actor      AccountName    `json:"actor"`
	Permission PermissionName `json:"permission"`
}

// NewPermissionLevel parses strings like `account@active`,
// `otheraccount@owner` and builds a PermissionLevel struct. It
// validates that there is a single optional @ (where permission
// defaults to 'active'), and validates length of account and
// permission names.
func NewPermissionLevel(in string) (out PermissionLevel, err error) {
	parts := strings.Split(in, "@")
	if len(parts) > 2 {
		return out, fmt.Errorf("permission %q invalid, use account[@permission]", in)
	}

	if len(parts[0]) > 12 {
		return out, fmt.Errorf("account name %q too long", parts[0])
	}

	out.Actor = AccountName(parts[0])
	out.Permission = PermissionName("active")
	if len(parts) == 2 {
		if len(parts[1]) > 12 {
			return out, fmt.Errorf("permission %q name too long", parts[1])
		}

		out.Permission = PermissionName(parts[1])
	}

	return
}

type PermissionLevelWeight struct {
	Permission PermissionLevel `json:"permission"`
	Weight     uint16          `json:"weight"` // weight_type
}

type Authority struct {
	Threshold uint32                  `json:"threshold"`
	Keys      []KeyWeight             `json:"keys,omitempty"`
	Accounts  []PermissionLevelWeight `json:"accounts,omitempty"`
	Waits     []WaitWeight            `json:"waits,omitempty"`
}

type KeyWeight struct {
	PublicKey ecc.PublicKey `json:"key"`
	Weight    uint16        `json:"weight"` // weight_type
}

type WaitWeight struct {
	WaitSec uint32 `json:"wait_sec"`
	Weight  uint16 `json:"weight"` // weight_type
}

type GetRawCodeAndABIResp struct {
	AccountName  AccountName `json:"account_name"`
	WASMasBase64 string      `json:"wasm"`
	ABIasBase64  string      `json:"abi"`
}

type GetCodeResp struct {
	AccountName AccountName `json:"account_name"`
	CodeHash    string      `json:"code_hash"`
	WASM        string      `json:"wasm"`
	ABI         zsw.ABI     `json:"abi"`
}

type GetCodeHashResp struct {
	AccountName AccountName `json:"account_name"`
	CodeHash    string      `json:"code_hash"`
}

type GetABIResp struct {
	AccountName AccountName `json:"account_name"`
	ABI         zsw.ABI     `json:"abi"`
}

type ABIJSONToBinResp struct {
	Binargs string `json:"binargs"`
}

type ABIBinToJSONResp struct {
	Args zsw.M `json:"args"`
}

// JSONTime

type JSONTime struct {
	time.Time
}

const JSONTimeFormat = "2006-01-02T15:04:05"

func (t JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", t.Format(JSONTimeFormat))), nil
}

func (t *JSONTime) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" {
		return nil
	}

	t.Time, err = time.Parse(`"`+JSONTimeFormat+`"`, string(data))
	return err
}

// ParseJSONTime will parse a string into a JSONTime object
func ParseJSONTime(date string) (JSONTime, error) {
	var t JSONTime
	var err error
	t.Time, err = time.Parse(JSONTimeFormat, string(date))
	return t, err
}

// HexBytes

type HexBytes []byte

func (t HexBytes) MarshalJSON() ([]byte, error) {
	return json.Marshal(hex.EncodeToString(t))
}

func (t *HexBytes) UnmarshalJSON(data []byte) (err error) {
	var s string
	err = json.Unmarshal(data, &s)
	if err != nil {
		return
	}

	*t, err = hex.DecodeString(s)
	return
}

func (t HexBytes) String() string {
	return hex.EncodeToString(t)
}

// Checksum256

type Checksum160 []byte

func (t Checksum160) MarshalJSON() ([]byte, error) {
	return json.Marshal(hex.EncodeToString(t))
}
func (t *Checksum160) UnmarshalJSON(data []byte) (err error) {
	var s string
	err = json.Unmarshal(data, &s)
	if err != nil {
		return
	}

	*t, err = hex.DecodeString(s)
	return
}

type Checksum256 []byte

func (t Checksum256) MarshalJSON() ([]byte, error) {
	return json.Marshal(hex.EncodeToString(t))
}
func (t *Checksum256) UnmarshalJSON(data []byte) (err error) {
	var s string
	err = json.Unmarshal(data, &s)
	if err != nil {
		return
	}

	*t, err = hex.DecodeString(s)
	return
}

func (t Checksum256) String() string {
	return hex.EncodeToString(t)
}

type Checksum512 []byte

func (t Checksum512) MarshalJSON() ([]byte, error) {
	return json.Marshal(hex.EncodeToString(t))
}
func (t *Checksum512) UnmarshalJSON(data []byte) (err error) {
	var s string
	err = json.Unmarshal(data, &s)
	if err != nil {
		return
	}

	*t, err = hex.DecodeString(s)
	return
}

// SHA256Bytes is deprecated and renamed to Checksum256 for
// consistency. Please update your code as this type will eventually
// be phased out.
type SHA256Bytes = Checksum256

type Varuint32 uint32
type Varint32 int32

// Tstamp

type Tstamp struct {
	time.Time
}

func (t Tstamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%d", t.UnixNano()))
}

func (t *Tstamp) UnmarshalJSON(data []byte) (err error) {
	var unixNano int64
	if data[0] == '"' {
		var s string
		if err = json.Unmarshal(data, &s); err != nil {
			return
		}

		unixNano, err = strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}

	} else {
		unixNano, err = strconv.ParseInt(string(data), 10, 64)
		if err != nil {
			return err
		}
	}

	*t = Tstamp{time.Unix(0, unixNano)}

	return nil
}

// BlockNum extracts the block number (or height) from a hex-encoded block ID.
func BlockNum(blockID string) uint32 {
	if len(blockID) < 8 {
		return 0
	}
	bin, err := hex.DecodeString(blockID[:8])
	if err != nil {
		return 0
	}
	return binary.BigEndian.Uint32(bin)
}

type BlockTimestamp struct {
	time.Time
}

// blockTimestampFormat
//
// We deal with timezone in a conditional matter so we allowed for example the
// unmarshalling to accept with and without timezone specifier.
const blockTimestampFormat = "2006-01-02T15:04:05.999"

func (t BlockTimestamp) MarshalJSON() ([]byte, error) {
	strTime := t.Format(blockTimestampFormat)
	if len(strTime) == len("2006-01-02T15:04:05.5") {
		strTime += "00"
	}

	return []byte(`"` + strTime + `"`), nil
}

func (t *BlockTimestamp) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" {
		return nil
	}

	t.Time, err = time.Parse(`"`+blockTimestampFormat+`"`, string(data))
	if err != nil {
		t.Time, err = time.Parse(`"`+blockTimestampFormat+`Z07:00"`, string(data))
	}

	return err
}

// TimePoint represents the number of microseconds since EPOCH (Jan 1st 1970)
type TimePoint uint64

const nodeosTimePointFormat = "2006-01-02T15:04:05.000"

func formatTimePoint(timePoint TimePoint, shouldFitNodeos bool) string {
	t := time.Unix(0, int64(timePoint*1000))
	if shouldFitNodeos {
		return t.UTC().Format(nodeosTimePointFormat)
	}

	return t.UTC().Format(standardTimePointFormat)
}

func (f TimePoint) String() string {
	return formatTimePoint(f, true)
}

func (f TimePoint) MarshalJSON() ([]byte, error) {
	return []byte(`"` + formatTimePoint(f, true) + `"`), nil
}

func (f *TimePoint) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	out, err := time.Parse(standardTimePointFormat, s)
	if err != nil {
		return err
	}

	*f = TimePoint(out.UnixNano() / 1000)
	return nil
}

const standardTimePointSecFormat = "2006-01-02T15:04:05"

func formatTimePointSec(timePoint TimePointSec) string {
	t := time.Unix(int64(timePoint), 0)

	return t.UTC().Format(standardTimePointSecFormat)
}

func formatFloat(v float64, fitNodeos bool) interface{} {
	switch {
	case math.IsInf(v, 1):
		return "inf"
	case math.IsInf(v, -1):
		return "-inf"
	case math.IsNaN(v): // cannot check equality on math.NaN()
		return "nan"
	default:
	}

	if fitNodeos {
		return strconv.FormatFloat(v, 'f', 17, 64)
	} else {
		return json.RawMessage(strconv.FormatFloat(float64(v), 'f', -1, 64))
	}

}

const standardTimePointFormat = "2006-01-02T15:04:05.999"

// TimePointSec represents the number of seconds since EPOCH (Jan 1st 1970)
type TimePointSec uint32

func (f TimePointSec) String() string {
	return formatTimePointSec(f)
}

func (f TimePointSec) MarshalJSON() ([]byte, error) {
	return []byte(`"` + formatTimePointSec(f) + `"`), nil
}

func (f *TimePointSec) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	out, err := time.Parse(standardTimePointFormat, s)
	if err != nil {
		return err
	}

	*f = TimePointSec(out.Unix())
	return nil
}

type JSONFloat64 = Float64

type Float64 float64

func (f *Float64) MarshalJSON() ([]byte, error) {
	switch {
	case math.IsInf(float64(*f), 1):
		return []byte("\"inf\""), nil
	case math.IsInf(float64(*f), -1):
		return []byte("\"-inf\""), nil
	case math.IsNaN(float64(*f)):
		return []byte("\"nan\""), nil
	default:
	}
	return json.Marshal(float64(*f))
}

func (f *Float64) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("empty value")
	}

	if data[0] == '"' {
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}

		switch s {
		case "inf":
			*f = Float64(math.Inf(1))
			return nil
		case "-inf":
			*f = Float64(math.Inf(-1))
			return nil
		case "nan":
			*f = Float64(math.NaN())
			return nil
		default:
		}

		val, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}

		*f = Float64(val)

		return nil
	}

	var fl float64
	if err := json.Unmarshal(data, &fl); err != nil {
		return err
	}

	*f = Float64(fl)

	return nil
}

// JSONInt64 is deprecated in favor of Int64.
type JSONInt64 = Int64

type Int64 int64

func (i Int64) MarshalJSON() (data []byte, err error) {
	if i > 0xffffffff || i < -0xffffffff {
		encodedInt, err := json.Marshal(int64(i))
		if err != nil {
			return nil, err
		}
		data = append([]byte{'"'}, encodedInt...)
		data = append(data, '"')
		return data, nil
	}
	return json.Marshal(int64(i))
}

func (i *Int64) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("empty value")
	}

	if data[0] == '"' {
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}

		val, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}

		*i = Int64(val)

		return nil
	}

	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = Int64(v)

	return nil
}

type Uint64 uint64

func (i Uint64) MarshalJSON() (data []byte, err error) {
	if i > 0xffffffff {
		encodedInt, err := json.Marshal(uint64(i))
		if err != nil {
			return nil, err
		}
		data = append([]byte{'"'}, encodedInt...)
		data = append(data, '"')
		return data, nil
	}
	return json.Marshal(uint64(i))
}

func (i *Uint64) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("empty value")
	}

	if data[0] == '"' {
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}

		val, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return err
		}

		*i = Uint64(val)

		return nil
	}

	var v uint64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = Uint64(v)

	return nil
}

/*
func (i Uint64) MarshalBinary(encoder *zsw.Encoder) error {
	return encoder.writeUint64N(uint64(i))
}
*/

// uint128
type Uint128 struct {
	Lo uint64
	Hi uint64
}

func (i Uint128) BigInt() *big.Int {
	buf := make([]byte, 16)
	binary.BigEndian.PutUint64(buf[:], i.Hi)
	binary.BigEndian.PutUint64(buf[8:], i.Lo)
	value := (&big.Int{}).SetBytes(buf)
	return value
}

func (i Uint128) String() string {
	//Same for Int128, Float128
	number := make([]byte, 16)
	binary.LittleEndian.PutUint64(number[:], i.Lo)
	binary.LittleEndian.PutUint64(number[8:], i.Hi)
	return fmt.Sprintf("0x%s%s", hex.EncodeToString(number[:8]), hex.EncodeToString(number[8:]))
}

func (i Uint128) DecimalString() string {
	return i.BigInt().String()
}
func (i Uint128) GetTypeACode() uint64 {
	return i.Lo
}
func (i Uint128) GetTypeBCode() uint64 {
	return i.Hi
}
func (i Uint128) Get40BitId() uint64 {
	return i.Lo & 0xffffffffff
}

func (i Uint128) MarshalJSON() (data []byte, err error) {
	return []byte(`"` + i.String() + `"`), nil
}

func (i *Uint128) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	if !strings.HasPrefix(s, "0x") && !strings.HasPrefix(s, "0X") {
		return fmt.Errorf("int128 expects 0x prefix")
	}

	truncatedVal := s[2:]
	if len(truncatedVal) != 32 {
		return fmt.Errorf("int128 expects 32 characters after 0x, had %d", len(truncatedVal))
	}

	loHex := truncatedVal[:16]
	hiHex := truncatedVal[16:]

	lo, err := hex.DecodeString(loHex)
	if err != nil {
		return err
	}

	hi, err := hex.DecodeString(hiHex)
	if err != nil {
		return err
	}

	loUint := binary.LittleEndian.Uint64(lo)
	hiUint := binary.LittleEndian.Uint64(hi)

	i.Lo = loUint
	i.Hi = hiUint

	return nil
}

func (i *Uint128) FromHexString(hexString string) error {
	if len(hexString) != 32 {
		return fmt.Errorf("Uint128 hex string must be 32 characters long")
	}

	loHex := hexString[:16]
	hiHex := hexString[16:]

	lo, err := hex.DecodeString(loHex)
	if err != nil {
		return err
	}

	hi, err := hex.DecodeString(hiHex)
	if err != nil {
		return err
	}

	loUint := binary.LittleEndian.Uint64(lo)
	hiUint := binary.LittleEndian.Uint64(hi)

	i.Lo = loUint
	i.Hi = hiUint

	return nil
}

func (i *Uint128) FromHexStringBigEndian(hexString string) error {
	if len(hexString) != 32 {
		return fmt.Errorf("Uint128 hex string must be 32 characters long")
	}

	loHex := hexString[:16]
	hiHex := hexString[16:]

	lo, err := hex.DecodeString(loHex)
	if err != nil {
		return err
	}

	hi, err := hex.DecodeString(hiHex)
	if err != nil {
		return err
	}

	loUint := binary.BigEndian.Uint64(lo)
	hiUint := binary.BigEndian.Uint64(hi)

	i.Lo = hiUint
	i.Hi = loUint

	return nil
}
func (i *Uint128) FromUuidString(uuidString string) error {
	return i.FromHexStringBigEndian(strings.Replace(uuidString, "-", "", -1))
}
func NewUint128FromUint64(i uint64) Uint128 {
	return Uint128{
		Hi: 0,
		Lo: i,
	}
}

// Int128
type Int128 Uint128

func (i Int128) BigInt() *big.Int {
	comp := byte(0x80)
	buf := make([]byte, 16)
	binary.BigEndian.PutUint64(buf[:], i.Hi)
	binary.BigEndian.PutUint64(buf[8:], i.Lo)

	var value *big.Int
	if (buf[0] & comp) == comp {
		buf = twosComplement(buf)
		value = (&big.Int{}).SetBytes(buf)
		value = value.Neg(value)
	} else {
		value = (&big.Int{}).SetBytes(buf)
	}
	return value
}

func (i Int128) DecimalString() string {
	return i.BigInt().String()
}

func (i Int128) MarshalJSON() (data []byte, err error) {
	return []byte(`"` + Uint128(i).String() + `"`), nil
}

func (i *Int128) UnmarshalJSON(data []byte) error {
	var el Uint128
	if err := json.Unmarshal(data, &el); err != nil {
		return err
	}

	out := Int128(el)
	*i = out

	return nil
}

type Float128 Uint128

func (i Float128) MarshalJSON() (data []byte, err error) {
	return []byte(`"` + Uint128(i).String() + `"`), nil
}

func (i *Float128) UnmarshalJSON(data []byte) error {
	var el Uint128
	if err := json.Unmarshal(data, &el); err != nil {
		return err
	}

	out := Float128(el)
	*i = out

	return nil
}

// Blob

// Blob is base64 encoded data
// https://github.com/EOSIO/fc/blob/0e74738e938c2fe0f36c5238dbc549665ddaef82/include/fc/variant.hpp#L47
type Blob string

// Data returns decoded base64 data
func (b Blob) Data() ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(b))
}

// String returns the blob as a string
func (b Blob) String() string {
	return string(b)
}

//
/// Variant (emulates `fc::static_variant` type)
//

type Variant interface {
	Assign(typeID uint, impl interface{})
	Obtain() (typeID uint, impl interface{})
}

type VariantType struct {
	Name string
	Type interface{}
}

type VariantDefinition struct {
	typeIDToType map[uint8]reflect.Type
	typeIDToName map[uint8]string
	typeNameToID map[string]uint8
}

// NewVariantDefinition creates a variant definition based on the *ordered* provided types.
// It's the ordering that defines the binary variant value just like in native `nodeos` C++
// and in Smart Contract via the `std::variant` type. It's important to pass the entries
// in the right order!
//
// This variant definition can now be passed to functions of `BaseVariant` to implement
// marshal/unmarshaling functionalities for binary & JSON.
func NewVariantDefinition(types []VariantType) (out *VariantDefinition) {
	if len(types) < 0 {
		panic("it's not valid to create a variant definition without any types")
	}

	typeCount := len(types)
	out = &VariantDefinition{
		typeIDToType: make(map[uint8]reflect.Type, typeCount),
		typeIDToName: make(map[uint8]string, typeCount),
		typeNameToID: make(map[string]uint8, typeCount),
	}

	for i, typeDef := range types {
		typeID := uint8(i)

		// FIXME: Check how the reflect.Type is used and cache all its usage in the definition.
		//        Right now, on each Unmarshal, we re-compute some expensive stuff that can be
		//        re-used like the `typeGo.Elem()` which is always the same. It would be preferable
		//        to have those already pre-defined here so we can actually speed up the
		//        Unmarshal code.
		out.typeIDToType[typeID] = reflect.TypeOf(typeDef.Type)
		out.typeIDToName[typeID] = typeDef.Name
		out.typeNameToID[typeDef.Name] = typeID
	}

	return out
}

func (d *VariantDefinition) TypeID(name string) uint8 {
	id, found := d.typeNameToID[name]
	if !found {
		knownNames := make([]string, len(d.typeNameToID))
		i := 0
		for name := range d.typeNameToID {
			knownNames[i] = name
			i++
		}

		panic(fmt.Errorf("trying to use an unknown type name %q, known names are %q", name, strings.Join(knownNames, ", ")))
	}

	return id
}

type VariantImplFactory = func() interface{}
type OnVariant = func(impl interface{}) error

type BaseVariant struct {
	TypeID uint8
	Impl   interface{}
}

func (a *BaseVariant) Assign(typeID uint8, impl interface{}) {
	a.TypeID = typeID
	a.Impl = impl
}

func (a *BaseVariant) Obtain(def *VariantDefinition) (typeID uint32, typeName string, impl interface{}) {
	return uint32(a.TypeID), def.typeIDToName[a.TypeID], a.Impl
}

func (a *BaseVariant) MarshalJSON(def *VariantDefinition) ([]byte, error) {
	typeName, found := def.typeIDToName[a.TypeID]
	if !found {
		return nil, fmt.Errorf("type %d is not know by variant definition", a.TypeID)
	}

	return json.Marshal([]interface{}{typeName, a.Impl})
}

func (a *BaseVariant) UnmarshalJSON(data []byte, def *VariantDefinition) error {
	typeResult := gjson.GetBytes(data, "0")
	implResult := gjson.GetBytes(data, "1")

	if !typeResult.Exists() || !implResult.Exists() {
		return fmt.Errorf("invalid format, expected '[<type>, <impl>]' pair, got %q", string(data))
	}

	typeName := typeResult.String()
	typeID, found := def.typeNameToID[typeName]
	if !found {
		return fmt.Errorf("type %q is not know by variant definition", typeName)
	}

	typeGo := def.typeIDToType[typeID]
	if typeGo == nil {
		return fmt.Errorf("no known type for %q", typeName)
	}

	a.TypeID = typeID

	if typeGo.Kind() == reflect.Ptr {
		a.Impl = reflect.New(typeGo.Elem()).Interface()
		if err := json.Unmarshal([]byte(implResult.Raw), a.Impl); err != nil {
			return err
		}
	} else {
		// This is not the most optimal way of doing things for "value"
		// types (over "pointer" types) as we always allocate a new pointer
		// element, unmarshal it and then either keep the pointer type or turn
		// it into a value type.
		//
		// However, in non-reflection based code, one would do like this and
		// avoid an `new` memory allocation:
		//
		// ```
		// name := zsw.Name("")
		// json.Unmarshal(data, &name)
		// ```
		//
		// This would work without a problem. In reflection code however, I
		// did not find how one can go from `reflect.Zero(typeGo)` (which is
		// the equivalence of doing `name := zsw.Name("")`) and take the
		// pointer to it so it can be unmarshalled correctly.
		//
		// A played with various iteration, and nothing got it working. Maybe
		// the next step would be to explore the `unsafe` package and obtain
		// an unsafe pointer and play with it.
		value := reflect.New(typeGo)
		if err := json.Unmarshal([]byte(implResult.Raw), value.Interface()); err != nil {
			return err
		}

		a.Impl = value.Elem().Interface()
	}

	return nil
}

func ptr(v reflect.Value) reflect.Value {
	pt := reflect.PtrTo(v.Type())
	pv := reflect.New(pt.Elem())
	pv.Elem().Set(v)
	return pv
}

func (a *BaseVariant) UnmarshalBinaryVariant(decoder *zsw.Decoder, def *VariantDefinition) error {
	typeID, err := decoder.ReadUvarint32()
	if err != nil {
		return fmt.Errorf("unable to read variant type id: %w", err)
	}

	a.TypeID = uint8(typeID)
	typeGo := def.typeIDToType[uint8(typeID)]
	if typeGo == nil {
		return fmt.Errorf("no known type for type %d", typeID)
	}

	if typeGo.Kind() == reflect.Ptr {
		a.Impl = reflect.New(typeGo.Elem()).Interface()
		if err = decoder.Decode(a.Impl); err != nil {
			return fmt.Errorf("unable to decode variant type %d: %w", typeID, err)
		}
	} else {
		// This is not the most optimal way of doing things for "value"
		// types (over "pointer" types) as we always allocate a new pointer
		// element, unmarshal it and then either keep the pointer type or turn
		// it into a value type.
		//
		// However, in non-reflection based code, one would do like this and
		// avoid an `new` memory allocation:
		//
		// ```
		// name := zsw.Name("")
		// json.Unmarshal(data, &name)
		// ```
		//
		// This would work without a problem. In reflection code however, I
		// did not find how one can go from `reflect.Zero(typeGo)` (which is
		// the equivalence of doing `name := zsw.Name("")`) and take the
		// pointer to it so it can be unmarshalled correctly.
		//
		// A played with various iteration, and nothing got it working. Maybe
		// the next step would be to explore the `unsafe` package and obtain
		// an unsafe pointer and play with it.
		value := reflect.New(typeGo)
		if err = decoder.Decode(value.Interface()); err != nil {
			return fmt.Errorf("unable to decode variant type %d: %w", typeID, err)
		}

		a.Impl = value.Elem().Interface()
	}
	return nil
}

func twosComplement(v []byte) []byte {
	buf := make([]byte, len(v))
	for i, b := range v {
		buf[i] = b ^ byte(0xff)
	}
	one := big.NewInt(1)
	value := (&big.Int{}).SetBytes(buf)
	return value.Add(value, one).Bytes()
}

// Implementation of `fc::variant` types

type fcVariantType uint32

const (
	fcVariantNullType fcVariantType = iota
	fcVariantInt64Type
	fcVariantUint64Type
	fcVariantDoubleType
	fcVariantBoolType
	fcVariantStringType
	fcVariantArrayType
	fcVariantObjectType
	fcVariantBlobType
)

func (t fcVariantType) String() string {
	switch t {
	case fcVariantNullType:
		return "null"
	case fcVariantInt64Type:
		return "int64"
	case fcVariantUint64Type:
		return "uint64"
	case fcVariantDoubleType:
		return "double"
	case fcVariantBoolType:
		return "bool"
	case fcVariantStringType:
		return "string"
	case fcVariantArrayType:
		return "array"
	case fcVariantObjectType:
		return "object"
	case fcVariantBlobType:
		return "blob"
	}

	return "unknown"
}

// FIXME: Ideally, we would re-use `BaseVariant` but that requires some
//        re-thinking of the decoder to make it efficient to read FCVariant types. For now,
//        let's re-code it a bit to make it as efficient as possible.
type fcVariant struct {
	TypeID fcVariantType
	Impl   interface{}
}

func (a fcVariant) IsNil() bool {
	return a.TypeID == fcVariantNullType
}

// ToNative transform the actual implementation, walking each sub-element like array
// and object, turning everything along the way in Go primitives types.
//
// **Note** For `Int64` and `Uint64`, we return `zsw.Int64` and `zsw.Uint64` types
//          so that JSON marshalling is done correctly for large numbers
func (a fcVariant) ToNative() interface{} {
	if a.TypeID == fcVariantNullType ||
		a.TypeID == fcVariantDoubleType ||
		a.TypeID == fcVariantBoolType ||
		a.TypeID == fcVariantStringType {
		return a.Impl
	}

	if a.TypeID == fcVariantInt64Type {
		return Int64(a.Impl.(int64))
	}

	if a.TypeID == fcVariantUint64Type {
		return Uint64(a.Impl.(uint64))
	}

	if a.TypeID == fcVariantArrayType {
		return a.Impl.(fcVariantArray).ToNative()
	}

	if a.TypeID == fcVariantObjectType {
		return a.Impl.(fcVariantObject).ToNative()
	}

	panic(fmt.Errorf("not implemented for %s yet", fcVariantBlobType))
}

// MustAsUint64 casts the underlying `impl` as a `uint64` type, panics if not of the correct type.
func (a fcVariant) MustAsUint64() uint64 {
	return a.Impl.(uint64)
}

// MustAsString casts the underlying `impl` as a `string` type, panics if not of the correct type.
func (a fcVariant) MustAsString() string {
	return a.Impl.(string)
}

// MustAsObject casts the underlying `impl` as a `fcObject` type, panics if not of the correct type.
func (a fcVariant) MustAsObject() fcVariantObject {
	return a.Impl.(fcVariantObject)
}

func (a *fcVariant) UnmarshalBinary(decoder *zsw.Decoder) error {
	typeID, err := decoder.ReadUvarint32()
	if err != nil {
		return fmt.Errorf("unable to read fc variant type ID: %w", err)
	}

	if typeID > uint32(fcVariantBlobType) {
		return fmt.Errorf("invalid fc variant type ID, should have been lower than or equal to %d", fcVariantBlobType)
	}

	a.TypeID = fcVariantType(typeID)
	if a.TypeID == fcVariantNullType {
		// There is probably no bytes to read here, but it's not super clear
		a.Impl = nil
		return nil
	}

	if a.TypeID == fcVariantInt64Type {
		if a.Impl, err = decoder.ReadInt64(); err != nil {
			return fmt.Errorf("unable to read int64 fc variant: %w", err)
		}
	} else if a.TypeID == fcVariantUint64Type {
		if a.Impl, err = decoder.ReadUint64(); err != nil {
			return fmt.Errorf("unable to read uint64 fc variant: %w", err)
		}
	} else if a.TypeID == fcVariantDoubleType {
		if a.Impl, err = decoder.ReadFloat64(); err != nil {
			return fmt.Errorf("unable to read double fc variant: %w", err)
		}
	} else if a.TypeID == fcVariantBoolType {
		if a.Impl, err = decoder.ReadBool(); err != nil {
			return fmt.Errorf("unable to read bool fc variant: %w", err)
		}
	} else if a.TypeID == fcVariantStringType {
		if a.Impl, err = decoder.ReadString(); err != nil {
			return fmt.Errorf("unable to read string fc variant: %w", err)
		}
	} else if a.TypeID == fcVariantArrayType {
		out := fcVariantArray(nil)
		if err = decoder.Decode(&out); err != nil {
			return fmt.Errorf("unable to read fc array variant: %w", err)
		}
		a.Impl = out
	} else if a.TypeID == fcVariantObjectType {
		out := fcVariantObject{}
		if err = decoder.Decode(&out); err != nil {
			return fmt.Errorf("unable to read fc object variant: %w", err)
		}
		a.Impl = out
	} else if a.TypeID == fcVariantBlobType {
		// FIXME: This one is really not clear what the output format looks like, do we even need an object for it?
		var out fcVariantBlob
		if err = decoder.Decode(&out); err != nil {
			return fmt.Errorf("unable to read fc blob variant: %w", err)
		}
		a.Impl = out
	}

	return nil
}

type fcVariantArray []fcVariant

func (o fcVariantArray) ToNative() interface{} {
	out := make([]interface{}, len(o))
	for i, element := range o {
		out[i] = element.ToNative()
	}

	return out
}

func (o *fcVariantArray) UnmarshalBinary(decoder *zsw.Decoder) error {
	elementCount, err := decoder.ReadUvarint64()
	if err != nil {
		return fmt.Errorf("unable to read length: %w", err)
	}

	array := make([]fcVariant, elementCount)
	for i := uint64(0); i < elementCount; i++ {
		err := decoder.Decode(&array[i])
		if err != nil {
			return fmt.Errorf("unable to read elememt at index %d: %w", i, err)
		}
	}

	*o = fcVariantArray(array)
	return nil
}

type fcVariantObject map[string]fcVariant

func (o fcVariantObject) ToNative() map[string]interface{} {
	out := map[string]interface{}{}
	for key, value := range o {
		out[key] = value.ToNative()
	}

	return out
}

func (o fcVariantObject) validateFields(nameToType map[string]fcVariantType) error {
	for key, fcType := range nameToType {
		if len(key) <= 0 {
			continue
		}

		optional := false
		if string(key[0]) == "?" {
			key = key[1:]
			optional = true
		}

		actualType := o[key].TypeID
		if optional && actualType == fcVariantNullType {
			continue
		}

		if !optional && actualType == fcVariantNullType {
			return fmt.Errorf("field %q of type %s is required but actual type is null", key, fcType)
		}

		if actualType != fcType {
			return fmt.Errorf("field %q should be a variant of type %s, got %s", key, fcType, actualType)
		}
	}

	return nil
}

func (o *fcVariantObject) UnmarshalBinary(decoder *zsw.Decoder) error {
	elementCount, err := decoder.ReadUvarint64()
	if err != nil {
		return fmt.Errorf("unable to read length: %w", err)
	}

	mappings := make(map[string]fcVariant, elementCount)
	for i := uint64(0); i < elementCount; i++ {
		key, err := decoder.ReadString()
		if err != nil {
			return fmt.Errorf("unable to read key of elememt at index %d: %w", i, err)
		}

		variant := fcVariant{}
		err = decoder.Decode(&variant)
		if err != nil {
			return fmt.Errorf("unable to read value of elememt with key %s at index %d: %w", key, i, err)
		}

		mappings[key] = variant
	}

	*o = fcVariantObject(mappings)
	return nil
}

// FIXME: This one I'm unsure, is this correct at all?
type fcVariantBlob Blob

func (o *fcVariantBlob) UnmarshalBinary(decoder *zsw.Decoder) error {
	var blob Blob
	err := decoder.Decode(&blob)
	if err != nil {
		return err
	}

	*o = fcVariantBlob(blob)
	return nil
}

var ZSWAttributeVariant = NewVariantDefinition([]VariantType{
	{Name: "int8", Type: int8(0)},
	{Name: "int16", Type: int16(0)},
	{Name: "int32", Type: int32(0)},
	{Name: "int64", Type: int64(0)},
	{Name: "uint8", Type: uint8(0)},
	{Name: "uint16", Type: uint16(0)},
	{Name: "uint32", Type: uint32(0)},
	{Name: "uint64", Type: uint64(0)},
	{Name: "float", Type: float32(0)},
	{Name: "double", Type: float64(0)},
	{Name: "string", Type: ""},
	{Name: "INT8_VEC", Type: []int8{}},
	{Name: "INT16_VEC", Type: []int16{}},
	{Name: "INT32_VEC", Type: []int32{}},
	{Name: "INT64_VEC", Type: []int64{}},
	{Name: "UINT8_VEC", Type: []uint8{}},
	{Name: "UINT16_VEC", Type: []uint16{}},
	{Name: "UINT32_VEC", Type: []uint32{}},
	{Name: "UINT64_VEC", Type: []uint64{}},
	{Name: "FLOAT_VEC", Type: []float32{}},
	{Name: "DOUBLE_VEC", Type: []float64{}},
	{Name: "STRING_VEC", Type: []string{}},
})

type InvalidTypeError struct {
	Label        string
	ExpectedType string
	Attribute    *ZSWAttribute
}

func (c *InvalidTypeError) Error() string {
	return fmt.Sprintf("received an unexpected type %T for metadata variant %T", c.ExpectedType, c.Attribute)
}

type AttributeMap map[string]*ZSWAttribute

type ZSWAttribute struct {
	*BaseVariant
}

func NewZSWAttribute(typeId string, value interface{}) *ZSWAttribute {
	return &ZSWAttribute{
		&BaseVariant{
			TypeID: ZSWAttributeVariant.TypeID(typeId),
			Impl:   value,
		},
	}
}

func ToZSWAttribute(value interface{}) *ZSWAttribute {
	switch v := value.(type) {
	case float32:
		return NewZSWAttribute("float", v)
	case float64:
		return NewZSWAttribute("double", v)
	case []int8, []int16, []int32, []int64, []uint8, []uint16, []uint32, []uint64, []string:
		typeId := fmt.Sprintf("%v_VEC", strings.ToUpper(strings.ReplaceAll(fmt.Sprintf("%T", v), "[]", "")))
		return NewZSWAttribute(typeId, v)
	case []float32:
		return NewZSWAttribute("FLOAT_VEC", v)
	case []float64:
		return NewZSWAttribute("DOUBLE_VEC", v)
	default:
		return NewZSWAttribute(fmt.Sprintf("%T", v), v)
	}
}

func (m *ZSWAttribute) InvalidTypeError(expectedType string) *InvalidTypeError {
	return &InvalidTypeError{
		Label:        fmt.Sprintf("received an unexpected type %T for variant %T", m.Impl, m),
		ExpectedType: "int8",
		Attribute:    m,
	}
}

func (m *ZSWAttribute) String() string {
	return fmt.Sprint(m.Impl)
}

func (m *ZSWAttribute) Int8() (int8, error) {
	switch v := m.Impl.(type) {
	case int8:
		return v, nil
	default:
		return 0, m.InvalidTypeError("int8")
	}
}

func (m *ZSWAttribute) Int16() (int16, error) {
	switch v := m.Impl.(type) {
	case int16:
		return v, nil
	default:
		return 0, m.InvalidTypeError("int16")
	}
}

func (m *ZSWAttribute) Int32() (int32, error) {
	switch v := m.Impl.(type) {
	case int32:
		return v, nil
	default:
		return 0, m.InvalidTypeError("int32")
	}
}

func (m *ZSWAttribute) Int64() (int64, error) {
	switch v := m.Impl.(type) {
	case int64:
		return v, nil
	default:
		return 0, m.InvalidTypeError("int64")
	}
}

func (m *ZSWAttribute) UInt8() (uint8, error) {
	switch v := m.Impl.(type) {
	case uint8:
		return v, nil
	default:
		return 0, m.InvalidTypeError("uint8")
	}
}

func (m *ZSWAttribute) UInt16() (uint16, error) {
	switch v := m.Impl.(type) {
	case uint16:
		return v, nil
	default:
		return 0, m.InvalidTypeError("uint16")
	}
}

func (m *ZSWAttribute) UInt32() (uint32, error) {
	switch v := m.Impl.(type) {
	case uint32:
		return v, nil
	default:
		return 0, m.InvalidTypeError("uint32")
	}
}

func (m *ZSWAttribute) UInt64() (uint64, error) {
	switch v := m.Impl.(type) {
	case uint64:
		return v, nil
	default:
		return 0, m.InvalidTypeError("uint64")
	}
}

func (m *ZSWAttribute) Float32() (float32, error) {
	switch v := m.Impl.(type) {
	case float32:
		return v, nil
	default:
		return 0, m.InvalidTypeError("float32")
	}
}

func (m *ZSWAttribute) Float64() (float64, error) {
	switch v := m.Impl.(type) {
	case float64:
		return v, nil
	default:
		return 0, m.InvalidTypeError("float64")
	}
}

func (m *ZSWAttribute) Int8Slice() ([]int8, error) {
	switch v := m.Impl.(type) {
	case []int8:
		return v, nil
	default:
		return nil, m.InvalidTypeError("[]int8")
	}
}

func (m *ZSWAttribute) Int16Slice() ([]int16, error) {
	switch v := m.Impl.(type) {
	case []int16:
		return v, nil
	default:
		return nil, m.InvalidTypeError("[]int16")
	}
}

func (m *ZSWAttribute) Int32Slice() ([]int32, error) {
	switch v := m.Impl.(type) {
	case []int32:
		return v, nil
	default:
		return nil, m.InvalidTypeError("[]int32")
	}
}

func (m *ZSWAttribute) Int64Slice() ([]int64, error) {
	switch v := m.Impl.(type) {
	case []int64:
		return v, nil
	default:
		return nil, m.InvalidTypeError("[]int64")
	}
}

func (m *ZSWAttribute) UInt8Slice() ([]uint8, error) {
	switch v := m.Impl.(type) {
	case []uint8:
		return v, nil
	default:
		return nil, m.InvalidTypeError("[]uint8")
	}
}

func (m *ZSWAttribute) UInt16Slice() ([]uint16, error) {
	switch v := m.Impl.(type) {
	case []uint16:
		return v, nil
	default:
		return nil, m.InvalidTypeError("[]uint16")
	}
}

func (m *ZSWAttribute) UInt32Slice() ([]uint32, error) {
	switch v := m.Impl.(type) {
	case []uint32:
		return v, nil
	default:
		return nil, m.InvalidTypeError("[]uint32")
	}
}

func (m *ZSWAttribute) UInt64Slice() ([]uint64, error) {
	switch v := m.Impl.(type) {
	case []uint64:
		return v, nil
	default:
		return nil, m.InvalidTypeError("[]uint64")
	}
}

func (m *ZSWAttribute) Float32Slice() ([]float32, error) {
	switch v := m.Impl.(type) {
	case []float32:
		return v, nil
	default:
		return nil, m.InvalidTypeError("[]float32")
	}
}

func (m *ZSWAttribute) Float64Slice() ([]float64, error) {
	switch v := m.Impl.(type) {
	case []float64:
		return v, nil
	default:
		return nil, m.InvalidTypeError("[]float64")
	}
}

func (m *ZSWAttribute) StringSlice() ([]string, error) {
	switch v := m.Impl.(type) {
	case []string:
		return v, nil
	default:
		return nil, m.InvalidTypeError("[]string")
	}
}

// IsEqual evaluates if the two zswattr have the same types and values (deep compare)
func (m *ZSWAttribute) IsEqual(m2 *ZSWAttribute) bool {

	if m.TypeID != m2.TypeID {
		log.Println("ZSWAttribute types inequal: ", m.TypeID, " vs ", m2.TypeID)
		return false
	}

	if m.String() != m2.String() {
		log.Println("ZSWAttribute Values.String() inequal: ", m.String(), " vs ", m2.String())
		return false
	}

	return true
}

// MarshalJSON translates to []byte
func (m *ZSWAttribute) MarshalJSON() ([]byte, error) {
	return m.BaseVariant.MarshalJSON(ZSWAttributeVariant)
}

// UnmarshalJSON translates ZSWAttributeVariant
func (m *ZSWAttribute) UnmarshalJSON(data []byte) error {
	return m.BaseVariant.UnmarshalJSON(data, ZSWAttributeVariant)
}

// UnmarshalBinary ...
func (m *ZSWAttribute) UnmarshalBinary(decoder *zsw.Decoder) error {
	return m.BaseVariant.UnmarshalBinaryVariant(decoder, ZSWAttributeVariant)
}