package zsw_test

import (
	"encoding/hex"
	"fmt"
	"strings"

	zsw "github.com/zhongshuwen/zswchain-go"
)

func ExampleABI_DecodeTableRowTyped() {
	abi, err := zsw.NewABI(strings.NewReader(abiJSON()))
	if err != nil {
		panic(fmt.Errorf("get ABI: %w", err))
	}

	tableDef := abi.TableForName(zsw.TableName("votes"))
	if tableDef == nil {
		panic(fmt.Errorf("table be should be present"))
	}

	bytes, err := abi.DecodeTableRowTyped(tableDef.Type, data())
	if err != nil {
		panic(fmt.Errorf("decode row: %w", err))
	}

	fmt.Println(string(bytes))
}

func data() []byte {
	bytes, err := hex.DecodeString(`424e544441505000`)
	if err != nil {
		panic(fmt.Errorf("decode data: %w", err))
	}

	return bytes
}

func abiJSON() string {
	return `{
			"structs": [
				{
					"name": "vote_t",
					"fields": [
						{ "name": "symbl", "type": "symbol_code" }
					]
				}
			],
			"actions": [],
			"tables": [
				{
					"name": "votes",
					"type": "vote_t"
				}
			]
	}`
}
