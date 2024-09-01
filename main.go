package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aspick/wtest/src/config"
	"github.com/aspick/wtest/src/handler"
	"github.com/jackc/pgx/v5/pgxpool"
)

// doMain はメイン関数の本体です。
func doMain() int {
	// TODO: Graceful shut down を追加する

	env := config.Env{}

	pool, err := setupDB(env)
	if err != nil {
		fmt.Printf("failed to setup db: %v", err)
		return 1
	}

	mux := http.NewServeMux()
	handler.RegisterHandlers(mux, pool)

	http.ListenAndServe(":8080", mux)

	return 0
}

func setupDB(env config.Env) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), env.GetDBURL())
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func main() {
	os.Exit(doMain())
}
