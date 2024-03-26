package block

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/cometbft/cometbft/types"
	"github.com/uwezukwechibuzor/indexing/client"
	queue "github.com/uwezukwechibuzor/indexing/sqs"
)

// ListenAndPublishBlocksToSQS subscribes to block events and publishes them to SQS
func ListenAndPublishBlocksToSQS(sqsSvc *sqs.SQS, queueURL string) {

	// Initialize the COMETBFT client
	httpClient := client.NewCometBFTClient()

	eventChan, err := httpClient.Subscribe(context.Background(), "example", "tm.event = 'NewBlock'")
	if err != nil {
		panic(err)
	}

	for eventReceived := range eventChan {

		// Process block event
		block := eventReceived.Data.(types.EventDataNewBlock).Block

		// Prepare message to send to SQS
		message := buildBlockMessage(block)

		// Publish message to SQS queue
		err := queue.PublishToQueue(sqsSvc, queueURL, message)
		if err != nil {
			fmt.Println("Error publishing message to SQS:", err)
			continue
		}

		// Print confirmation
		fmt.Println("Published block event to SQS")
	}
}

// buildBlockMessage returns selected block data fields from a block event
// block height
// Total Number of Txs
// Time
// Proposer Address
func buildBlockMessage(block *types.Block) string {
	return fmt.Sprintf("New Block Height: %d, Time: %s, Number of Txs: %d, Proposer Hex: %s",
		block.Header.Height, block.Header.Time.Format(time.RFC3339), len(block.Data.Txs), block.Header.ProposerAddress)
}
