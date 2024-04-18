package main

import (
	"log"
	"time"

	protobuf "grpc/grpc_communication_template/proto"
)

var treat = [5]string{
	"\xF0\x9F\x8C\x8B",
	"\xF0\x9F\x94\xA5",
	"\xF0\x9F\x8E\xB2",
	"\xF0\x9F\x8E\xA5",
	"\xF0\x9F\x8E\xA0",
}

func (server *comServer) ComServerStreaming(request *protobuf.ContentList, stream protobuf.ComService_ComServerStreamingServer) error {

	log.Printf("got request with content: %v", request.Content)
	for i, content := range request.Content {
		response := &protobuf.ComResponse{
			Message: "Your words: " + content + treat[i%5],
		}
		if err := stream.Send(response); err != nil {
			return err
		}
		time.Sleep(2 * time.Second) // Time delay to mock computation
	}
	return nil

}
