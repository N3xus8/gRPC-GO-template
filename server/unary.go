package main

import (
	"context"
	protobuf "grpc/grpc_communication_template/proto"
)

func (server *comServer) Healthz(context context.Context, request *protobuf.NoParam) (*protobuf.StatusResponse, error) {
	return &protobuf.StatusResponse{
		Message: "Server up and running \xF0\x9F\x91\xBE",
	}, nil
}
