package clients

import (
	"context"
	"errors"
	"fmt"
	"server-test/server-database/pb_database"
	"server-test/server-proccess-data/model"

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

func (databaseClient *DatabaseClient) UploadDataFileExcelClient(ctx context.Context, infofile []model.InfoFile, fileName string) (*pb_database.ImportFileExcelRespone, error) {
	if err := prepareDatabaseGrpcClient(ctx); err != nil {
		return nil, err

	}

	var temp []*pb_database.ImportFileExcel

	for i := 0; i < len(infofile); i++ {
		temp = append(temp, &pb_database.ImportFileExcel{
			Filename: fileName + " " + infofile[i].Name,
			Content:  infofile[i].Content,
		})
	}

	// infoFileExcel := &pb_database.{Filename: fileName, Typefile: typeFile, Iduser: idUser, Size: size, Link: link}

	res, err := databaseGrpcServiceClient.ImportFileExcel(ctx, &pb_database.ImportFileExcelResquest{ImportFileExcel: temp})

	if err != nil {
		return nil, err

	}
	fmt.Println("UploadFileClient Storage", res)

	// noti := res.Noti
	return &pb_database.ImportFileExcelRespone{Noti: "noti"}, nil
}

func (databaseClient *DatabaseClient) UpdateStatusProcessingClient(ctx context.Context, status, idFile string) (*pb_database.StatusProcessingFileRespone, error) {
	if err := prepareDatabaseGrpcClient(ctx); err != nil {
		fmt.Println("prepareStorageGrpcClient failed", err)
		return nil, err
	}

	resp, err := databaseGrpcServiceClient.UpdateStatusProcessingFileExcel(ctx, &pb_database.StatusProcessingFileResquest{Status: status, IdFile: idFile})

	if err != nil {
		fmt.Println("UpdateStatusProcessingFileExcel failed", err)
		return nil, err
	}
	fmt.Println("UpdateStatusProcessingClient ", resp)

	return resp, nil
}
