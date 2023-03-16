package clients

import (
	"context"
	"errors"
	"fmt"
	"server-test/server-proccess-data/pb_processing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type StorageClient struct {
}

var (
	processingGrpcServiceAddr   = "0.0.0.0:3001"
	processingGrpcServiceClient pb_processing.ProcessingClient
)

func prepareProcessingGrpcClient(ctx context.Context) error {

	conn, err := grpc.DialContext(ctx, processingGrpcServiceAddr, []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock()}...)

	if err != nil {
		processingGrpcServiceClient = nil
		return errors.New("connection to advice gRPC service failed")
	}

	if processingGrpcServiceClient != nil {
		conn.Close()
		return nil
	}

	processingGrpcServiceClient = pb_processing.NewProcessingClient(conn)
	return nil
}

func (stogareClient *StorageClient) TestData(ctx context.Context, data string) (string, error) {
	fmt.Println(ctx)
	if err := prepareProcessingGrpcClient(ctx); err != nil {
		return "prepareStorageGrpcClient failed", err
	}

	res, err := processingGrpcServiceClient.TestData2(ctx, &pb_processing.DataInfoTestResquest1{DataTest: data})
	if err != nil {
		return "failed", errors.New(status.Convert(err).Message())
	}
	fmt.Println("TestData server Storage")

	noti := res.DataResp + " Done TestData2"
	return noti, nil
}
