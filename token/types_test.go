package token

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/zhongshuwen/zswchain-go"
	"github.com/stretchr/testify/assert"
)

//func TestPackAction(t *testing.T) {
//	a := &zsw.Action{
//		Account: AN("zswhq"),
//		Name:    ActN("transfer"),
//		Authorization: []zsw.PermissionLevel{
//			{AN("zswhq"), PN("active")},
//		},
//		Data: Transfer{
//			From:     AN("abourget"),
//			To:       AN("zswhq"),
//			Quantity: zsw.Asset{Amount: 123123, Symbol: zsw.ZSWCCSymbol},
//		},
//	}
//
//	buf, err := zsw.MarshalBinary(a)
//	assert.NoError(t, err)
//	assert.Equal(t, `0000000000ea3055000000572d3ccdcd010000000000ea305500000000a8ed32322100000059b1abe9310000000000ea3055f3e001000000000004454f530000000000`, hex.EncodeToString(buf))
//
//	buf, err = json.Marshal(a)
//	assert.NoError(t, err)
//	assert.Equal(t, `{"account":"zswhq","authorization":[{"actor":"zswhq","permission":"active"}],"data":"00000059b1abe9310000000000ea3055f3e001000000000004454f530000000000","name":"transfer"}`, string(buf))
//
//	/* 0000000000ea3055 000000572d3ccdcd 01 0000000000ea3055 00000000a8ed3232
//	   21
//	   00000059b1abe931 0000000000ea3055 f3e0010000000000 04 454f5300000000 00 */
//}

func TestUnpackActionTransfer(t *testing.T) {
	tests := []struct {
		in  string
		out Transfer
	}{
		{
			"00000003884ed1c900000000884ed1c90900000000000000000000000000000000",
			Transfer{AN("tbcox2.3"), AN("tbcox2"), zsw.Asset{Amount: 9}, ""},
		},
		{
			"00000003884ed1c900000000884ed1c90900000000000000000000000000000004616c6c6f",
			Transfer{AN("tbcox2.3"), AN("tbcox2"), zsw.Asset{Amount: 9}, "allo"},
		},
	}

	for idx, test := range tests {
		buf, err := hex.DecodeString(test.in)
		assert.NoError(t, err)

		var res Transfer
		assert.NoError(t, zsw.UnmarshalBinary(buf, &res), fmt.Sprintf("Index %d", idx))
		assert.Equal(t, test.out, res)
	}

}
