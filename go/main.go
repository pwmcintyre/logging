package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"

	"github.com/pwmcintyre/logging/go/appcontext"
	"github.com/pwmcintyre/logging/go/logger"
)

func Handle(ctx context.Context, event events.APIGatewayProxyRequest) error {

	// harvest context for logging
	// https://docs.aws.amazon.com/lambda/latest/dg/golang-context.html
	if lc, ok := lambdacontext.FromContext(ctx); ok {
		ctx = appcontext.WithAWSRequestID(ctx, appcontext.AWSRequestID(lc.AwsRequestID))
	}

	ctx = appcontext.WithCorrelationID(ctx, appcontext.CorrelationID(event.Headers["X-Correlation-Id"]))

	// logger pacakge which harvests `context.Context`
	logger.WithContext(ctx).Info("foo")

	return nil

}

type Observer interface {
	didThing()
}

type service struct {
	observer Observer
}

func (s *server) thing() {

	// do thing

	// log
	logger.log("did thing")
	metrics.emit(1, "thing")

}

func (s *server) thing() {

	// log
	s.observer.doThing(func() {
		//do thing
	})

	lambda.Handle(datadog.Trace(handler))

}
