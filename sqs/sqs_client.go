package sqs

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// NewSQSService initializes a new SQS service with provided endpoint
func NewSQSService(endpoint string) *sqs.SQS {
	// Create AWS session
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		),
		Endpoint: aws.String(endpoint),
		Region:   aws.String("us-east-1"),
	}))
	return sqs.New(sess)
}
