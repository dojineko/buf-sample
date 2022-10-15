package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	bazv1 "github.com/dojineko/buf-sample/namespace-confliction-go/reproduction/baz/protobuf/gen/go"
	foov1 "github.com/dojineko/buf-sample/namespace-confliction-go/reproduction/foo/protobuf/gen/go"
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

	// OK.
	bazv1.RegisterBazServiceServer(grpcServer, &BazServiceServer{})
	// NG, cuz dummy.proto for foo is already registered.
	foov1.RegisterFooServiceServer(grpcServer, &FooServiceServer{})

	fmt.Println("gRPC Server started! Press Ctrl+C to exit.")

	grpcServer.Serve(lis)
}
