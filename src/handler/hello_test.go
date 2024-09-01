package handler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aspick/wtest/src/handler"
)

func TestHelloHandler(t *testing.T) {
	mux := http.NewServeMux()
	handler.RegisterHandlers(mux)

	t.Run("GET /hello", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/hello", nil)
		if err != nil {
			t.Fatal(err)
		}

		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)

		res := rec.Result()

		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, res.StatusCode)
		}

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		if string(resBody) != "Hello, World!" {
			t.Errorf("expected body %q, got %q", "Hello, World!", string(resBody))
		}
	})
}
