package context

import (
	basecontext "context"
	"testing"

	"github.com/go-playground/assert/v2"
)

var (
	mapValues = map[string]interface{}{
		"keyA": "valueA",
		"keyB": "valueB",
		"keyC": "valueC",
	}
)

func TestMapping(t *testing.T) {
	ctx := basecontext.Background()

	ctx = WithValues(ctx, mapValues)

	for k, v := range mapValues {
		assert.Equal(t, v, ctx.Value(k).(string))
	}
}
