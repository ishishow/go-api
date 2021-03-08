package interceptor

import (
	"context"
	"log"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	// XRequestIDKey is a key for getting request id.
	XRequestIDKey    = "x-token"
	unknownRequestID = "<unknown>"
)

// RequestIDInterceptor is a interceptor of access control list.
func RequestIDInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		requestID := requestIDFromContext(ctx)
		ctx = context.WithValue(ctx, "key", requestID)
		return handler(ctx, req)
	}
}

func requestIDFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("metadata.FromIncomingContext(ctx) is failed")
		return unknownRequestID
	}

	key := strings.ToLower(XRequestIDKey)
	header, ok := md[key]
	if !ok || len(header) == 0 {
		log.Println("key is missing")
		return unknownRequestID
	}

	requestID := header[0]
	if requestID == "" {
		log.Println("requestID is empty")
		return unknownRequestID
	}

	return requestID
}
