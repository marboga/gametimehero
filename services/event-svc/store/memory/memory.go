// This package implements the store layer using in-memory data store.
// No business logic inside this package, only CRUD operations.
package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/golang/protobuf/ptypes"
	"github.com/pborman/uuid"
	"github.com/sirupsen/logrus"

	eventproto "github.com/marboga/gametimehero/proto/event-svc"
	"github.com/marboga/gametimehero/services/event-svc/store"
)

// Options contains the options to create a memory store
type Options struct {
	Log *logrus.Logger
}

// memory implements store.Store interface.
// Represents the store logic using in-memory data store.
type memory struct {
	sync.Mutex

	data map[string]*eventproto.Event
	log  *logrus.Logger
}

// New is the constructor of memory
func New(opts *Options) store.Store {
	return &memory{
		data: make(map[string]*eventproto.Event),
		log:  opts.Log,
	}
}

// CreateEvent implements store.Store interface.
// This function stores the given event.
func (m *memory) CreateEvent(ctx context.Context, input *eventproto.Event) (*eventproto.Event, error) {
	// Protect the data from race condition and data race.
	m.Lock()
	defer m.Unlock()

	// Generate a new event ID.
	input.Id = uuid.New()

	// Set timestamps
	now := ptypes.TimestampNow()
	input.CreatedAt = now
	input.UpdatedAt = now

	// Store the event
	m.data[input.Id] = input

	return input, nil
}

// ReadEvent implements store.Store interface.
// This function reads an existing event by its ID.
func (m *memory) ReadEvent(ctx context.Context, id string) (*eventproto.Event, error) {
	// Protect the data from race condition and data race.
	m.Lock()
	defer m.Unlock()

	// Retrieve event with the given ID.
	event, ok := m.data[id]
	if !ok {
		// Return the not found errors.
		// Here should be custom not found error implementation
		// to convert in to the proto status instead of return this error.
		return nil, fmt.Errorf("event with ID '%s' doesn't found", id)
	}

	return event, nil
}

// ListEvents implements store.Store interface.
// This function lists all events.
func (m *memory) ListEvents(ctx context.Context) ([]*eventproto.Event, error) {
	// Protect the data from race condition and data race.
	m.Lock()
	defer m.Unlock()

	// Prepare data to return.
	var events []*eventproto.Event
	for _, event := range m.data {
		events = append(events, event)
	}

	return events, nil
}

// UpdateEvent implements store.Store interface.
// This function updates an existing event by its ID.
func (m *memory) UpdateEvent(ctx context.Context, id string, input *eventproto.Event) (*eventproto.Event, error) {
	// Protect the data from race condition and data race.
	m.Lock()
	defer m.Unlock()

	// Retrieve event with the given ID.
	if _, ok := m.data[id]; !ok {
		// Return the not found errors.
		// Here should be custom not found error implementation
		// to convert in to the proto status instead of return this error.
		return nil, fmt.Errorf("event with ID '%s' doesn't found", id)
	}

	// Update event record.
	input.UpdatedAt = ptypes.TimestampNow()
	m.data[id] = input

	return input, nil
}

// DeleteEvent implements store.Store interface.
// This function deletes an existing event by its ID.
func (m *memory) DeleteEvent(ctx context.Context, id string) error {
	// Protect the data from race condition and data race.
	m.Lock()
	defer m.Unlock()

	// Retrieve event with the given ID.
	if _, ok := m.data[id]; !ok {
		// Return the not found errors.
		// Here should be custom not found error implementation
		// to convert in to the proto status instead of return this error.
		return fmt.Errorf("event with ID '%s' doesn't found", id)
	}

	// Delete record.
	delete(m.data, id)

	return nil
}
