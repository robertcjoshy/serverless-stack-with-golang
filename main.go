package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/robert/serverless-stack-with-golang/src/handlers"
)

var ginLambda *ginadapter.GinLambda

func init() {
	r := gin.Default()

	r.POST("/app", handlers.CreateUser)

	r.GET("app/:email", handlers.GetUser)

	//r.DELETE("/app/det", handlers.DeleteUser)

	//r.PUT("/app/update", handlers.UpdateUser)

	ginLambda = ginadapter.New(r)
}

func main() {
	lambda.Start(lambdahandler)
}

func lambdahandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}
