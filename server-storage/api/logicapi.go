package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"server-test/server-storage/database/mongodb"
	"server-test/server-storage/pb_storage"
	"strings"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// import (
// 	"context"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"path/filepath"
// 	"server-test/server-storage/pb_storage"
// 	"strings"

// 	"github.com/xuri/excelize/v2"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"

// 	log "github.com/sirupsen/logrus"
// )

// func (serverStorage *ServerStorage) TestData(ctx context.Context, res *pb_storage.DataInfoTestResquest) (*pb_storage.DataInfoTestRespone, error) {

// 	temp := res.GetDataTest() + "trong nhat 1 processing"

// 	// temp1, err := serverStorage.client.TestData(ctx, temp)
// 	// fmt.Println("TestData server Storage  3= ", ctx)
// 	if res.GetDataTest() == "1" {
// 		fmt.Println("TestData server Storage  1")
// 		go func(temp1 string) {
// 			temp2, err := serverStorage.clientTest.TestData(context.Background(), temp1)

// 			fmt.Println("TestData server Storage  1", temp2)
// 			fmt.Println("TestData server Storage  1", err)

// 		}(temp)
// 	}
// 	fmt.Println("TestData server Storage")
// 	// fmt.Println("TestData server Storage", err)
// 	return &pb_storage.DataInfoTestRespone{DataResp: temp}, nil
// }

func (serverStorage *ServerStorage) UploadFile(ctx context.Context, res *pb_storage.FileInfoResquest) (*pb_storage.FileInfoRespone, error) {

	file := res.GetFile()

	fmt.Println(file.Filename)
	fmt.Println(file.Size)

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

	fmt.Println("Temporary file created:", float32(file.Size))
	fmt.Println("Temporary file created1:", file.Size)

	link := absPath + "/" + strings.Split(tempFile.Name(), "/")[2]

	err = mongodb.AddInfoUploadFile(serverStorage.config, strings.Split(tempFile.Name(), "/")[2], IdUser, "xlsx", link, float32(file.Size))
	if err != nil {
		log.Error("Upload info file  faied ", err)
		return nil, status.Errorf(codes.Unimplemented, "Upload info file  faied")
	}

	idFile, err := mongodb.FilterIdFile(serverStorage.config, strings.Split(tempFile.Name(), "/")[2])

	if err != nil {
		log.Error("Upload info file  faied ", err)
		return nil, status.Errorf(codes.Unimplemented, "Get ID file faied")
	}

	// resp, err := serverStorage.clientDatabase.UploadFileClient(ctx, strings.Split(tempFile.Name(), "/")[2], "xlsx", IdUser, link, float32(file.Size))

	rsp := &pb_storage.FileInfoRespone{
		// Link: absPath + "/" + strings.Split(tempFile.Name(), "/")[2],
		Link: idFile,
	}
	return rsp, nil
}

func (serverStorage *ServerStorage) GetListFileUpload(ctx context.Context, res *pb_storage.GetListFileUploadResquest) (*pb_storage.GetListFileUploadRespone, error) {

	idUser := res.GetIduser()

	listFile, err := mongodb.GetListFile(serverStorage.config, idUser)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "get list file failed")
	}

	var temp []*pb_storage.FileInfo

	for i := 0; i < len(listFile); i++ {
		temp = append(temp, &pb_storage.FileInfo{
			Filename:     listFile[0].FileName,
			Typefile:     listFile[0].TypeFile,
			Size:         int64(listFile[0].Size),
			Link:         listFile[0].Link,
			Timecreateat: listFile[0].CreateAt,
		})
	}

	return &pb_storage.GetListFileUploadRespone{Fileinfo: temp}, nil
}

func (serverStorage *ServerStorage) DownloafFile(ctx context.Context, res *pb_storage.DownloadFileResquest) (*pb_storage.DownloadFileRespone, error) {

	idFile := res.GetIdFile()

	dir, err := mongodb.GetDirFile(serverStorage.config, idFile)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "method GetDirFile failded")
	}

	fmt.Println(dir)

	file, err := os.Open(dir)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "method DownloafFile failded")
	}

	content, err := ioutil.ReadAll(file)

	name := file.Name()

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "method DownloafFile ReadAll failded")
	}

	return &pb_storage.DownloadFileRespone{Name: name, Content: content}, nil
}

func (serverStorage *ServerStorage) UpdateStatusUploadFile(ctx context.Context, res *pb_storage.UpdateStatusResquest) (*pb_storage.UpdateStatusRespone, error) {

	fmt.Println("trongnhat")

	status := res.GetStatus()
	idFile := res.GetIdFile()

	err := mongodb.UpdateStatus(serverStorage.config, idFile, status)

	if err != nil {
		log.Error(" UpdateStatus failed:", err)
	}
	return &pb_storage.UpdateStatusRespone{Noti: "Done"}, nil
}

func (serverStorage *ServerStorage) GetShortInfoFileUpload(ctx context.Context, res *pb_storage.GetShortInfoFileUploadResquest) (*pb_storage.GetShortInfoFileUploadRespone, error) {

	idFile := res.GetIdFile()

	shortInfo, err := mongodb.GetShortInfoFile(serverStorage.config, idFile)
	// fmt.Println("test= ", shortInfo)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "get short info file failed")
	}

	// var temp *pb_storage.FileInfo

	temp := &pb_storage.FileInfo{
		Filename:     shortInfo[0].FileName,
		Typefile:     shortInfo[0].TypeFile,
		Size:         int64(shortInfo[0].Size),
		Link:         shortInfo[0].Link,
		Timecreateat: shortInfo[0].CreateAt,
	}
	// return nil, status.Errorf(codes.InvalidArgument, "get short info file failed")

	return &pb_storage.GetShortInfoFileUploadRespone{Fileinfo: temp}, nil
}
