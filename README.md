# SQS golang worker

* v1
* v2

## Localstack

### List SQS queues

`awslocal sqs list-queues`

### Create an SQS queue on LocalStack

`awslocal sqs create-queue --queue-name test`

### Add message to the queue

`awslocal sqs send-message --queue-url http://sqs.eu-central-1.localhost.localstack.cloud:4566/000000000000/test --message-body '{ "message": "this is a test" }'`

### List 10 first messages

`awslocal sqs receive-message --queue-url <http://sqs.eu-central-1.localhost.localstack.cloud:4566/000000000000/test> --max-number-of-messages 10`
