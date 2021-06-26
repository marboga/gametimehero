package event

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	eventproto "github.com/marboga/gametimehero/proto/event-svc"
	"github.com/marboga/gametimehero/services/rest-api-svc/swaggergen/restapi/operations"
)

// eventDelete is the handler of the event deletion endpoint.
// This func calls the event deletion endpoint of event-svc with the given data.
func (h *RestHandler) eventDelete(params operations.EventDeleteParams) middleware.Responder {
	// Call endpoint to delete an existing event with the given ID.
	resp, err := h.eventService.DeleteEvent(params.HTTPRequest.Context(), &eventproto.DeleteEventRequest{
		EventId: params.EventID.String(),
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

	// Return nothing, just 204 status code.
	return operations.NewEventDeleteNoContent()
}
