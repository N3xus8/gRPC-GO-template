package main

import (
	"log"

	protobuf "grpc/grpc_communication_template/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {

	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials())) // Using insecure but need to change it to https . Note by default gRPC uses HTTP2 with uses TLS.
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := protobuf.NewComServiceClient(conn)

	content := &protobuf.ContentList{
		Content: []string{"Ubik", "Tears", "Rain", "Resistance", "Futile"},
	}

	messageList := &protobuf.MessagesList{
		Messages: []string{"Replicants", "Unicorn", "More", "Human...", "...than human"},
	}

	callHealthz(client)

	callComServerStream(client, content)

	callComClientStream(client, messageList)

	callComBidirectionalStream(client, messageList)

}
