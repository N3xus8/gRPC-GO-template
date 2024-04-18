package main

import (
	protobuf "grpc/grpc_communication_template/proto"
	"io"
	"log"
)

func (service *comServer) ComClientStreaming(stream protobuf.ComService_ComClientStreamingServer) error {
	var messages []string

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&protobuf.MessagesList{Messages: messages})
		}

		if err != nil {
			return err
		}
		log.Printf("Got request with name: %v", request.Message)
		messages = append(messages, "Acknowledgment: Roger that! \xF0\x9F\x92\xAF", request.Message)
	}
}
