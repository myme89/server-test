package clients

import (
	"context"
	"errors"
	"fmt"
	"server-test/server-proccess-data/pb_processing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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

func (processingClient *ProcessingClient) ProcessingDataClient(ctx context.Context, idFile, fileName, linkFile string) (*pb_processing.ProcessingFileRespone, error) {
	fmt.Println(ctx)
	if err := prepareProcessingGrpcClient(ctx); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "method prepareProcessingGrpcClient in ProcessingDataClient failed")
	}

	infoFile := &pb_processing.FileInfoProcess{Filename: fileName, LinkFile: linkFile, Idfile: idFile}
	resp, err := processingGrpcServiceClient.ProcessingFileExcel(ctx, &pb_processing.ProcessingFileResquest{Fileinfoprocess: infoFile})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "method ProcessingFileExcel in ProcessingDataClient failed")
	}
	fmt.Println("ProcessingDataClient server Storage")

	return resp, nil
}

func (processingClient *ProcessingClient) ExportFileTemplateExcelClient(ctx context.Context, templateName string) (*pb_processing.ExportFileRespone, error) {

	if err := prepareProcessingGrpcClient(ctx); err != nil {
		return nil, err
	}

	fmt.Println("trongnhat test1", processingGrpcServiceClient)

	resp, err := processingGrpcServiceClient.ExportTemplateFileUpload(ctx, &pb_processing.ExportFileResquest{TemplateExport: templateName})

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "method ExportFileTemplateExcelClient in client process failed")
	}
	fmt.Println("UploadFileClient Process")

	return resp, nil
	// return nil, status.Errorf(codes.InvalidArgument, "method ExportFileTemplateExcelClient in client process failed")

}

func (processingClient *ProcessingClient) ExportFuntionClient(ctx context.Context, accountRec string) (*pb_processing.GetTransactionByAccountRespone, error) {

	if err := prepareProcessingGrpcClient(ctx); err != nil {
		return nil, err
	}

	resp, err := processingGrpcServiceClient.GetTransactionByAccount(ctx, &pb_processing.GetTransactionByAccountResquest{Account: accountRec})

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "method GetTransactionByAccount in client process failed")
	}
	fmt.Println("GetTransactionByAccount Process")

	return resp, nil
	// return nil, status.Errorf(codes.InvalidArgument, "method ExportFileTemplateExcelClient in client process failed")

}
