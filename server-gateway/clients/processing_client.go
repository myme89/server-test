package clients

import (
	"context"
	"errors"
	"fmt"
	"server-test/server-proccess-data/pb_processing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type ProcessingClient struct {
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

// func (stogareClient *StorageClient) TestData(ctx context.Context, data string) (string, error) {
// 	fmt.Println(ctx)
// 	if err := prepareProcessingGrpcClient(ctx); err != nil {
// 		return "prepareStorageGrpcClient failed", err
// 	}

// 	res, err := processingGrpcServiceClient.TestData2(ctx, &pb_processing.DataInfoTestResquest1{DataTest: data})
// 	if err != nil {
// 		return "failed", errors.New(status.Convert(err).Message())
// 	}
// 	fmt.Println("TestData server Storage")

// 	noti := res.DataResp + " Done TestData2"
// 	return noti, nil
// }

func (processingClient *ProcessingClient) ProcessingDataClient(ctx context.Context, idFile, fileName string, content []byte) (string, error) {
	fmt.Println(ctx)
	if err := prepareProcessingGrpcClient(ctx); err != nil {
		return "ProcessingDataClient failed", err
	}

	infoFile := &pb_processing.FileInfoProcess{Filename: fileName, Content: content, Idfile: idFile}
	resp, err := processingGrpcServiceClient.ProcessingFileExcel(ctx, &pb_processing.ProcessingFileResquest{Fileinfoprocess: infoFile})
	if err != nil {
		return "processingGrpcServiceClient ProcessingFileExcel failed", errors.New(status.Convert(err).Message())
	}
	fmt.Println("TestData server Storage")

	return resp.Noti, nil
}
