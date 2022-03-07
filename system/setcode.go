package system

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	zsw "github.com/zhongshuwen/zswchain-go"
)

func NewSetContract(account zsw.AccountName, wasmPath, abiPath string) (out []*zsw.Action, err error) {
	codeAction, err := NewSetCode(account, wasmPath)
	if err != nil {
		return nil, err
	}

	abiAction, err := NewSetABI(account, abiPath)
	if err != nil {
		return nil, err
	}

	return []*zsw.Action{codeAction, abiAction}, nil
}

func NewSetContractContent(account zsw.AccountName, wasmContent, abiContent []byte) (out []*zsw.Action, err error) {
	codeAction := NewSetCodeContent(account, wasmContent)

	abiAction, err := NewSetAbiContent(account, abiContent)
	if err != nil {
		return nil, err
	}

	return []*zsw.Action{codeAction, abiAction}, nil
}

func NewSetCode(account zsw.AccountName, wasmPath string) (out *zsw.Action, err error) {
	codeContent, err := ioutil.ReadFile(wasmPath)
	if err != nil {
		return nil, err
	}
	return NewSetCodeContent(account, codeContent), nil
}

func NewSetCodeContent(account zsw.AccountName, codeContent []byte) *zsw.Action {
	return &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("setcode"),
		Authorization: []zsw.PermissionLevel{
			{
				Actor:      account,
				Permission: zsw.PermissionName("active"),
			},
		},
		ActionData: zsw.NewActionData(SetCode{
			Account:   account,
			VMType:    0,
			VMVersion: 0,
			Code:      zsw.HexBytes(codeContent),
		}),
	}
}

func NewSetABI(account zsw.AccountName, abiPath string) (out *zsw.Action, err error) {
	abiContent, err := ioutil.ReadFile(abiPath)
	if err != nil {
		return nil, err
	}

	return NewSetAbiContent(account, abiContent)
}

func NewSetAbiContent(account zsw.AccountName, abiContent []byte) (out *zsw.Action, err error) {
	var abiPacked []byte
	if len(abiContent) > 0 {
		var abiDef zsw.ABI
		if err := json.Unmarshal(abiContent, &abiDef); err != nil {
			return nil, fmt.Errorf("unmarshal ABI file: %w", err)
		}

		abiPacked, err = zsw.MarshalBinary(abiDef)
		if err != nil {
			return nil, fmt.Errorf("packing ABI: %w", err)
		}
	}

	return &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("setabi"),
		Authorization: []zsw.PermissionLevel{
			{
				Actor:      account,
				Permission: zsw.PermissionName("active"),
			},
		},
		ActionData: zsw.NewActionData(SetABI{
			Account: account,
			ABI:     zsw.HexBytes(abiPacked),
		}),
	}, nil
}

func NewSetAbiFromAbi(account zsw.AccountName, abi zsw.ABI) (out *zsw.Action, err error) {
	var abiPacked []byte
	abiPacked, err = zsw.MarshalBinary(abi)
	if err != nil {
		return nil, fmt.Errorf("packing ABI: %w", err)
	}

	return &zsw.Action{
		Account: AN("zswhq"),
		Name:    ActN("setabi"),
		Authorization: []zsw.PermissionLevel{
			{
				Actor:      account,
				Permission: zsw.PermissionName("active"),
			},
		},
		ActionData: zsw.NewActionData(SetABI{
			Account: account,
			ABI:     zsw.HexBytes(abiPacked),
		}),
	}, nil
}

// NewSetCodeTx is _deprecated_. Use NewSetContract instead, and build
// your transaction yourself.
func NewSetCodeTx(account zsw.AccountName, wasmPath, abiPath string) (out *zsw.Transaction, err error) {
	actions, err := NewSetContract(account, wasmPath, abiPath)
	if err != nil {
		return nil, err
	}
	return &zsw.Transaction{Actions: actions}, nil
}

// SetCode represents the hard-coded `setcode` action.
type SetCode struct {
	Account   zsw.AccountName `json:"account"`
	VMType    byte            `json:"vmtype"`
	VMVersion byte            `json:"vmversion"`
	Code      zsw.HexBytes    `json:"code"`
}

// SetABI represents the hard-coded `setabi` action.
type SetABI struct {
	Account zsw.AccountName `json:"account"`
	ABI     zsw.HexBytes    `json:"abi"`
}
