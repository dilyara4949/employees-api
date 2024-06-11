package server

import (
	"context"
	"log"
	"time"

	"github.com/dilyara4949/employees-api/internal/middleware"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()
	correlationID := ctx.Value(middleware.CorrelationID)
	h, err := handler(ctx, req)

	log.Printf("Method: %s, CorrelationID: %s, Request: %+v, Duration: %s, Error: %v", info.FullMethod, correlationID, req, time.Since(start), err)

	return h, err
}

func CorrelationIDInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		correlationID := getCorrelationIDFromContext(ctx)
		ctx = context.WithValue(ctx, middleware.CorrelationID, correlationID)
		return handler(ctx, req)
	}
}

func getCorrelationIDFromContext(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if values := md[middleware.CorrelationID]; len(values) > 0 {
			return values[0]
		}
	}
	return uuid.New().String()
}
