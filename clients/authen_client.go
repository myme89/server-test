package clients

import (
	"context"
	"errors"
	"server-test/server-authen/pb_authen"

	"google.golang.org/grpc"
)

type AuthenClient struct {
}

var (
	authenGrpcServiceAddr   = "0.0.0.0:3002"
	authenGrpcServiceClient pb_authen.AuthenClient
)

func prepareAuthenGrpcClient(ctx context.Context) error {

	conn, err := grpc.DialContext(ctx, stogareGrpcServiceAddr, []grpc.DialOption{
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

// func (stogareClient *StorageClient) TestData(ctx context.Context, data string) (string, error) {

// 	if err := prepareStorageGrpcClient(ctx); err != nil {
// 		return "prepareStorageGrpcClient failed", err
// 	}

// 	res, err := stogareGrpcServiceClient.TestData(ctx, &pb_storage.DataInfoTestResquest{DataTest: data})
// 	if err != nil {
// 		return "failed", errors.New(status.Convert(err).Message())
// 	}
// 	fmt.Println("TestData server Storage")

// 	noti := res.DataResp + " Done"
// 	return noti, nil
// }

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
