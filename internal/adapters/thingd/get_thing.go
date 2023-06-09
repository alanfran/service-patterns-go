package thingd

import (
	"context"

	"github.com/alanfran/ebay-go-examples/z_external/thingd"
	"github.com/google/uuid"
)

func (c *Client) GetThing(ctx context.Context, req GetThingRequest) (Thing, error) {
	externalThing, err := c.external.GetThing(ctx, &thingd.GetThingRequest{
		ID: req.ID.String(),
	})
	if err != nil {
		return Thing{}, err
	}

	id, err := uuid.Parse(externalThing.ID)
	if err != nil {
		return Thing{}, err
	}

	thing := Thing{
		ID:   id,
		Name: externalThing.Name,
	}

	return thing, nil
}
