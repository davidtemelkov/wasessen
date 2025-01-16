package main

import (
	"context"
	"net/http"
	"os"

	"github.com/davidtemelkov/wasessen/api"
	"github.com/davidtemelkov/wasessen/data"
)

func main() {
	var ctx = context.Background()

	var err error
	data.Db, err = data.NewDynamoDbClient(ctx)
	if err != nil {
		os.Exit(1)
	}

	r := api.SetUpRoutes()
	if err := http.ListenAndServe(":8080", r); err != nil {
		os.Exit(1)
	}
}
