package main

import (
	"context"
	protobuf "grpc/grpc_communication_template/proto"
	"log"
	"time"
)

func callHealthz(client protobuf.ComServiceClient) {
	context, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.Healthz(context, &protobuf.NoParam{})
	if err != nil {
		log.Fatalf("could not check : %v", err)
	}
	log.Printf("%s", response.Message)
}
