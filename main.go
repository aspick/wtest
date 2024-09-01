package main

import (
	"net/http"
	"os"

	"github.com/aspick/wtest/src/handler"
)

// doMain はメイン関数の本体です。
func doMain() int {
	// TODO: Graceful shut down を追加する

	mux := http.NewServeMux()
	handler.RegisterHandlers(mux)

	http.ListenAndServe(":8080", mux)

	return 0
}

func main() {
	os.Exit(doMain())
}
