package main

import (
	"context"
	protobuf "grpc/grpc_communication_template/proto"
	"io"
	"log"
	"time"
)

func callComBidirectionalStream(client protobuf.ComServiceClient, messages *protobuf.MessagesList) {
	log.Printf("Bidirectional Streaming Started \xF0\x9F\x93\xA1")
	stream, err := client.ComBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send messages: %v", err)

	}

	// Create a channel to send close call to the GO routine.
	waitc := make(chan struct{})

	// Go routine to be able to send and receive.
	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, message := range messages.Messages {
		request := &protobuf.ComRequest{
			Message: message,
		}
		if err := stream.Send(request); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		time.Sleep(1 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional streaming finished")
}
