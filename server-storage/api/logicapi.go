package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"server-test/server-storage/pb_storage"
	"strings"
)

func (serverStorage *ServerStorage) TestData(ctx context.Context, res *pb_storage.DataInfoTestResquest) (*pb_storage.DataInfoTestRespone, error) {

	temp := res.GetDataTest() + "trong nhat 1 processing"

	// temp1, err := serverStorage.client.TestData(ctx, temp)
	// fmt.Println("TestData server Storage  3= ", ctx)
	if res.GetDataTest() == "1" {
		fmt.Println("TestData server Storage  1")
		go func(temp1 string) {
			temp2, err := serverStorage.clientTest.TestData(context.Background(), temp1)

			fmt.Println("TestData server Storage  1", temp2)
			fmt.Println("TestData server Storage  1", err)

		}(temp)
	}
	fmt.Println("TestData server Storage")
	// fmt.Println("TestData server Storage", err)
	return &pb_storage.DataInfoTestRespone{DataResp: temp}, nil
}

func (serverStorage *ServerStorage) UploadFile(ctx context.Context, res *pb_storage.FileInfoResquest) (*pb_storage.FileInfoRespone, error) {

	file := res.GetFile()

	fmt.Println(file.Filename)
	fmt.Println(file.Typefile)

	IdUser := res.GetIduser()

	// Create the folder
	folderName := "storage"
	// err := os.Mkdir(folderName, 0755)
	// if err != nil {
	// 	panic(err)
	// }

	fileName := strings.Split(file.Filename, ".")

	fmt.Println(fileName)
	tempFile, err := ioutil.TempFile("./storage", fileName[0]+"-*.xlsx")

	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	tempFile.Write(file.Content)

	// os.Remove(tempFile.Name()) // Remove the temporary file when the program exits

	// Get the directory of the temporary file
	// dir := filepath.Dir(tempFile.Name())

	fmt.Println("Directory of the temporary file:", strings.Split(tempFile.Name(), "/")[2])

	absPath, err := filepath.Abs(folderName)
	if err != nil {
		panic(err)
	}

	// Print the absolute path of the folder
	fmt.Println("Folder created:", absPath)

	fmt.Println("Temporary file created:", IdUser)

	resp, err := serverStorage.clientDatabase.UploadFileClient(ctx, strings.Split(tempFile.Name(), "/")[2], "xlsx", IdUser, 10.32)

	rsp := &pb_storage.FileInfoRespone{
		// Link: absPath + "/" + strings.Split(tempFile.Name(), "/")[2],
		Link: resp,
	}
	return rsp, nil
}
