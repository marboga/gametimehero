// This package contains business logic of the service.
// No store or transport logic inside.
package controller

import (
	"context"

	eventproto "github.com/marboga/gametimehero/proto/event-svc"
)

// Controller represents the behavior of the business/controller logic of the service.
// Currently, proto models are used in the controller layer.
// This is not really the right way because proto models are more about the transport layer, not the controller one.
// So if you would like to implement domain/controller representation of proto models, you can do so.
// Don't forget to implement a conversion logic between proto models and controller ones :)
type Controller interface {
	// HealthCheck returns an error if there is a problem with the service.
	HealthCheck() error

	// CreateEvent creates a new Event by the given input.
	CreateEvent(context.Context, *eventproto.Event) (*eventproto.Event, error)

	// ReadEvent reads an existing Event by its ID.
	ReadEvent(context.Context, string) (*eventproto.Event, error)

	// ListEvents lists all events.
	ListEvents(context.Context) ([]*eventproto.Event, error)

	// UpdateEvent updates an existing Event by its ID using the given input.
	UpdateEvent(context.Context, string, *eventproto.Event) (*eventproto.Event, error)

	// DeleteEvent deletes an existing Event by its ID.
	DeleteEvent(context.Context, string) error
}
