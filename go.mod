module github.com/marboga/gametimehero

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/go-openapi/errors v0.19.8
	github.com/go-openapi/loads v0.19.6
	github.com/go-openapi/runtime v0.19.24
	github.com/go-openapi/spec v0.19.14
	github.com/go-openapi/strfmt v0.19.11
	github.com/go-openapi/swag v0.19.12
	github.com/go-openapi/validate v0.19.14
	github.com/gogo/protobuf v1.3.0 // indirect
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/jessevdk/go-flags v1.4.0
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/nats/v2 v2.9.1
	github.com/micro/go-plugins/registry/nats/v2 v2.9.1
	github.com/micro/go-plugins/transport/nats/v2 v2.9.1
	github.com/pborman/uuid v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.2.1 // indirect
	github.com/sirupsen/logrus v1.5.0
	github.com/stretchr/testify v1.6.1
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4
	google.golang.org/genproto v0.0.0-20210624195500-8bfb893ecb84
	google.golang.org/grpc v1.38.0 // indirect
	google.golang.org/protobuf v1.27.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
