package thingd

import (
	"context"

	"github.com/alanfran/ebay-go-examples/z_external/thingd"
)

var _ API = &Client{}

type Client struct {
	external ExternalClient
}

// ExternalClient is the subset of the external interface we use.
// This helps wall off our codebase from ripple effects stemming from changes in external packages.
type ExternalClient interface {
	GetThing(context.Context, *thingd.GetThingRequest, ...thingd.DialOpts) (thingd.Thing, error)
}

func NewClient(externalClient ExternalClient) *Client {
	return &Client{
		external: externalClient,
	}
}
