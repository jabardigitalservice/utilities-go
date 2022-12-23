package context

import (
	basecontext "context"
)

// nolint: staticcheck
func WithValues(ctx basecontext.Context, mapValues map[string]interface{}) basecontext.Context {
	for k, v := range mapValues {
		ctx = basecontext.WithValue(ctx, k, v)
	}

	return ctx
}
