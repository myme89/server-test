package api

import (
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"
	"server-test/server-database/database/mongodb"
	"server-test/server-database/model"
	"server-test/server-database/pb_database"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	log "github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
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

	err := mongodb.AddInfoUploadFile(serverDatabase.config, infoUploadFile.Filename, infoUploadFile.Iduser, infoUploadFile.Typefile, infoUploadFile.Link, infoUploadFile.Size)
	if err != nil {
		log.Error("Upload info file  faied ", err)
		return nil, status.Errorf(codes.Unimplemented, "Upload info file  faied")
	}

	idFile, err := mongodb.FilterIdFile(serverDatabase.config, infoUploadFile.Filename)

	if err != nil {
		log.Error("Upload info file  faied ", err)
		return nil, status.Errorf(codes.Unimplemented, "Get ID file faied")
	}

	return &pb_database.UploadFileRespone{Noti: idFile}, nil
}

func (serverDatabase *ServerDatabase) GetListUploadFile(ctx context.Context, res *pb_database.GetListFileResquest) (*pb_database.GetListFileRespone, error) {

	idUser := res.GetIdUser()

	listFile, err := mongodb.GetListFile(serverDatabase.config, idUser)

	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "get list file failed")
	}

	var ListFileInfo []model.FileInfoUpload

	for i := 0; i < len(listFile); i++ {
		ListFileInfo = append(ListFileInfo, model.FileInfoUpload{FileName: listFile[i].FileName, TypeFile: listFile[i].TypeFile, Size: listFile[i].Size, Link: listFile[i].Link, CreateAt: listFile[i].CreateAt})
	}

	rsp := &pb_database.GetListFileRespone{
		Fileinfo: ConvertData(ListFileInfo),
	}

	return rsp, nil
}

func (serverDatabase *ServerDatabase) ExportTemplateFile(ctx context.Context, res *pb_database.ExportTemplateFileResquest) (*pb_database.ExportTemplateFileRespone, error) {

	templateName := res.GetTemplateName()

	dataInfo, err := mongodb.GetAllInfo(serverDatabase.config, templateName)

	var DataInfo []model.TemplateInfoPerson
	var expenseData = [][]interface{}{}

	for i := 0; i < len(dataInfo); i++ {
		DataInfo = append(DataInfo,
			model.TemplateInfoPerson{
				LastName:    dataInfo[i].LastName,
				FullName:    dataInfo[i].FullName,
				FistName:    dataInfo[i].FistName,
				PhoneNumber: dataInfo[i].PhoneNumber,
				Address:     dataInfo[i].Address,
			})
		temp := []interface{}{dataInfo[i].LastName, dataInfo[i].FistName, dataInfo[i].FullName, dataInfo[i].PhoneNumber, dataInfo[i].Address}
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
	if err := f.SaveAs("./export/TemplateInfoPerson.xlsx"); err != nil {
		fmt.Println(err)
	}

	if err != nil {
		log.Error("ExportData: Error SaveAs: ", err)
	}

	dir, err := filepath.Abs(filepath.Dir("TemplateInfoPerson.xlsx"))
	if err != nil {
		log.Fatal(err)
	}
	pathExport := dir + "/export/DataExportFromDB.xlsx"

	rsp := &pb_database.ExportTemplateFileRespone{
		PathExport: pathExport,
	}
	return rsp, nil
	// return nil, status.Errorf(codes.Unimplemented, "get list file failed")
}

func (serverDatabase *ServerDatabase) ImportFileExcel(ctx context.Context, res *pb_database.ImportFileExcelResquest) (*pb_database.ImportFileExcelRespone, error) {

	infoFileEx := res.GetImportFileExcel()

	for i := 0; i < len(infoFileEx); i++ {
		// Create an empty map to unmarshal JSON into
		var person []map[string]interface{}

		// Unmarshal the JSON data into the map
		err := json.Unmarshal([]byte(infoFileEx[i].Content), &person)
		if err != nil {
			fmt.Println(err)
			return nil, status.Errorf(codes.Unimplemented, "Unmarshal failed")
		}

		infos := make([]interface{}, len(person))
		for i, s := range person {
			infos[i] = s
		}

		infos = infos[1:]

		err = mongodb.AddManyInfoNotModel(serverDatabase.config, infos, infoFileEx[i].Filename)
		if err != nil {
			log.Error(" Post Data to Mongo Database failed:", err)
			return nil, status.Errorf(codes.Unimplemented, "AddManyInfoNotModel failed")
		}

	}
	return &pb_database.ImportFileExcelRespone{Noti: "Proccesing Done"}, nil
}

func (serverDatabase *ServerDatabase) UpdateStatusProcessingFileExcel(ctx context.Context, res *pb_database.StatusProcessingFileResquest) (*pb_database.StatusProcessingFileRespone, error) {

	status := res.GetStatus()
	idFile := res.GetIdFile()

	err := mongodb.UpdateStatus(serverDatabase.config, idFile, status)

	if err != nil {
		log.Error(" UpdateStatus failed:", err)
	}

	return &pb_database.StatusProcessingFileRespone{Noti: "Done"}, nil
}
