package controller

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc/metadata"
)

func appendHeaderAccessToken(header http.Header, ctx context.Context) context.Context {
	fmt.Println(header.Get("access-token"))
	return metadata.AppendToOutgoingContext(ctx, "access-token", header.Get("access-token"))
}
