// cometbft_connection/cometbft_connection.go
package client

import (
	"github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/cosmos-sdk/client"
)

// NewCometBFTClient initializes a new client
func NewCometBFTClient() *http.HTTP {
	nodeUrl := "https://seda-testnet.rpc.kjnodes.com"
	httpClient, err := client.NewClientFromNode(nodeUrl)
	if err != nil {
		panic(err)
	}

	// Start the client
	err = httpClient.Start()
	if err != nil {
		panic(err)
	}
	return httpClient
}
