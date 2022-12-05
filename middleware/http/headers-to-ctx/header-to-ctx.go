package headerstoctx

import (
	"context"
	"net/http"
)

// nolint: staticcheck
func Mapping(maps map[string]interface{}) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			for ctxKey, headerKey := range maps {
				ctxVal := r.Header.Get(headerKey.(string))
				ctx = context.WithValue(ctx, ctxKey, ctxVal)
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
