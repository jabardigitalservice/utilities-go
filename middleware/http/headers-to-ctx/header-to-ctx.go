package headerstoctx

import (
	"context"
	"net/http"
)

// nolint: staticcheck
func Mapping(ctx context.Context, maps map[string]interface{}, h http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {

		for ctxKey, headerKey := range maps {
			ctxVal := r.Header.Get(headerKey.(string))
			ctx = context.WithValue(ctx, ctxKey, ctxVal)
		}

		h.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
