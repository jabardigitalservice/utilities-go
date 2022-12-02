package headerstoctx

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMapping(t *testing.T) {

	mappingHandler := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		for ctxKey, ctxValue := range ctxtValuesExpectedTest {
			assert.Equal(t, ctxValue, ctx.Value(ctxKey).(string))
		}

	}

	req := httptest.NewRequest(http.MethodGet, "http://www.your-domain.com", nil)

	for ctxKey, ctxValue := range ctxtValuesExpectedTest {
		headerKey := mapsTest[ctxKey]
		req.Header.Add(headerKey.(string), ctxValue.(string))
	}

	res := httptest.NewRecorder()

	mapping := Mapping(context.Background(), mapsTest, mappingHandler)
	mapping.ServeHTTP(res, req)

}

var (
	mapsTest = map[string]interface{}{
		"SessionID": "X-Session-ID",
		"UserID":    "X-User-ID",
	}

	ctxtValuesExpectedTest = map[string]interface{}{
		"SessionID": "SessionIDValue",
		"UserID":    "UserIDValue",
	}
)
