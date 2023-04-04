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

func (stogareClient *StorageClient) UploadFile(ctx context.Context, fileName, fileType, idUser string, size int64, content []byte) (*pb_storage.FileInfoRespone, error) {

	fmt.Println("stogareClient Upload file")
	if err := prepareStorageGrpcClient(ctx); err != nil {
		return nil, err
	}

	file := &pb_storage.FileInfo{Filename: fileName, Typefile: fileType, Content: content, Size: size}

	res, err := stogareGrpcServiceClient.UploadFile(ctx, &pb_storage.FileInfoResquest{File: file, Iduser: idUser})
	if err != nil {
		return nil, err
	}

	// noti := res.Link
	return res, nil
}

func (stogareClient *StorageClient) GetUploadFileInfoClient(ctx context.Context, idUser string) (*pb_storage.GetListFileUploadRespone, error) {
	if err := prepareStorageGrpcClient(ctx); err != nil {
		return nil, err
	}

	resp, err := stogareGrpcServiceClient.GetListFileUpload(ctx, &pb_storage.GetListFileUploadResquest{Iduser: idUser})
	if err != nil {
		return nil, err
	}
	fmt.Println("UploadFileClient Storage")

	return resp, nil
}

func (stogareClient *StorageClient) DownloadFileClient(ctx context.Context, idFile string) (*pb_storage.DownloadFileRespone, error) {

	if err := prepareStorageGrpcClient(ctx); err != nil {
		return nil, err
	}

	resp, err := stogareGrpcServiceClient.DownloafFile(ctx, &pb_storage.DownloadFileResquest{IdFile: idFile})

	if err != nil {
		return nil, err
	}
	fmt.Println("UploadFileClient Storage")

	return resp, nil
	// return nil, status.Errorf(codes.InvalidArgument, "method ExportFileTemplateExcelClient in client stogare failed")

}

func (stogareClient *StorageClient) UpdateStatusFileClient(ctx context.Context, idFile, status string) (*pb_storage.UpdateStatusRespone, error) {

	if err := prepareStorageGrpcClient(ctx); err != nil {
		return nil, err
	}

	resp, err := stogareGrpcServiceClient.UpdateStatusUploadFile(ctx, &pb_storage.UpdateStatusResquest{Status: status, IdFile: idFile})

	if err != nil {
		return nil, err
	}
	fmt.Println("UpdateStatusFileClient Storage")

	return resp, nil

}

func (stogareClient *StorageClient) GetUploadFileShortInfoClient(ctx context.Context, idFile string) (*pb_storage.GetShortInfoFileUploadRespone, error) {
	if err := prepareStorageGrpcClient(ctx); err != nil {
		return nil, err
	}

	resp, err := stogareGrpcServiceClient.GetShortInfoFileUpload(ctx, &pb_storage.GetShortInfoFileUploadResquest{IdFile: idFile})
	if err != nil {
		return nil, err
	}
	fmt.Println("GetUploadFileShortInfoClient Storage")

	return resp, nil
}
