package logger

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/pwmcintyre/logging/go/appcontext"
)

// WithContext returns a logger with known fields from context
// It is recommended to use this anytime you log inside your microservice
func WithContext(ctx context.Context) *logrus.Entry {
	return logrus.StandardLogger().WithFields(ContextFields(ctx))
}

// ContextFields returns loggable fields from context
func ContextFields(ctx context.Context) logrus.Fields {

	f := logrus.Fields{}
	if ctx == nil {
		return f
	}

	if val, ok := appcontext.GetSystemContext(ctx); ok {
		f["system"] = val
	}

	if val, ok := appcontext.GetHTTPRequestContext(ctx); ok {
		f["request"] = val
	}

	if val, ok := appcontext.GetClientContext(ctx); ok {
		f["client"] = val
	}

	if val, ok := appcontext.GetCorrelationID(ctx); ok {
		f["correlation_id"] = val
	}

	if val, ok := appcontext.GetAWSRequestID(ctx); ok {
		f["aws_request_id"] = val
	}

	return f
}
