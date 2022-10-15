package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	bazv1 "github.com/dojineko/buf-sample/namespace-confliction-go/fixed/baz/protobuf/gen/go/baz/v1"
	foov1 "github.com/dojineko/buf-sample/namespace-confliction-go/fixed/foo/protobuf/gen/go/foo/v1"
	"google.golang.org/grpc"
)

type BazServiceServer struct {
	bazv1.UnimplementedBazServiceServer
}
type FooServiceServer struct {
	foov1.UnimplementedFooServiceServer
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	// All OK :)
	bazv1.RegisterBazServiceServer(grpcServer, &BazServiceServer{})
	foov1.RegisterFooServiceServer(grpcServer, &FooServiceServer{})

	fmt.Println("gRPC Server started! Press Ctrl+C to exit.")

	grpcServer.Serve(lis)
}
