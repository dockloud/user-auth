package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dockloud/user-auth/pkg/infrastructure/verifier"
)

type contextKey struct{}

func UserAuthMiddleware(verifier *verifier.Verifier) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			if strings.HasSuffix(path, "/auth") || strings.HasSuffix(path, "/register") {
				next.ServeHTTP(w, r)
				return
			}

			header := r.Header.Get("Authorization")
			parts := strings.Fields(header) // ["Bearer", "auth_token_here"]

			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Unaithorized", http.StatusUnauthorized)
				return
			}

			uid, err := verifier.Verify(parts[1])
			if err != nil {
				http.Error(w, "Unaithorized", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(
				r.Context(),
				contextKey{},
				uid,
			)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func UIDFromContext(ctx context.Context) (string, bool) {
	uid, ok := ctx.Value(contextKey{}).(string)
	return uid, ok
}
