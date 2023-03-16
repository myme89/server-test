package api

import (
	"fmt"
	"net"
	"server-test/server-proccess-data/pb_processing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	log "github.com/sirupsen/logrus"
)

type ServerProcessing struct {
	pb_processing.UnimplementedProcessingServer
	Addr string
	// Handler http.Handler
	// config *config.Config
}

// func GRPCSever(serverAddr string, config *config.Config) {
func GRPCSeverStorage(serverAddr string) {
	srv := &ServerProcessing{
		Addr: serverAddr,
		// config: config,
	}

	grpcServer := grpc.NewServer()

	pb_processing.RegisterProcessingServer(grpcServer, srv)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatal("cannot create listener server processing")
	}
	log.Info("Successfully connected server processing")
	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Println(err)
		log.Fatal("cannot start gRPC  server processing")
	}

	log.Info("Successfully connected server stogare")
}
