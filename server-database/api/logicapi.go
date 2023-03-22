package api

import (
	"context"
	"fmt"
	"server-test/server-database/database/mongodb"
	"server-test/server-database/pb_database"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	log "github.com/sirupsen/logrus"
)

func (serverDatabase *ServerDatabase) SignUpAcc(ctx context.Context, res *pb_database.SignUpAccResquest) (*pb_database.SignUpAccRespone, error) {

	infoUser := res.GetUserinfo()

	err := mongodb.AddUserInfoSignUp(serverDatabase.config, infoUser.Username, infoUser.Password, infoUser.Lastname, infoUser.Firstname)

	if err != nil {
		log.Error("Create user sign up acc faied ", err)
		return nil, status.Errorf(codes.Unimplemented, "get Data failed")
	}

	return &pb_database.SignUpAccRespone{Noti: "Done"}, nil
}

func (serverDatabase *ServerDatabase) LoginAcc(ctx context.Context, res *pb_database.LoginAccResquest) (*pb_database.LoginAccRespone, error) {

	infoUser := res.GetUserinfo()

	infoAcc := mongodb.GetHashPassword(serverDatabase.config, infoUser.Username)

	// if err != nil {
	// 	log.Error("Create user sign up acc faied ", err)
	// 	return nil, status.Errorf(codes.Unimplemented, "get Data failed")
	// }

	if infoAcc.Password != infoUser.Password || len(infoAcc.LastName) == 0 || len(infoAcc.FistName) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "User or password is incorrect")
	}

	fmt.Println("nhatnt")
	return &pb_database.LoginAccRespone{Userinfo: &pb_database.UserAccInfo{Username: infoUser.Username, Lastname: infoAcc.LastName, Firstname: infoAcc.FistName, Id: infoAcc.Id}}, nil

}

func (serverDatabase *ServerDatabase) UploadFile(ctx context.Context, res *pb_database.UploadFileResquest) (*pb_database.UploadFileRespone, error) {
	infoUploadFile := res.GetFileUploadInfo()

	err := mongodb.AddInfoUploadFile(serverDatabase.config, infoUploadFile.Filename, infoUploadFile.Iduser, infoUploadFile.Typefile, infoUploadFile.Size)
	if err != nil {
		log.Error("Upload info file  faied ", err)
		return nil, status.Errorf(codes.Unimplemented, "Upload info file  faied")
	}

	return &pb_database.UploadFileRespone{Noti: "Done"}, nil
}
