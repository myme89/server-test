package clients

import (
	"context"
	"errors"
	"fmt"
	"server-test/server-storage/pb_storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type StorageClient struct {
}

var (
	stogareGrpcServiceAddr   = "0.0.0.0:3003"
	stogareGrpcServiceClient pb_storage.StogareClient
)

func prepareStorageGrpcClient(ctx context.Context) error {

	conn, err := grpc.DialContext(ctx, stogareGrpcServiceAddr, []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock()}...)

	if err != nil {
		stogareGrpcServiceClient = nil
		return errors.New("connection to advice gRPC service failed")
	}

	if stogareGrpcServiceClient != nil {
		conn.Close()
		return nil
	}

	stogareGrpcServiceClient = pb_storage.NewStogareClient(conn)
	return nil
}

func (stogareClient *StorageClient) TestData(ctx context.Context, data string) (string, error) {

	if err := prepareStorageGrpcClient(ctx); err != nil {
		return "prepareStorageGrpcClient failed", err
	}

	res, err := stogareGrpcServiceClient.TestData(ctx, &pb_storage.DataInfoTestResquest{DataTest: data})
	if err != nil {
		return "failed", errors.New(status.Convert(err).Message())
	}
	fmt.Println("TestData server Storage")

	noti := res.DataResp + " Done"
	return noti, nil
}
