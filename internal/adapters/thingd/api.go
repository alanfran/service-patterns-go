package thingd

import (
	"context"

	"github.com/google/uuid"
)

// API defines our internal API to interact with an external service.
type API interface {
	GetThing(context.Context, GetThingRequest) (Thing, error)
}

// Thing is an internal view of some external type.
type Thing struct {
	ID   uuid.UUID
	Name string
}

type GetThingRequest struct {
	ID uuid.UUID
}
