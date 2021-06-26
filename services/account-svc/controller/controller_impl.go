package controller

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	accountproto "github.com/marboga/gametimehero/proto/account-svc"
	"github.com/marboga/gametimehero/services/account-svc/store"
)

// Options contains options to create a controller.
type Options struct {
	Store store.Store
	Log   *logrus.Logger
}

// controller implements the business/controller logic of the service.
type controller struct {
	store store.Store
	log   *logrus.Logger
}

// New is the constructor of controller.
func New(opts *Options) Controller {
	return &controller{
		store: opts.Store,
		log:   opts.Log,
	}
}

// HealthCheck implements Controller interface.
func (d *controller) HealthCheck() error {
	return nil
}

// CreateUser implements Controller interface.
// The business logic of the user creation operation can be implemented within this function.
// For now, it's not implemented because this is just an example of an architecture.
func (d *controller) CreateUser(ctx context.Context, input *accountproto.User) (*accountproto.User, error) {
	// Call the store directly.
	createdUser, err := d.store.CreateUser(ctx, input)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create user in the store layer")
	}

	return createdUser, nil
}

// ReadUser implements Controller interface.
// The business logic of the user reading operation can be implemented within this function.
// For now, it's not implemented because this is just an example of an architecture.
func (d *controller) ReadUser(ctx context.Context, id string) (*accountproto.User, error) {
	// Call the store directly.
	user, err := d.store.ReadUser(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read user in the store layer with ID '%s'", id)
	}

	return user, nil
}

// ListUsers implements Controller interface.
// The business logic of the user listing operation can be implemented within this function.
// For now, it's not implemented because this is just an example of an architecture.
func (d *controller) ListUsers(ctx context.Context) ([]*accountproto.User, error) {
	// Call the store directly.
	users, err := d.store.ListUsers(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list users in the store layer")
	}

	return users, nil
}

// UpdateUser implements Controller interface.
// The business logic of the user updating operation can be implemented within this function.
// For now, it's not implemented because this is just an example of an architecture.
func (d *controller) UpdateUser(ctx context.Context, id string, input *accountproto.User) (*accountproto.User, error) {
	// Call the store directly.
	updatedUser, err := d.store.UpdateUser(ctx, id, input)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to update user in the store layer with ID '%s'", id)
	}

	return updatedUser, nil
}

// DeleteUser implements Controller interface.
// The business logic of the user deletion operation can be implemented within this function.
// For now, it's not implemented because this is just an example of an architecture.
func (d *controller) DeleteUser(ctx context.Context, id string) error {
	// Call the store directly.
	if err := d.store.DeleteUser(ctx, id); err != nil {
		return errors.Wrapf(err, "unable to delete user in the store layer with ID '%s'", id)
	}

	return nil
}
