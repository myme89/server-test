package api

import (
	"context"
	"fmt"
	"server-test/server-authen/pb_authen"
)

func (serverAuthen *ServerAuthen) SignUp(ctx context.Context, res *pb_authen.UserResquest) (*pb_authen.UserRespone, error) {

	infoSignUp := res.GetUserinfo()

	fmt.Println("username: ", infoSignUp.Username)
	fmt.Println("password: ", infoSignUp.Password)
	fmt.Println("lastname: ", infoSignUp.Lastname)
	fmt.Println("firstname: ", infoSignUp.Firstname)

	resp, err := serverAuthen.clientDatabase.SignUpAcc(ctx, infoSignUp.Username, infoSignUp.Password, infoSignUp.Firstname, infoSignUp.Lastname)

	if err != nil {
		fmt.Println(err)
	}

	noti := &pb_authen.UserRespone{Noti: resp}

	return noti, nil
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
