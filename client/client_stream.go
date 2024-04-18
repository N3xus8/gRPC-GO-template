package main

import (
	"context"
	protobuf "grpc/grpc_communication_template/proto"
	"log"
	"time"
)

func callComClientStream(client protobuf.ComServiceClient, messages *protobuf.MessagesList) {

	log.Printf("Client streaming started")
	stream, err := client.ComClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send messages: %v", err)
	}

	for _, message := range messages.Messages {
		request := &protobuf.ComRequest{
			Message: message,
		}
		if err := stream.Send(request); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("Sent the request with message: %s", message)
		time.Sleep(2 * time.Second)

	}

	response, err := stream.CloseAndRecv()
	log.Printf("Client streaming finished.")
	if err != nil {
		log.Fatalf("Error while receiving %v", err)
	}
	log.Printf("%v", response.Messages)

}
