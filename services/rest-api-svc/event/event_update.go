package event

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	eventproto "github.com/marboga/gametimehero/proto/event-svc"
	"github.com/marboga/gametimehero/services/rest-api-svc/swaggergen/restapi/operations"
)

// eventUpdate is the handler of the event updating endpoint.
// This func calls the event updating endpoint of event-svc with the given data.
func (h *RestHandler) eventUpdate(params operations.EventUpdateParams) middleware.Responder {
	// Call endpoint to update an existing event with the given input.
	resp, err := h.eventService.UpdateEvent(params.HTTPRequest.Context(), &eventproto.UpdateEventRequest{
		EventId: params.EventID.String(),
		Event: &eventproto.Event{
			Name: params.Seed.Name,
		},
	})
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

	// Convert proto model to the Swagger model.
	model := toEventModel(resp.GetEvent())

	// Return the updated event model.
	return operations.NewEventUpdateOK().WithPayload(model)
}
