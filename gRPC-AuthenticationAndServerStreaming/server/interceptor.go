// server/interceptor.go
package main

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func authInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("missing metadata")
	}

	token := md["authorization"]
	if len(token) == 0 || token[0] != "Bearer secret123" {
		return nil, errors.New("unauthorized")
	}

	// Token valid → call actual RPC
	return handler(ctx, req)
}
