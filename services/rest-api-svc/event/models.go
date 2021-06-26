package event

import (
	"github.com/go-openapi/strfmt"
	"github.com/golang/protobuf/ptypes"

	eventproto "github.com/marboga/gametimehero/proto/event-svc"
	"github.com/marboga/gametimehero/services/rest-api-svc/swaggergen/models"
)

// toEventModel converts the event proto model to the Swagger model.
func toEventModel(u *eventproto.Event) *models.Event {
	updatedAt, _ := ptypes.Timestamp(u.GetUpdatedAt())
	createdAt, _ := ptypes.Timestamp(u.GetCreatedAt())

	return &models.Event{
		ID:        u.GetId(),
		Name:      u.GetName(),
		UpdatedAt: strfmt.DateTime(updatedAt),
		CreatedAt: strfmt.DateTime(createdAt),
	}
}
