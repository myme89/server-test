package clients

import (
	"context"
	"errors"
	"fmt"
	"server-test/server-database/pb_database"

	"google.golang.org/grpc"
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

func (databaseClient *DatabaseClient) SignUpAcc(ctx context.Context, userName, password, firstName, lastName string) (string, error) {
	if err := prepareDatabaseGrpcClient(ctx); err != nil {
		return "prepareStorageGrpcClient failed", err
	}

	userInfo := &pb_database.UserAccInfo{Username: userName, Password: password, Firstname: firstName, Lastname: lastName}

	res, err := databaseGrpcServiceClient.SignUpAcc(ctx, &pb_database.SignUpAccResquest{Userinfo: userInfo})
	if err != nil {
		return "failed", errors.New(status.Convert(err).Message())
	}
	fmt.Println("TestData server Storage")

	noti := res.Noti
	return noti, nil
}

func (databaseClient *DatabaseClient) LoginAccClient(ctx context.Context, userName, password string) (*pb_database.LoginAccRespone, error) {
	if err := prepareDatabaseGrpcClient(ctx); err != nil {
		return nil, err
	}

	userInfo := &pb_database.UserAccInfo{Username: userName, Password: password}

	res, err := databaseGrpcServiceClient.LoginAcc(ctx, &pb_database.LoginAccResquest{Userinfo: userInfo})

	if err != nil {
		return nil, errors.New(status.Convert(err).Message())
	}

	fmt.Println("TestData server Storage", res)

	return res, nil
}
