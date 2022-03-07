package zsw_test

import "os"

func getAPIURL() string {
	apiURL := os.Getenv("ZSW_CHAIN_API_URL")
	if apiURL != "" {
		return apiURL
	}

	return "https://api.eosn.io/"
}
