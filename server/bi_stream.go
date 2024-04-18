package main

import (
	protobuf "grpc/grpc_communication_template/proto"
	"io"
	"log"
)

func (server *comServer) ComBidirectionalStreaming(stream protobuf.ComService_ComBidirectionalStreamingServer) error {

	// Note: Server side handling is kind of sequential?
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with message: %v", request.Message)
		response := &protobuf.ComResponse{
			Message: "Server side bi-dir mode:" + request.Message,
		}
		if err := stream.Send(response); err != nil {
			return err
		}
	}
}
