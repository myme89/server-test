package api

import (
	"context"
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

// func (serverStorage *ServerStorage) UploadFile(ctx context.Context, res *pb_storage.FileInfoResquest) (*pb_storage.FileInfoRespone, error) {

// 	file := res.GetFile()

// 	fmt.Println(file.Filename)
// 	fmt.Println(file.Typefile)
// 	// fmt.Println(file.Content)

// 	// xlsx, err := excelize.OpenReader(bytes.NewReader(file.Content))
// 	// xlsx, err := excelize.OpenReader(file_ex)

// 	// if err != nil {
// 	// 	logs.Logger.Error("ImportDataWithHttp: Failed to open Excel file ", err)
// 	// 	return nil, status.Errorf(codes.Unimplemented, "Failed to open Excel file")
// 	// }

// 	// fmt.Println(xlsx)

// 	tempFile, err := ioutil.TempFile("./folder-storage-file", "upload-*.xlsx")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer tempFile.Close()

// 	tempFile.Write(file.Content)

// 	// os.Remove(tempFile.Name()) // Remove the temporary file when the program exits

// 	// Get the directory of the temporary file
// 	dir := filepath.Dir(tempFile.Name())
// 	fmt.Println("Directory of the temporary file:", dir)

// 	folderName := "storage1"

// 	// Create the folder
// 	err = os.Mkdir(folderName, 0755)
// 	if err != nil {
// 		panic(err)
// 	}

// 	absPath, err := filepath.Abs(folderName)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Print the absolute path of the folder
// 	fmt.Println("Folder created:", absPath)

// 	fmt.Println("Temporary file created:", tempFile.Name())
// 	rsp := &pb_storage.FileInfoRespone{
// 		Link: serverStorage.Addr + tempFile.Name(),
// 	}
// 	return rsp, nil
// }
