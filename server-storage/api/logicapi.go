package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"server-test/server-storage/pb_storage"
	"strings"

	"github.com/xuri/excelize/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	log "github.com/sirupsen/logrus"
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

	resp, err := serverStorage.clientDatabase.UploadFileClient(ctx, strings.Split(tempFile.Name(), "/")[2], "xlsx", IdUser, link, float32(file.Size))

	rsp := &pb_storage.FileInfoRespone{
		// Link: absPath + "/" + strings.Split(tempFile.Name(), "/")[2],
		Link: resp,
	}
	return rsp, nil
}

func (serverStorage *ServerStorage) GetListFileUpload(ctx context.Context, res *pb_storage.GetListFileUploadResquest) (*pb_storage.GetListFileUploadRespone, error) {

	idUser := res.GetIduser()

	respDatabase, err := serverStorage.clientDatabase.GetUploadFileInfoClient(ctx, idUser)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Get list file failed")
	}

	var temp []*pb_storage.FileInfo

	for i := 0; i < len(respDatabase.Fileinfo); i++ {
		temp = append(temp, &pb_storage.FileInfo{
			Filename:     respDatabase.Fileinfo[0].Filename,
			Typefile:     respDatabase.Fileinfo[0].Typefile,
			Size:         int64(respDatabase.Fileinfo[0].Size),
			Link:         respDatabase.Fileinfo[0].Link,
			Timecreateat: respDatabase.Fileinfo[0].Createat,
		})
	}

	return &pb_storage.GetListFileUploadRespone{Fileinfo: temp}, nil
	// return nil, status.Errorf(codes.Unimplemented, "method GetListFileUpload not implemented")
}

func (serverStorage *ServerStorage) ExportTemplateFileUpload(ctx context.Context, res *pb_storage.ExportFileResquest) (*pb_storage.ExportFileRespone, error) {

	nameTemplate := res.GetTemplateExport()
	resp, err := serverStorage.clientDatabase.ExportFileTemplateExcelClient(ctx, nameTemplate)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "ExportTemplateFileUpload  failed")
	}

	// var DataInfo []model.TemplateInfoPerson
	var expenseData = [][]interface{}{}

	// var DataInfo []*pb_database.TemplateFilePersonInfo

	for i := 0; i < len(resp.TemplateFilePersonInfo); i++ {
		// DataInfo = append(DataInfo,
		// 	&pb_database.TemplateFilePersonInfo{
		// 		Lastname:    dataInfo[i].LastName,
		// 		Fullname:    dataInfo[i].FullName,
		// 		Firstname:   dataInfo[i].FistName,
		// 		Phonenumber: dataInfo[i].PhoneNumber,
		// 		Address:     dataInfo[i].Address,
		// 	})
		temp := []interface{}{resp.TemplateFilePersonInfo[i].Lastname, resp.TemplateFilePersonInfo[i].Firstname, resp.TemplateFilePersonInfo[i].Fullname, resp.TemplateFilePersonInfo[i].Phonenumber, resp.TemplateFilePersonInfo[i].Address}
		expenseData = append(expenseData, temp)
	}

	f := excelize.NewFile()
	index, _ := f.NewSheet("Sheet1")
	f.SetActiveSheet(index)

	err = f.SetSheetRow("Sheet1", "A1", &[]interface{}{"Last Name", "First Name", "Full Name", "Phone Number", "Address"})
	if err != nil {

		log.Error("ExportData: Error SetSheetRow: ", err)
	}
	err = f.SetColWidth("Sheet1", "A", "G", 30)

	if err != nil {

		log.Error("ExportData: Error SetColWidth: ", err)
	}

	startRow := 2
	for i := startRow; i < (len(expenseData) + startRow); i++ {
		err = f.SetSheetRow("Sheet1", fmt.Sprintf("A%d", i), &expenseData[i-2])
		if err != nil {
			log.Error("ExportData: Error SetSheetRow: ", err)
		}
	}

	// Save spreadsheet by the given path.
	if err := f.SaveAs("./storage-export/TemplateInfoPerson.xlsx"); err != nil {
		fmt.Println(err)
	}

	if err != nil {
		log.Error("ExportData: Error SaveAs: ", err)
	}

	// dir, err := filepath.Abs(filepath.Dir("TemplateInfoPerson.xlsx"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// pathExport := dir + "/stogare-export/DataExportFromDB.xlsx"
	pathExport := "localhost:3000/v1/downloadlink?dir=" + "storage-export/TemplateInfoPerson.xlsx"
	return &pb_storage.ExportFileRespone{PathExport: pathExport}, nil
}

func (serverStorage *ServerStorage) DownloafFile(ctx context.Context, res *pb_storage.DownloadFileResquest) (*pb_storage.DownloadFileRespone, error) {

	dir := res.GetDir()

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
