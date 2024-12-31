package http

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goblinus/httplib/v2/buildmeta"
	"github.com/stretchr/testify/assert"
)

type wantValues struct {
	HttpCode int
	Content  string
}

func setup() *HTTPRouter {
	meta := buildmeta.NewBuildMeta("v0.0.1", "release_test", "test", "2006-01-02 15:04:05")
	router := NewHTTPRouter()
	router.Init(WithDefaultHandlers(meta))

	return router
}

func TestHandlers(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		descr   string
		request *http.Request
		want    *wantValues
	}{
		{
			"ready request",
			httptest.NewRequest(http.MethodGet, "/ready", nil),
			&wantValues{
				HttpCode: 200,
				Content:  "{\"state\":\"OK\"}",
			},
		},
		{
			"health request",
			httptest.NewRequest(http.MethodGet, "/health", nil),
			&wantValues{
				HttpCode: 200,
				Content:  "{\"builder\":\"test\",\"dateTime\":\"2006-01-02 15:04:05\",\"release\":\"release_test\",\"version\":\"v0.0.1\"}",
			},
		},
	}

	for _, tt := range testCases {
		want := tt.want
		request := tt.request
		t.Run(tt.descr, func(t *testing.T) {
			t.Parallel()

			router := setup()
			w := httptest.NewRecorder()
			router.Routes().ServeHTTP(w, request)
			data, err := io.ReadAll(w.Result().Body)
			if err != nil {
				t.Error("body should be readable")
			}

			defer w.Result().Body.Close()

			assert.Equal(t, want.HttpCode, w.Code, "status should be OK (status=200)")
			assert.Equal(t, want.Content, string(data), "content should be equal")
		})
	}
}
