package api

import (
	"fmt"
	"net"
	"server-test/server-storage/clients"
	"server-test/server-storage/pb_storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	log "github.com/sirupsen/logrus"
)

type ServerStorage struct {
	pb_storage.UnimplementedStogareServer
	Addr string
	// Handler http.Handler
	// config *config.Config
	clientTest     clients.StorageClient
	clientDatabase clients.DatabaseClient
}

// func GRPCSever(serverAddr string, config *config.Config) {
func GRPCSeverStorage(serverAddr string) {
	srv := &ServerStorage{
		Addr: serverAddr,
		// config: config,
	}

	grpcServer := grpc.NewServer()

	pb_storage.RegisterStogareServer(grpcServer, srv)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatal("cannot create listener server storage")
	}
	log.Info("Successfully connected server stogare")
	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Println(err)
		log.Fatal("cannot start gRPC  server stogare")
	}

	log.Info("Successfully connected server stogare")
}
