package zsw_test

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	zsw "github.com/zhongshuwen/zswchain-go"
	"github.com/zhongshuwen/zswchain-go/token"
)

func ExampleAPI_PushTransaction_transfer_EOS() {
	api := zsw.New(getAPIURL())

	keyBag := &zsw.KeyBag{}
	err := keyBag.ImportPrivateKey(context.Background(), readPrivateKey())
	if err != nil {
		panic(fmt.Errorf("import private key: %w", err))
	}
	api.SetSigner(keyBag)

	from := zsw.AccountName("zswuser1")
	to := zsw.AccountName("zswuser2")
	quantity, err := zsw.NewZSWAssetFromString("1.0000 ZSWCC")
	memo := ""

	if err != nil {
		panic(fmt.Errorf("invalid quantity: %w", err))
	}

	txOpts := &zsw.TxOptions{}
	if err := txOpts.FillFromChain(context.Background(), api); err != nil {
		panic(fmt.Errorf("filling tx opts: %w", err))
	}

	tx := zsw.NewTransaction([]*zsw.Action{token.NewTransfer(from, to, quantity, memo)}, txOpts)
	signedTx, packedTx, err := api.SignTransaction(context.Background(), tx, txOpts.ChainID, zsw.CompressionNone)
	if err != nil {
		panic(fmt.Errorf("sign transaction: %w", err))
	}

	content, err := json.MarshalIndent(signedTx, "", "  ")
	if err != nil {
		panic(fmt.Errorf("json marshalling transaction: %w", err))
	}

	fmt.Println(string(content))
	fmt.Println()

	response, err := api.PushTransaction(context.Background(), packedTx)
	if err != nil {
		panic(fmt.Errorf("push transaction: %w", err))
	}

	fmt.Printf("Transaction [%s] submitted to the network succesfully.\n", hex.EncodeToString(response.Processed.ID))
}

func readPrivateKey() string {
	// Right now, the key is read from an environment variable, it's an example after all.
	// In a real-world scenario, would you probably integrate with a real wallet or something similar
	envName := "ZSW_CHAIN_PRIVATE_KEY"
	privateKey := os.Getenv(envName)
	if privateKey == "" {
		panic(fmt.Errorf("private key environment variable %q must be set", envName))
	}

	return privateKey
}
