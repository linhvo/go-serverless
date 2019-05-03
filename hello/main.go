package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"

	//"github.com/aws/aws-lambda-go/lambda"

)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	//var buf bytes.Buffer
	//
	//body, err := json.Marshal(map[string]interface{}{
	//	"message": "Go Serverless v1.0! Your function executed successfully!",
	//})
	//if err != nil {
	//	return Response{StatusCode: 404}, err
	//}
	//json.HTMLEscape(&buf, body)
	//
	//resp := Response{
	//	StatusCode:      200,
	//	IsBase64Encoded: false,
	//	Body:            buf.String(),
	//	Headers: map[string]string{
	//		"Content-Type":           "application/json",
	//		"X-MyCompany-Func-Reply": "hello-handler",
	//	},
	//}
	//
	//return resp, nil
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)
	resp := Response{}

	// URL to our queue
	qURL := "https://sqs.us-east-1.amazonaws.com/502004135298/sqs-queue"
	result, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Title": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("The Whistler"),
			},
			"Author": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("John Grisham"),
			},
			"WeeksOn": &sqs.MessageAttributeValue{
				DataType:    aws.String("Number"),
				StringValue: aws.String("6"),
			},
		},
		MessageBody: aws.String("Test"),
		QueueUrl:    &qURL,
	})


	if err != nil {
		fmt.Println("Error", err)
		return resp, err
	}
	fmt.Println("Success", *result.MessageId)
	resp = Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            *result.MessageId,
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil

}

func main() {
	lambda.Start(Handler)

}
