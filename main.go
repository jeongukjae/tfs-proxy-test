package main

import (
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	tfs_api_pb "github.com/tensorflow/serving/tensorflow_serving/apis"
)

var (
	host   = flag.String("host", ":8000", "server host")
	server = flag.String("target", ":8500", "target server")
)

func main() {
	lis, err := net.Listen("tcp", *host)
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer(grpc.UnknownServiceHandler(handleTfsRequest))
	log.Println("Serve gRPC server @", *host)
	srv.Serve(lis)
}

func handleTfsRequest(srv interface{}, stream grpc.ServerStream) error {
	fullMethodName, ok := grpc.MethodFromServerStream(stream)
	if !ok {
		return status.Errorf(codes.Internal, "lowLevelServerStream not exists in context")
	}
	log.Println(fullMethodName)

	var requestPb tfs_api_pb.PredictRequest
	err := stream.RecvMsg(&requestPb)
	if err != nil {
		log.Println(err)
	}
	log.Println(srv, requestPb)
	return nil
}
