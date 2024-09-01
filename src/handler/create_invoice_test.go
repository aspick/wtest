package handler_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aspick/wtest/src/config"
	"github.com/aspick/wtest/src/handler"
	"github.com/aspick/wtest/src/schema"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestCreateInvoice(t *testing.T) {
	env, err := config.NewEnv()
	if err != nil {
		t.Fatal(err)
	}

	pool, err := pgxpool.New(context.Background(), env.GetDBURL())
	if err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	handler.RegisterHandlers(mux, pool)

	t.Run("POST /invoice", func(t *testing.T) {
		reqBody := `{
			"customer_id": 1,
			"issue_date": "2021-01-01T00:00:00Z",
			"payment_amount": 100000,
			"due_date": "2021-01-31T00:00:00Z"
		}`
		req, err := http.NewRequest("POST", "/api/invoice", strings.NewReader(reqBody))
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

		var resJson schema.CreateInvoiceResponse
		fmt.Printf("resBody: %s\n", resBody)
		if err := json.Unmarshal(resBody, &resJson); err != nil {
			t.Fatal(err)
		}

		if resJson.InvoiceID == 0 {
			t.Errorf("expected InvoiceID > 0, got %d", resJson.InvoiceID)
		}
	})
}
