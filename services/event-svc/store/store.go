// This package contains the store representation of this service.
// No business logic inside of this package.
package store

import (
	"context"

	eventproto "github.com/marboga/gametimehero/proto/event-svc"
)

// Store represents the behavior of the store layer.
// Currently, proto models are used in the store layer as well.
// The same comment as for controller.Controller interface.
type Store interface {
	// CreateEvent creates a new event by the given input in the store.
	// This function only creates a new record using the given input. No business logic there.
	CreateEvent(context.Context, *eventproto.Event) (*eventproto.Event, error)

	// ReadEvent reads an existing event by its ID from the store.
	ReadEvent(context.Context, string) (*eventproto.Event, error)

	// ListEvents lists all events from the store.
	ListEvents(context.Context) ([]*eventproto.Event, error)

	// UpdateEvent updates an existing event in the store by its ID using the given input.
	// This function only updates the record using the given input. No business logic there.
	UpdateEvent(context.Context, string, *eventproto.Event) (*eventproto.Event, error)

	// DeleteEvent deletes an existing event from the store by its ID.
	// This function only deletes the record using the given input. No business logic there.
	DeleteEvent(context.Context, string) error
}
