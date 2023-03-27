package clients

import (
	"context"
	"errors"
	"server-test/server-database/pb_database"

	"google.golang.org/grpc"
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

// func (databaseClient *DatabaseClient) GetUploadFileInfoClient(ctx context.Context, idUser string) (*pb_database.GetListFileRespone, error) {
// 	if err := prepareDatabaseGrpcClient(ctx); err != nil {
// 		return nil, err
// 	}

// 	resp, err := databaseGrpcServiceClient.GetListUploadFile(ctx, &pb_database.GetListFileResquest{IdUser: idUser})
// 	if err != nil {
// 		return nil, errors.New(status.Convert(err).Message())
// 	}
// 	fmt.Println("UploadFileClient Storage")

// 	return resp, nil
// }

// func (databaseClient *DatabaseClient) ExportFileTemplateExcelClient(ctx context.Context, templateName string) (string, error) {
// 	if err := prepareDatabaseGrpcClient(ctx); err != nil {
// 		return "prepareDatabaseGrpcClient faild", err
// 	}

// 	resp, err := databaseGrpcServiceClient.ExportTemplateFile(ctx, &pb_database.ExportTemplateFileResquest{TemplateName: templateName})

// 	if err != nil {
// 		return "failed", errors.New(status.Convert(err).Message())
// 	}
// 	fmt.Println("UploadFileClient Storage")

// 	return resp.PathExport, nil
// }
