package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var (
	endpoint          = "http://localhost.localstack.cloud:4566"
	queue             = "http://sqs.eu-central-1.localhost.localstack.cloud:4566/000000000000/test"
	maxMessages int64 = 10
	visibility  int64 = 20
)

func main() {
	config := &aws.Config{
		Endpoint: &endpoint,
		Region:   aws.String("eu-central-1"),
	}

	link, err := session.NewSession(config)
	if err != nil {
		fmt.Println("Error creating AWS session:", err)
		return
	}

	svc := sqs.New(link)

	for {
		input := &sqs.ReceiveMessageInput{
			QueueUrl:            &queue,
			MaxNumberOfMessages: &maxMessages,
			VisibilityTimeout:   &visibility,
		}

		result, err := svc.ReceiveMessageWithContext(context.TODO(), input)
		if err != nil {
			fmt.Println("Error receiving messages:", err)
			continue
		}

		for _, message := range result.Messages {
			input := &sqs.DeleteMessageInput{
				QueueUrl:      aws.String(queue),
				ReceiptHandle: message.ReceiptHandle,
			}

			if _, err := svc.DeleteMessageWithContext(context.TODO(), input); err != nil {
				fmt.Println("Error deleting message:", err)
			}
		}
	}
}
