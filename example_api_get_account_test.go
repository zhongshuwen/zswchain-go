package zsw_test

import (
	"context"
	"encoding/json"
	"fmt"

	zsw "github.com/zhongshuwen/zswchain-go"
)

func ExampleAPI_GetAccount() {
	api := zsw.New(getAPIURL())

	account := zsw.AccountName("zsw.rex")
	info, err := api.GetAccount(context.Background(), account)
	if err != nil {
		if err == zsw.ErrNotFound {
			fmt.Printf("unknown account: %s", account)
			return
		}

		panic(fmt.Errorf("get account: %w", err))
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		panic(fmt.Errorf("json marshal response: %w", err))
	}

	fmt.Println(string(bytes))
}
