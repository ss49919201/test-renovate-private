package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.CloudWatchEvent) error {
	b, err := json.Marshal(event)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

func main() {
	lambda.Start(handler)
}
