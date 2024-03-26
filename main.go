package main

import (
	"os"

	"github.com/uwezukwechibuzor/indexing/client"
	"github.com/uwezukwechibuzor/indexing/messages/block"
	"github.com/uwezukwechibuzor/indexing/sqs"
)

func main() {
	// Load environment variables
	sqsQueueURL := os.Getenv("SQS_QUEUE_URL")
	sqsEndpoint := os.Getenv("SQS_ENDPOINT")

	// Initialize COMETBFT connection
	client.NewCometBFTClient()

	// Initialize SQS connection
	sqsSvc := sqs.NewSQSService(sqsEndpoint)

	// Subscribe to block events and publish to SQS
	block.ListenAndPublishBlocksToSQS(sqsSvc, sqsQueueURL)
}
