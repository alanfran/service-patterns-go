// Package thingd is an example of an external service client,
// in this case following a grpc-like interface.
package thingd

import (
	"context"

	"github.com/google/uuid"
)

var _ GeneratedInterface = &Client{}

type Client struct{}

type GeneratedInterface interface {
	GetThing(context.Context, *GetThingRequest, ...DialOpts) (Thing, error)
	GetSomethingElse(context.Context, ...DialOpts) error
}

type Thing struct {
	ID   string
	Name string
}

type GetThingRequest struct {
	ID string
}

type DialOpts struct{}

// GetThing is a generated grpc client method.
func (c *Client) GetThing(ctx context.Context, req *GetThingRequest, opts ...DialOpts) (Thing, error) {
	// In this example, every Thing has the name "Carl".
	return Thing{
		ID:   uuid.New().String(),
		Name: "Carl",
	}, nil
}

// GetSomethingElse is a generated grpc client method. This is not used by our application,
// and its existence can be ignored by using an "external interface" pattern in an adapter.
func (c *Client) GetSomethingElse(ctx context.Context, opts ...DialOpts) error {
	return nil
}
