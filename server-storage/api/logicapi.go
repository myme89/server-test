package api

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"server-test/server-storage/database/mongodb"
	"server-test/server-storage/pb_storage"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

	IdUser := res.GetIduser()

	// Create the folder
	folderName := "storage"

	fileNameSplit := strings.Split(file.Filename, ".")

	now := time.Now()

	dateTimeNow := now.Format("20060102150405")

	fileName := fileNameSplit[0] + dateTimeNow + "." + fileNameSplit[1]

	hashNameFile := HashNameFile(fileNameSplit[0]+dateTimeNow) + "." + fileNameSplit[1]

	fmt.Println(hashNameFile)

	absPath, err := filepath.Abs(folderName)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(absPath+"/"+hashNameFile, file.Content, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}

	// Print the absolute path of the folder
	// fmt.Println("Folder created:", absPath)

	// fmt.Println("Temporary file created:", float32(file.Size))
	// fmt.Println("Temporary file created1:", file.Size)

	link := absPath + "/" + fileName

	err = mongodb.AddInfoUploadFile(serverStorage.config, fileName, IdUser, "xlsx", link, float32(file.Size))
	if err != nil {
		log.Error("Upload info file  faied ", err)
		return nil, status.Errorf(codes.Unimplemented, "Upload info file  faied")
	}

	idFile, err := mongodb.FilterIdFile(serverStorage.config, fileName)

	if err != nil {
		log.Error("Upload info file  faied ", err)
		return nil, status.Errorf(codes.Unimplemented, "Get ID file faied")
	}

	rsp := &pb_storage.FileInfoRespone{
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
			Fileid:       listFile[i].Id,
			Filename:     listFile[i].FileName,
			Typefile:     listFile[i].TypeFile,
			Size:         int64(listFile[i].Size),
			Link:         listFile[i].Link,
			Timecreateat: listFile[i].CreateAt,
		})
	}

	return &pb_storage.GetListFileUploadRespone{Fileinfo: temp}, nil
}

func (serverStorage *ServerStorage) DownloafFile(ctx context.Context, res *pb_storage.DownloadFileResquest) (*pb_storage.DownloadFileRespone, error) {

	idFile := res.GetIdFile()
	dir, nameFile, err := mongodb.GetDirFile(serverStorage.config, idFile)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "method GetDirFile failded")
	}
	hashNameFile := HashNameFile((strings.Split(nameFile, "."))[0]) + "." + (strings.Split(nameFile, "."))[1]
	fmt.Println((strings.Split(nameFile, "."))[0])

	fmt.Println(hashNameFile)
	dirSpilt := strings.Split(dir, "/")

	dirSpilt[len(dirSpilt)-1] = hashNameFile

	dirFile := strings.Join(dirSpilt, "/")

	fmt.Println(dirFile)
	file, err := os.Open(dirFile)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "method DownloafFile failded")
	}

	content, err := ioutil.ReadAll(file)

	// name := file.Name()

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "method DownloafFile ReadAll failded")
	}

	return &pb_storage.DownloadFileRespone{Name: nameFile, Content: content}, nil
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

func HashNameFile(nameFile string) string {

	key := []byte("trongnhat99tn")

	// compute HMAC-SHA256 hash of the message with key
	hash := hmac.New(sha256.New, key)
	hash.Write([]byte(nameFile))
	mac := hash.Sum(nil)

	macString := hex.EncodeToString(mac)
	return macString
}
