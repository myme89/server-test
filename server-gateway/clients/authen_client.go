package clients

import (
	"context"
	"errors"
	"fmt"
	"server-test/server-authen/pb_authen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type AuthenClient struct {
}

var (
	authenGrpcServiceAddr   = "0.0.0.0:3002"
	authenGrpcServiceClient pb_authen.AuthenClient
)

func prepareAuthenGrpcClient(ctx context.Context) error {

	conn, err := grpc.DialContext(ctx, authenGrpcServiceAddr, []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock()}...)

	if err != nil {
		authenGrpcServiceClient = nil
		return errors.New("connection to advice gRPC service failed")
	}

	if authenGrpcServiceClient != nil {
		conn.Close()
		return nil
	}

	authenGrpcServiceClient = pb_authen.NewAuthenClient(conn)
	return nil
}

func (authenClient *AuthenClient) SignUp(ctx context.Context, userName, password, lastName, firstName string) (string, error) {

	fmt.Println("SignUp AuthenClient")
	if err := prepareAuthenGrpcClient(ctx); err != nil {
		return "prepareStorageGrpcClient failed", err
	}

	userInfo := &pb_authen.UserInfo{Username: userName, Password: password, Firstname: firstName, Lastname: lastName}

	res, err := authenGrpcServiceClient.SignUp(ctx, &pb_authen.UserResquest{Userinfo: userInfo})
	if err != nil {
		return "failed", errors.New(status.Convert(err).Message())
	}

	// noti :=
	return res.Noti, nil
}

// func (stogareClient *StorageClient) UploadFile(ctx context.Context, fileName, fileType string, content []byte) (string, error) {

// 	fmt.Println("stogareClient Upload file")
// 	if err := prepareStorageGrpcClient(ctx); err != nil {
// 		return "prepareStorageGrpcClient failed", err
// 	}

// 	file := &pb_storage.FileInfo{Filename: fileName, Typefile: fileType, Content: content}

// 	res, err := stogareGrpcServiceClient.UploadFile(ctx, &pb_storage.FileInfoResquest{File: file})
// 	if err != nil {
// 		return "failed", errors.New(status.Convert(err).Message())
// 	}

// 	noti := res.Link
// 	return noti, nil
// }
