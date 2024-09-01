package server

import (
	"fmt"
	"net/http"
)

// RegisterMiddlewares は複数のミドルウェアを登録する関数です。
func RegisterMiddlewares(middlewares []func(http.HandlerFunc) http.HandlerFunc, handler http.HandlerFunc) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

// RecoverMiddleware は panic を回復してエラーレスポンスを返すミドルウェアです。
func RecoverMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(res, fmt.Sprintf("%v", err), http.StatusInternalServerError)
			}
		}()
		next(res, req)
	}
}

// LoggingMiddleware はリクエストをログ出力するミドルウェアです。
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Printf("%s %s\n", req.Method, req.URL)
		next(res, req)
	}
}
