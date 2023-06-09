package app

import "github.com/alanfran/ebay-go-examples/internal/adapters/thingd"

type App struct {
	thingd thingd.API
}

type Params struct {
	Thingd thingd.API
}

func New(p Params) App {
	return App{
		thingd: p.Thingd,
	}
}

// Run brings up the server, blocks this coroutine, and returns any error when it terminates.
func (a *App) Run() error {
	// TODO: graceful shutdown handling

	return nil
}
