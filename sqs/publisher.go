package sqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// PublishToQueue publishes a message to an SQS queue
func PublishToQueue(sqsSvc *sqs.SQS, queueURL, message string) error {
	sendMessageInput := &sqs.SendMessageInput{
		MessageBody: aws.String(message),
		QueueUrl:    aws.String(queueURL),
	}

	_, err := sqsSvc.SendMessage(sendMessageInput)
	if err != nil {
		return err
	}

	return nil
}
