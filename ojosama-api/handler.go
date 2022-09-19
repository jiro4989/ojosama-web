package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/jiro4989/ojosama"
)

type RequestPostOjosama struct {
	Text string
}

type ResponseOjosama struct {
	Result string
}

// Handler is executed by AWS Lambda in the main function. Once the request
// is processed, it returns an Amazon API Gateway response object to AWS Lambda
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := request.Body

	var r RequestPostOjosama
	if err := json.Unmarshal([]byte(body), &r); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if r.Text == "" {
		return events.APIGatewayProxyResponse{}, nil
	}

	result, err := ojosama.Convert(r.Text, nil)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	var resp ResponseOjosama
	resp.Result = result

	b, err := json.Marshal(resp)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(b),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
	}, nil
}
