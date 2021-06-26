package event

import (
	"github.com/sirupsen/logrus"

	eventproto "github.com/marboga/gametimehero/proto/event-svc"
	"github.com/marboga/gametimehero/services/rest-api-svc/swaggergen/restapi/operations"
)

// RestHandlerOptions contains required options for the handler.
// This handler implements REST endpoints with handling incoming data.
type RestHandlerOptions struct {
	EventService eventproto.EventService
	Logger       logrus.FieldLogger
}

// RestHandler defines the REST interface for the business service.
type RestHandler struct {
	eventService eventproto.EventService
	logger       logrus.FieldLogger
}

// NewRestHandler creates a new Handler.
func NewRestHandler(opts *RestHandlerOptions) *RestHandler {
	return &RestHandler{
		eventService: opts.EventService,
		logger:       opts.Logger,
	}
}

// Register registers endpoints to the handler.
func (h *RestHandler) Register(api *operations.RestAPISvcAPI) {
	api.EventCreateHandler = operations.EventCreateHandlerFunc(h.eventCreate)
	api.EventReadHandler = operations.EventReadHandlerFunc(h.eventRead)
	api.EventsListHandler = operations.EventsListHandlerFunc(h.eventsList)
	api.EventUpdateHandler = operations.EventUpdateHandlerFunc(h.eventUpdate)
	api.EventDeleteHandler = operations.EventDeleteHandlerFunc(h.eventDelete)
}
