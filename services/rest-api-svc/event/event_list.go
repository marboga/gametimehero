package event

import (
	"net/http"

	"github.com/marboga/gametimehero/services/rest-api-svc/swaggergen/models"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	eventproto "github.com/marboga/gametimehero/proto/event-svc"

	"github.com/marboga/gametimehero/services/rest-api-svc/swaggergen/restapi/operations"
)

// eventsList is the handler of the events listing endpoint.
// This func calls the events listing endpoint of event-svc.
func (h *RestHandler) eventsList(params operations.EventsListParams) middleware.Responder {
	// Call endpoint to list all events.
	resp, err := h.eventService.ListEvents(params.HTTPRequest.Context(), &eventproto.ListEventsRequest{})
	if err != nil {
		// Handle the given error and return 500 status code.
		// Also, write error message into the response.
		// Error handling can be with more clear way, but it's just an example.
		return middleware.ResponderFunc(func(w http.ResponseWriter, _ runtime.Producer) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		})
	} else if resp.GetError().GetCode() != 0 {
		// Handle the given logic error and return 400 status code.
		// Also, write error message into the response.
		// Error handling can be with more clear way, but it's just an example.
		// Here gonna be mapping between RPC and HTTP status codes.
		return middleware.ResponderFunc(func(w http.ResponseWriter, p runtime.Producer) {
			http.Error(w, resp.GetError().GetMessage(), http.StatusBadRequest)
		})
	}

	// Convert proto models to the Swagger models.
	events := make([]*models.Event, len(resp.GetData().GetEvents()))
	for i, event := range resp.GetData().GetEvents() {
		events[i] = toEventModel(event)
	}

	// Return event models.
	return operations.NewEventsListOK().WithPayload(events)
}
