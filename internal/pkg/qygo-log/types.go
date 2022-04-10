package qygo_log

import "go.uber.org/zap"

type key int

// Defines common log fields.
const (
	KeyRequestID string = "requestID"
	KeyUsername  string = "username"
)
const (
	logContextKey key = iota
)

type Field = zap.Field
