package main

import (
	"context"
	protobuf "grpc/grpc_communication_template/proto"
	"io"
	"log"
)

func callComServerStream(client protobuf.ComServiceClient, content *protobuf.ContentList) {
	log.Printf("Streaming has started...")
	stream, err := client.ComServerStreaming(context.Background(), content)

	if err != nil {
		log.Fatalf("could not send content: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streaming: %v", err)
		}
		log.Println(message)
		log.Printf("Streaming finished")
	}
}
