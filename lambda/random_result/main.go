package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type SqsBatchResponse struct {
	BatchItemFailures []BatchItemFailure `json:"batchItemFailures"`
}

type BatchItemFailure struct {
	ItemIdentifier string `json:"itemIdentifier"`
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) (res SqsBatchResponse, err error) {
	fmt.Printf("Received %d records\n", len(sqsEvent.Records))
	b, err := json.Marshal(sqsEvent)
	if err != nil {
		return
	}
	fmt.Println(string(b))

	for i, message := range sqsEvent.Records {
		fmt.Printf("The message %s for event source %s = %s \n", message.MessageId, message.EventSource, message.Body)
		if i%2 == 0 {
			fmt.Printf("Failuer message %s for event source %s = %s \n", message.MessageId, message.EventSource, message.Body)
			res.BatchItemFailures = append(res.BatchItemFailures, BatchItemFailure{message.MessageId})
		}
	}
	b, err = json.Marshal(res)
	if err != nil {
		return
	}
	fmt.Println(string(b))

	return
}

func main() {
	lambda.Start(handler)
}
