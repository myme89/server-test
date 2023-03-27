package clients

import (
	"context"
	"errors"
	"fmt"
	"server-test/server-database/pb_database"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DatabaseClient struct {
}

var (
	databaseGrpcServiceAddr   = "0.0.0.0:3004"
	databaseGrpcServiceClient pb_database.DatabaseClient
)

func prepareDatabaseGrpcClient(ctx context.Context) error {

	conn, err := grpc.DialContext(ctx, databaseGrpcServiceAddr, []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock()}...)

	if err != nil {
		databaseGrpcServiceClient = nil
		return errors.New("connection to advice gRPC service failed")
	}

	if databaseGrpcServiceClient != nil {
		conn.Close()
		return nil
	}

	databaseGrpcServiceClient = pb_database.NewDatabaseClient(conn)
	return nil
}

func (databaseClient *DatabaseClient) UploadFileClient(ctx context.Context, fileName, typeFile, idUser, link string, size float32) (string, error) {
	if err := prepareDatabaseGrpcClient(ctx); err != nil {
		return "prepareStorageGrpcClient failed", err
	}

	upLoadFileInfo := &pb_database.FileUploadInfo{Filename: fileName, Typefile: typeFile, Iduser: idUser, Size: size, Link: link}

	res, err := databaseGrpcServiceClient.UploadFile(ctx, &pb_database.UploadFileResquest{FileUploadInfo: upLoadFileInfo})
	if err != nil {
		return "failed", errors.New(status.Convert(err).Message())
	}
	fmt.Println("UploadFileClient Storage")

	noti := res.Noti
	return noti, nil
}

func (databaseClient *DatabaseClient) GetUploadFileInfoClient(ctx context.Context, idUser string) (*pb_database.GetListFileRespone, error) {
	if err := prepareDatabaseGrpcClient(ctx); err != nil {
		return nil, err
	}

	resp, err := databaseGrpcServiceClient.GetListUploadFile(ctx, &pb_database.GetListFileResquest{IdUser: idUser})
	if err != nil {
		return nil, errors.New(status.Convert(err).Message())
	}
	fmt.Println("UploadFileClient Storage")

	return resp, nil
}

func (databaseClient *DatabaseClient) ExportFileTemplateExcelClient(ctx context.Context, templateName string) (*pb_database.ExportTemplateFileRespone, error) {
	if err := prepareDatabaseGrpcClient(ctx); err != nil {
		return nil, err
	}

	resp, err := databaseGrpcServiceClient.ExportTemplateFile(ctx, &pb_database.ExportTemplateFileResquest{TemplateName: templateName})

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "ExportFileTemplateExcelClient in stogare failed")
	}
	fmt.Println("UploadFileClient Storage")

	return resp, nil
}
