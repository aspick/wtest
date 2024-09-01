package handler

import (
	"encoding/json"
	"net/http"

	"github.com/aspick/wtest/src/schema"
	"github.com/aspick/wtest/src/usecase"
	"github.com/jackc/pgx/v5/pgxpool"
)

// CreateInvoiceHandler は,請求書を作成するエンドポイントを実装したハンドラです。
// TODO: Server の struct を作成し、pool などを持たせる
func CreateInvoiceHandler(pool *pgxpool.Pool) func(res http.ResponseWriter, req *http.Request) {
	uc := usecase.NewCreateInvoice(pool)

	return func(res http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		var request schema.CreateInvoiceRequest
		json.NewDecoder(req.Body).Decode(&request)

		// TODO: ログインユーザーの企業をリクエストから取得する
		companyID := 1

		result, err := uc.Execute(ctx, companyID, request)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		resJson, err := json.Marshal(result)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Write(resJson)
	}
}
