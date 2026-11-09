package metadata

import (
	"context"
	"net/http"

	"github.com/dockloud/user-auth/pkg/transport/http/middleware"
	"google.golang.org/grpc/metadata"
)

func MetadataWithUID(ctx context.Context, r *http.Request) metadata.MD {
	md := metadata.Pairs()

	if uid, ok := middleware.UIDFromContext(ctx); ok {
		md.Set("x-user-id", uid)
	}

	return md
}
