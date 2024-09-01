package handler

import "net/http"

// HelloHandler は,テスト用のエンドポイントを実装したハンドラです。
func HelloHandler(res http.ResponseWriter, _ *http.Request) {
	res.Write([]byte("Hello, World!"))
}
