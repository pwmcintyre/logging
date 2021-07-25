package appcontext

import "context"

const (
	systemContextKey int = iota
	requestContextKey
	authContextKey
	correlationKey
	awsRequestKey
)

type SystemContext struct {
	Application string `json:"application,omitempty"`
	Version     string `json:"version,omitempty"`
	Environment string `json:"environment,omitempty"`
}

func WithSystemContext(ctx context.Context, val SystemContext) context.Context {
	return context.WithValue(ctx, systemContextKey, val)
}

func GetSystemContext(ctx context.Context) (val SystemContext, ok bool) {
	val, ok = ctx.Value(systemContextKey).(SystemContext)
	return
}

type HTTPRequestContext struct {
	Method    string `json:"method,omitempty"`
	Path      string `json:"path,omitempty"`
	Query     string `json:"query,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

func WithHTTPRequestContext(ctx context.Context, val HTTPRequestContext) context.Context {
	return context.WithValue(ctx, requestContextKey, val)
}

func GetHTTPRequestContext(ctx context.Context) (val HTTPRequestContext, ok bool) {
	val, ok = ctx.Value(requestContextKey).(HTTPRequestContext)
	return
}

type ClientContext struct {
	ClientID string `json:"client_id"`
}

func WithClientContext(ctx context.Context, val ClientContext) context.Context {
	return context.WithValue(ctx, authContextKey, val)
}

func GetClientContext(ctx context.Context) (val ClientContext, ok bool) {
	val, ok = ctx.Value(authContextKey).(ClientContext)
	return
}

type CorrelationID string

func WithCorrelationID(ctx context.Context, val CorrelationID) context.Context {
	return context.WithValue(ctx, authContextKey, val)
}

func GetCorrelationID(ctx context.Context) (val CorrelationID, ok bool) {
	val, ok = ctx.Value(correlationKey).(CorrelationID)
	return
}

type AWSRequestID string

func WithAWSRequestID(ctx context.Context, val AWSRequestID) context.Context {
	return context.WithValue(ctx, authContextKey, val)
}

func GetAWSRequestID(ctx context.Context) (val AWSRequestID, ok bool) {
	val, ok = ctx.Value(awsRequestKey).(AWSRequestID)
	return
}
