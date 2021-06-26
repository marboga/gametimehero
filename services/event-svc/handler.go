// This is the transport layer of the service. Here we use RPC communication type.
// You can implement any transport you want, e.g. HTTP, WS, etc.
package eventsvc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	eventproto "github.com/marboga/gametimehero/proto/event-svc"
	"github.com/marboga/gametimehero/proto/health"
	proto "github.com/marboga/gametimehero/proto/status"
	"github.com/marboga/gametimehero/services/event-svc/controller"
	"github.com/marboga/gametimehero/utils/rpc"
)

// To make sure Handler implements eventproto.EventService interface.
var _ eventproto.EventServiceHandler = &Handler{}

// Options serves as the dependency injection container to create a new handler.
type Options struct {
	Service        controller.Controller
	SelfPingClient *health.SelfPingClient
	Log            *logrus.Logger
}

// Handler implements authzproto.AuthorizationServiceHandler interface.
type Handler struct {
	service        controller.Controller
	selfPingClient *health.SelfPingClient
	log            *logrus.Logger
}

// NewHandler returns a new handler for the event-svc.
func NewHandler(opts *Options) *Handler {
	return &Handler{
		service:        opts.Service,
		selfPingClient: opts.SelfPingClient,
		log:            opts.Log,
	}
}

// CreateEvent implements eventproto.EventServiceHandler interface.
// Calls the service's method to create a new event by the given input.
func (h *Handler) CreateEvent(ctx context.Context, req *eventproto.CreateEventRequest, resp *eventproto.CreateEventResponse) error {
	// Create event.
	createdEvent, err := h.service.CreateEvent(ctx, req.GetEvent())
	if err != nil {
		// Try to convert the given error to the proto status.
		if resStatus, ok := h.errorAsStatus(ctx, err); ok {
			resp.Result = &eventproto.CreateEventResponse_Error{
				Error: resStatus,
			}
			return nil
		}

		// Otherwise just return this error wrapped to a description.
		return errors.Wrap(err, "unable to create event")
	}

	// Prepare RPC response data.
	resp.Result = &eventproto.CreateEventResponse_Event{
		Event: createdEvent,
	}
	return nil
}

// ReadEvent implements eventproto.EventServiceHandler interface.
// Calls the service's method to read an existing event by the given ID.
func (h *Handler) ReadEvent(ctx context.Context, req *eventproto.ReadEventRequest, resp *eventproto.ReadEventResponse) error {
	// Read event.
	event, err := h.service.ReadEvent(ctx, req.GetEventId())
	if err != nil {
		// Try to convert the given error to the proto status.
		if resStatus, ok := h.errorAsStatus(ctx, err); ok {
			resp.Result = &eventproto.ReadEventResponse_Error{
				Error: resStatus,
			}
			return nil
		}

		// Otherwise just return this error wrapped to a description.
		return errors.Wrapf(err, "unable to read event with ID '%s'", req.GetEventId())
	}

	// Prepare RPC response data.
	resp.Result = &eventproto.ReadEventResponse_Event{
		Event: event,
	}
	return nil
}

// ListEvents implements eventproto.EventServiceHandler interface.
// Calls the service's method to list all events.
func (h *Handler) ListEvents(ctx context.Context, req *eventproto.ListEventsRequest, resp *eventproto.ListEventsResponse) error {
	// List all events.
	events, err := h.service.ListEvents(ctx)
	if err != nil {
		// Try to convert the given error to the proto status.
		if resStatus, ok := h.errorAsStatus(ctx, err); ok {
			resp.Result = &eventproto.ListEventsResponse_Error{
				Error: resStatus,
			}
			return nil
		}

		// Otherwise just return this error wrapped to a description.
		return errors.Wrapf(err, "unable to list events")
	}

	// Prepare RPC response data.
	resp.Result = &eventproto.ListEventsResponse_Data{
		Data: &eventproto.ListEventsResponseOK{
			Events: events,
		},
	}
	return nil
}

// UpdateEvent implements eventproto.EventServiceHandler interface.
// Calls the service's method to update an existing event by the given ID and input.
func (h *Handler) UpdateEvent(ctx context.Context, req *eventproto.UpdateEventRequest, resp *eventproto.UpdateEventResponse) error {
	// Update event by its ID.
	event, err := h.service.UpdateEvent(ctx, req.GetEventId(), req.GetEvent())
	if err != nil {
		// Try to convert the given error to the proto status.
		if resStatus, ok := h.errorAsStatus(ctx, err); ok {
			resp.Result = &eventproto.UpdateEventResponse_Error{
				Error: resStatus,
			}
			return nil
		}

		// Otherwise just return this error wrapped to a description.
		return errors.Wrapf(err, "unable to update event with ID '%s'", req.GetEventId())
	}

	// Prepare RPC response data.
	resp.Result = &eventproto.UpdateEventResponse_Event{
		Event: event,
	}
	return nil
}

// DeleteEvent implements eventproto.EventServiceHandler interface.
// Calls the service's method to delete an existing event by the given ID.
func (h *Handler) DeleteEvent(ctx context.Context, req *eventproto.DeleteEventRequest, resp *eventproto.DeleteEventResponse) error {
	// Delete event by its ID.
	if err := h.service.DeleteEvent(ctx, req.GetEventId()); err != nil {
		// Try to convert the given error to the proto status.
		if resStatus, ok := h.errorAsStatus(ctx, err); ok {
			resp.Result = &eventproto.DeleteEventResponse_Error{
				Error: resStatus,
			}
			return nil
		}

		// Otherwise just return this error wrapped to a description.
		return errors.Wrapf(err, "unable to delete event with ID '%s'", req.GetEventId())
	}

	// Prepare RPC response data.
	resp.Result = &eventproto.DeleteEventResponse_Empty{
		Empty: &empty.Empty{},
	}
	return nil
}

// Health implements eventproto.EventServiceHandler interface
func (h *Handler) Health(ctx context.Context, _ *empty.Empty, res *health.HealthResponse) error {
	// Check database
	if err := h.service.HealthCheck(); err != nil {
		res.Status = health.HealthResponse_NOT_SERVING
		return err
	}

	// Check nats, i.e. call ourselves (exactly this node) via nats
	if err := h.selfPingClient.Ping(ctx); err != nil {
		res.Status = health.HealthResponse_NOT_SERVING
		return errors.Wrapf(err, "unable to ping ourselves")
	}

	res.Status = health.HealthResponse_SERVING
	return nil
}

// Ping implements eventproto.EventServiceHandler and helath.Pinger interface.
// This is needed to implement self-pinger functionality.
func (h *Handler) Ping(ctx context.Context, _ *empty.Empty, _ *empty.Empty) error {
	return nil
}

// errorAsStatus converts the given error to the proto status.
// This function have to be implemented according to the logic of your project.
// For now, it always returns the ErrAborted RPC status code.
// What will be returned:
// - the first parameter if the proto status of the error;
// - the second boolean value is true, if the error has been matched with one of RPC statuses;
func (h *Handler) errorAsStatus(ctx context.Context, err error) (*proto.Status, bool) {
	return rpc.ErrAbortedf(err.Error()), true
}
