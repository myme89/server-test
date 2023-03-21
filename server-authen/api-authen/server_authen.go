package api

import (
	"fmt"
	"net"
	"server-test/server-authen/pb_authen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	log "github.com/sirupsen/logrus"
)

type ServerAuthen struct {
	pb_authen.UnimplementedAuthenServer
	Addr string
	// Handler http.Handler
	// config *config.Config
	// client clients.StorageClient
}

// func GRPCSever(serverAddr string, config *config.Config) {
func GRPCSeverAuthen(serverAddr string) {
	srv := &ServerAuthen{
		Addr: serverAddr,
		// config: config,
	}
	log.Info("trongnhat")
	grpcServer := grpc.NewServer()

	pb_authen.RegisterAuthenServer(grpcServer, srv)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatal("cannot create listener server authen")
	}
	log.Info("Successfully connected server authen")
	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Println(err)
		log.Fatal("cannot start gRPC  server authen")
	}

	log.Info("Successfully connected server authen")
}
