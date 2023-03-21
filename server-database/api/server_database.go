package api

import (
	"fmt"
	"net"
	"server-test/server-database/config"
	"server-test/server-database/pb_database"
	"server-test/server-storage/clients"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	log "github.com/sirupsen/logrus"
)

type ServerDatabase struct {
	pb_database.UnimplementedDatabaseServer
	Addr string
	// Handler http.Handler
	config *config.Config
	client clients.StorageClient
}

// func GRPCSever(serverAddr string, config *config.Config) {
func GRPCSeverDatabase(serverAddr string, config *config.Config) {
	srv := &ServerDatabase{
		Addr:   serverAddr,
		config: config,
	}

	grpcServer := grpc.NewServer()

	pb_database.RegisterDatabaseServer(grpcServer, srv)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatal("cannot create listener server database")
	}
	log.Info("Successfully connected server database")
	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Println(err)
		log.Fatal("cannot start gRPC  server database")
	}

	log.Info("Successfully connected server database")
}
