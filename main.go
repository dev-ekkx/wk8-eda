package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type Response struct {
	Message string `json:"message"`
}

func handler(ctx context.Context, s3Event events.S3Event) (Response, error) {
	// SNS client
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return Response{Message: "Failed to load AWS config"}, err
	}
	snsClient := sns.NewFromConfig(cfg)
	topicArn := "arn:aws:sns:eu-west-1:124355655853:wk8-eda-topic-dev"

	// Process S3 event
	for _, record := range s3Event.Records {
		bucket := record.S3.Bucket.Name
		key := record.S3.Object.Key
		message := fmt.Sprintf("New object uploaded to S3: %s/%s", bucket, key)

		// Publish to SNS
		_, err := snsClient.Publish(ctx, &sns.PublishInput{
			Message:  &message,
			TopicArn: &topicArn,
		})
		if err != nil {
			return Response{Message: "Failed to publish to SNS"}, err
		}
	}

	return Response{Message: "Notification sent successfully"}, nil
}

func main() {
	lambda.Start(handler)
}
