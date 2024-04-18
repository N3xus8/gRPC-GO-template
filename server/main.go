//https://www.youtube.com/watch?v=a6G5-LUlFO4

package main

import (
	"log"
	"net"

	protobuf "grpc/grpc_communication_template/proto"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type comServer struct {
	protobuf.ComServiceServer
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	// New grpc server
	grpcServer := grpc.NewServer()

	//Register the server with service server
	protobuf.RegisterComServiceServer(grpcServer, &comServer{})

	log.Printf("server started at %v  \xF0\x9F\x93\xA1", listen.Addr())

	// Start gRPC server.
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
}
