package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"server-test/server-gateway/clients"
	"server-test/server-gateway/config"
	"server-test/server-gateway/pb"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	pb.UnimplementedDeliverooServer
	Addr string
	// Handler http.Handler
	config        *config.Config
	clientStogare clients.StorageClient
	clientAuthen  clients.AuthenClient
	// clientDatabase   clients.DatabaseClient
	clientProcessing clients.ProcessingClient
}

func GRPCSever(serverAddr string) {

	srv := &Server{
		Addr: serverAddr,
		// config: config,
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

	// err = RegisterDeliverooHandlerServerCustom(ctx, grpcMux, srv)

	grpcMux.HandlePath("POST", "/v1/upload-file", func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		srv.ImportDataWithHttp(w, req)
	})

	grpcMux.HandlePath("GET", "/v1/download-link", func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		srv.DowloadLinkWithHttp(w, req)
	})

	grpcMux.HandlePath("GET", "/v1/export-template", func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		srv.ExportDataHttp(w, req)
	})

	grpcMux.HandlePath("GET", "/v1/export-function", func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		srv.ExportFuntionWithHttp(w, req)
	})

	// grpcMux.HandlePath("GET", "/v1/export", func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
	// 	srv.ExportDataHttp(w, req)
	// })

	// grpcMux.HandlePath("GET", "/v1/preview", func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
	// 	srv.PreviewWithHttp(w, req)
	// })

	// grpcMux.HandlePath("GET", "/v1/preview1", func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
	// 	fmt.Println("trongnhat1")

	// 	fmt.Fprintf(w, "<html><body><img src=\"/image\" /></body></html>")
	// })

	// grpcMux.HandlePath("GET", "/v1/preview/test", func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
	// 	fmt.Println("trongnhat")
	// 	http.ServeFile(w, req, "/home/nhatnt/nhatnt/probationary-project/server-test/server-gateway/DataImportToDB.xlsx")
	// })

	grpcMux.HandlePath("POST", "/v1/status", func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		status := req.Header.Get("status")
		idFile := req.Header.Get("idfile")
		resp, err := srv.clientStogare.UpdateStatusFileClient(context.Background(), idFile, status)
		fmt.Println("trongnhat HandlePath")
		if err != nil {
			// logs.Logger.Error("ImportDataWithHttp: Failed to retrieve file from form data")
			http.Error(w, "Failed to retrieve file from form data", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	if err != nil {
		log.Fatal("cannot register handler server custom", err)
	}

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

// func RegisterDeliverooHandlerServerCustom(ctx context.Context, mux *runtime.ServeMux, server pb.DeliverooServer) error {

// 	mux.HandlePath("POST", "/v1/uploadfile", func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
// 		data_ex, err := ImportDataWithHttp(w, req)

// 		if err != nil {
// 			log.Error("cannot hander data to file or form-data")
// 		}
// 		// fmt.Println("test_data", data_ex)

// 		for i := 0; i < len(data_ex); i++ {

// 			data := &pb.ImportDataResquest{
// 				Data: data_ex[i].content,
// 				Name: data_ex[i].name,
// 			}
// 			_, err := server.ImportData(ctx, data)
// 			if err != nil {
// 				log.Error("cannot import data to database")
// 			}
// 		}
// 		msg := &pb.ImportDataRespone{
// 			Notice: "Post Done",
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(msg)
// 	})

// 	return nil
// }
