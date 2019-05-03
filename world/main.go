package main

import (

	"context"
	//"encoding/json"
	"strings"

	"fmt"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-lambda-go/events"
	//"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/aws/session"
	//"github.com/aws/aws-sdk-go/service/sqs"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, sqsEvent events.SQSEvent) (Response, error) {
	var results []string
	for _, message := range sqsEvent.Records {
		fmt.Printf("The message %s for event source %s = %s \n", message.MessageId, message.EventSource, message.Body)
		results = append(results, message.Body)
		fmt.Println(strings.Join(results, ", "))
	}

	//sess := session.Must(session.NewSessionWithOptions(session.Options{
	//	SharedConfigState: session.SharedConfigEnable,
	//}))
	//
	//svc := sqs.New(sess)
	//qURL := "https://sqs.us-east-1.amazonaws.com/502004135298/sqs-queue"
	//resp := Response{}
	//result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
	//	AttributeNames: []*string{
	//		aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
	//	},
	//	MessageAttributeNames: []*string{
	//		aws.String(sqs.QueueAttributeNameAll),
	//	},
	//	QueueUrl:            &qURL,
	//	MaxNumberOfMessages: aws.Int64(10),
	//	VisibilityTimeout:   aws.Int64(60), // 60 seconds
	//	WaitTimeSeconds:     aws.Int64(0),
	//})
	//if err != nil {
	//	fmt.Println("Error", err)
	//	return resp, err
	//}
	//if len(result.Messages) == 0 {
	//	resp.Body = "Received no messages"
	//	return resp, nil
	//}
	//
	//body := fmt.Sprintf("Success: %+v\n", result.Messages)
	//body, _ := json.Marshal(results)
	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:     strings.Join(results, "\n")       ,
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "world-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)

}
