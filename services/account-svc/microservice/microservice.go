package microservice

import (
	"github.com/micro/go-micro/v2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	accountproto "github.com/marboga/gametimehero/proto/account-svc"
	"github.com/marboga/gametimehero/proto/health"
	accountsvc "github.com/marboga/gametimehero/services/account-svc"
	"github.com/marboga/gametimehero/services/account-svc/controller"
	"github.com/marboga/gametimehero/services/account-svc/store/memory"
	"github.com/marboga/gametimehero/utils/healthchecker"
	"github.com/marboga/gametimehero/utils/rpc"
)

// MicroService is the micro-service.
type MicroService struct {
	svc     micro.Service
	handler *accountsvc.Handler
	log     *logrus.Logger
}

// Init initializes the service.
func Init(clientOpts *ClientOptions) (*MicroService, error) {
	// Create micro-service.
	svc := micro.NewService(
		micro.Name(rpc.AccountServiceName),
		micro.Version(clientOpts.Version),
		micro.Flags(flags...),
		micro.BeforeStart(func() error {
			return opts.Validate()
		}),
	)

	// Parse command-line arguments.
	svc.Init()

	return New(svc, clientOpts)
}

// New is the constructor of the service.
func New(svc micro.Service, clientOpts *ClientOptions) (*MicroService, error) {
	// Create a self-pinger client.
	selfPingClient := health.NewSelfPingClient(svc, accountproto.NewAccountService(rpc.AccountServiceName, svc.Client()))

	// Create store layer using in-memory data store.
	// Here can be any implementation of the store layer.
	store := memory.New(&memory.Options{
		Log: clientOpts.Log,
	})

	// Create business layer.
	service := controller.New(&controller.Options{
		Store: store,
		Log:   clientOpts.Log,
	})

	// Create RPC handler.
	handler := accountsvc.NewHandler(&accountsvc.Options{
		Service:        service,
		SelfPingClient: selfPingClient,
		Log:            clientOpts.Log,
	})

	// Register the service.
	if err := accountproto.RegisterAccountServiceHandler(svc.Server(), handler); err != nil {
		return nil, errors.Wrap(err, "failed to register handler")
	}

	return &MicroService{
		svc:     svc,
		handler: handler,
		log:     clientOpts.Log,
	}, nil
}

// Run runs the service.
func (s *MicroService) Run() error {
	if opts.IsTest {
		s.log.Info("Running in test mode!")
	}

	// Run helathcheck endpoint.
	shutdown := healthchecker.Run(s.log, healthchecker.WrapRPC(s.handler.Health), nil)

	// Stop helathcheck endpoint after RPC service stop.
	s.svc.Init(micro.AfterStop(shutdown))

	// Start service.
	if err := s.svc.Run(); err != nil {
		return errors.Wrap(err, "failed to run")
	}

	return nil
}
