package main

import (
	"log"

	"github.com/alanfran/ebay-go-examples/app"
	"github.com/alanfran/ebay-go-examples/internal/adapters/thingd"
	thingdGrpc "github.com/alanfran/ebay-go-examples/z_external/thingd"
)

func main() {
	// construct a grpc client
	grpcClient := thingdGrpc.Client{}

	// construct adapters
	thingdClient := thingd.NewClient(&grpcClient)

	// construct the app
	a := app.New(app.Params{
		Thingd: thingdClient,
	})

	// and run it
	err := a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
