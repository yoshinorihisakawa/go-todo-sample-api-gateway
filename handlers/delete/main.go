package main

import (
	"fmt"
	"strconv"

	"github.com/SrcHndWng/go-todo-sample-api-gateway/model/todo"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is the only one entry point.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var id int
	var err error

	if id, err = strconv.Atoi(request.PathParameters["id"]); err != nil {
		return errorResponse(err)
	}
	if err = todo.Delete(id); err != nil {
		return errorResponse(err)
	}
	return successResponse("")
}

func successResponse(body string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: body}, nil
}

func errorResponse(err error) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("%+v\n", err)
	return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Internal Server Error!"}, nil
}

func main() {
	lambda.Start(Handler)
}
