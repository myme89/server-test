package api

import (
	"context"
	"net"
	"net/http"
	"server-test/config"
	"server-test/pb"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	pb.UnimplementedDeliverooServer
	Addr    string
	Handler http.Handler
	config  *config.Config
}

func GRPCSever(serverAddr string, config *config.Config) {

	srv := &Server{
		Addr:   serverAddr,
		config: config,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDeliverooServer(grpcServer, srv)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatal("cannot create listener")
	}
	log.Info("Successfully connected")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}

	log.Info("Successfully connected")
}

func GatewaySever(serverAddr string, config *config.Config) {
	srv := &Server{
		Addr:   serverAddr,
		config: config,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	grpcMux := runtime.NewServeMux()
	err := pb.RegisterDeliverooHandlerServer(ctx, grpcMux, srv)

	if err != nil {
		log.Fatal("cannot register handler server", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatal("cannot create listener", err)
	}

	log.Info("Start server at : ", serverAddr)

	err = http.Serve(listener, mux)

	if err != nil {
		log.Fatal("cannot start http gateway server")
	}
}
