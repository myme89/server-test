package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"server-test/server-storage/pb_storage"
)

func (serverStorage *ServerStorage) TestData(ctx context.Context, res *pb_storage.DataInfoTestResquest) (*pb_storage.DataInfoTestRespone, error) {

	temp := res.GetDataTest() + "trong nhat 1 processing"

	// temp1, err := serverStorage.client.TestData(ctx, temp)
	// fmt.Println("TestData server Storage  3= ", ctx)
	if res.GetDataTest() == "1" {
		fmt.Println("TestData server Storage  1")
		go func(temp1 string) {
			temp2, err := serverStorage.client.TestData(context.Background(), temp1)

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
	// fmt.Println(file.Content)

	// xlsx, err := excelize.OpenReader(bytes.NewReader(file.Content))
	// xlsx, err := excelize.OpenReader(file_ex)

	// if err != nil {
	// 	logs.Logger.Error("ImportDataWithHttp: Failed to open Excel file ", err)
	// 	return nil, status.Errorf(codes.Unimplemented, "Failed to open Excel file")
	// }

	// fmt.Println(xlsx)

	tempFile, err := ioutil.TempFile("./folder-storage-file", "upload-*.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	tempFile.Write(file.Content)

	// os.Remove(tempFile.Name()) // Remove the temporary file when the program exits

	// Get the directory of the temporary file
	dir := filepath.Dir(tempFile.Name())
	fmt.Println("Directory of the temporary file:", dir)

	folderName := "storage1"

	// Create the folder
	err = os.Mkdir(folderName, 0755)
	if err != nil {
		panic(err)
	}

	absPath, err := filepath.Abs(folderName)
	if err != nil {
		panic(err)
	}

	// Print the absolute path of the folder
	fmt.Println("Folder created:", absPath)

	fmt.Println("Temporary file created:", tempFile.Name())
	rsp := &pb_storage.FileInfoRespone{
		Link: serverStorage.Addr + tempFile.Name(),
	}
	return rsp, nil
}
