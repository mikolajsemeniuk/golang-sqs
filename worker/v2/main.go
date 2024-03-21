package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

var (
	endpoint                = "http://localhost:4566" // Update if needed
	queue                   = "http://sqs.eu-central-1.localhost.localstack.cloud:4566/000000000000/test"
	maxNumberOfMsgs   int32 = 10
	visibilityTimeout int32 = 30
)

func main() {
	c := aws.Config{
		BaseEndpoint: &endpoint,
		Region:       "eu-central-1",
	}

	svc := sqs.NewFromConfig(c)

	for {
		input := &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(queue),
			MaxNumberOfMessages: maxNumberOfMsgs,
			VisibilityTimeout:   visibilityTimeout,
		}

		result, err := svc.ReceiveMessage(context.Background(), input)
		if err != nil {
			fmt.Println("Error receiving messages:", err)
			continue
		}

		for _, message := range result.Messages {
			input := &sqs.DeleteMessageInput{
				QueueUrl:      aws.String(queue),
				ReceiptHandle: message.ReceiptHandle,
			}

			if _, err := svc.DeleteMessage(context.TODO(), input); err != nil {
				fmt.Println("Error deleting message:", err)
			}
		}
	}
}
