package handler

import (
	"net/http"

	"github.com/aspick/wtest/src/server"
	"github.com/jackc/pgx/v5/pgxpool"
)

// HelloHandler は,サーバーにエンドポイントを追加するためのハンドラです。
func RegisterHandlers(mux *http.ServeMux, pool *pgxpool.Pool) {
	middlewares := []func(http.HandlerFunc) http.HandlerFunc{
		server.RecoverMiddleware,
		server.LoggingMiddleware,
		// TODO: 認証・認可 に関する middleware を追加する
	}

	mux.HandleFunc("GET /hello", server.RegisterMiddlewares(middlewares, HelloHandler))
	mux.HandleFunc("POST /api/invoice", server.RegisterMiddlewares(middlewares, CreateInvoiceHandler(pool)))
}
