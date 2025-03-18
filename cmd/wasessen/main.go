package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/davidtemelkov/wasessen/internal/api"
	"github.com/davidtemelkov/wasessen/internal/data"
)

func main() {
	// TODO: Refactor env loading for local dev and for deployed app
	// err := godotenv.Load("../../.env")
	// if err != nil {
	// 	os.Exit(1)
	// }

	var err error
	var ctx = context.Background()
	data.Db, err = data.NewDynamoDbClient(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r := api.SetUpRoutes()
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
