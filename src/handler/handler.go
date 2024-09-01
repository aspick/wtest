package handler

import (
	"net/http"

	"github.com/aspick/wtest/src/server"
)

// HelloHandler は,サーバーにエンドポイントを追加するためのハンドラです。
func RegisterHandlers(mux *http.ServeMux) {
	middlewares := []func(http.HandlerFunc) http.HandlerFunc{
		server.RecoverMiddleware,
		server.LoggingMiddleware,
	}

	mux.HandleFunc("GET /hello", server.RegisterMiddlewares(middlewares, HelloHandler))
}
