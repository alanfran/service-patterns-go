package device

import "github.com/alanfran/ebay-go-examples/internal/models/device/status"

type Model struct {
	Name   string
	Status status.Enum
}
