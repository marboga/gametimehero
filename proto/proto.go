package proto

//go:generate protoc --proto_path=$GOPATH/src:. --micro_out=$GOPATH/src --go_out=$GOPATH/src $GOPATH/src/github.com/marboga/gametimehero/proto/status/status.proto
//go:generate protoc --proto_path=$GOPATH/src:. --micro_out=$GOPATH/src --go_out=$GOPATH/src $GOPATH/src/github.com/marboga/gametimehero/proto/health/health.proto
//go:generate protoc --proto_path=$GOPATH/src:. --micro_out=$GOPATH/src --go_out=$GOPATH/src $GOPATH/src/github.com/marboga/gametimehero/proto/common/error_response.proto
//go:generate protoc --proto_path=$GOPATH/src:. --micro_out=$GOPATH/src --go_out=$GOPATH/src $GOPATH/src/github.com/marboga/gametimehero/proto/common/types.proto

//go:generate protoc --proto_path=$GOPATH/src:. --micro_out=$GOPATH/src --go_out=$GOPATH/src $GOPATH/src/github.com/marboga/gametimehero/proto/account-svc/account.proto
