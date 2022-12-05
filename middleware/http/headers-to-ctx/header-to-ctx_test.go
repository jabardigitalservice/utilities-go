package headerstoctx

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/go-playground/assert/v2"
)

func TestMappingUsingGoChi(t *testing.T) {

	r := chi.NewRouter()
	r.Use(Mapping(mapsTest))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		w.Header().Set("X-Session-ID", ctx.Value("SessionID").(string))
		w.Header().Set("X-User-ID", ctx.Value("UserID").(string))
		w.Write([]byte("ok"))
	})

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest(http.MethodGet, ts.URL+"/", nil)
	if err != nil {
		t.Fatal(err)
	}

	for ctxKey, ctxValue := range ctxtValuesExpectedTest {
		headerKey := mapsTest[ctxKey]
		req.Header.Add(headerKey.(string), ctxValue.(string))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	for ctxKey, headerKey := range mapsTest {
		assert.Equal(t, ctxtValuesExpectedTest[ctxKey], resp.Header.Get(headerKey.(string)))
	}

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
