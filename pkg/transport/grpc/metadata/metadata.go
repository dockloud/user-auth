package metadata

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func GetUIDFromMetadata(ctx context.Context) (string, bool) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if uids := md.Get("x-user-id"); len(uids) > 0 {
			return uids[0], true
		}
	}
	return "", false
}
