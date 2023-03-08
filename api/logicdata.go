package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"server-test/cache"
	"server-test/database/db"
	"server-test/database/levedb"
	"server-test/database/mongodb"
	"server-test/database/mysqldb"
	"server-test/model"
	"server-test/pb"

	log "github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetData(ctx context.Context, res *pb.DataInfoResquest) (*pb.DataInfoRespone, error) {

	valueCache := cache.GetFromRedis(ctx, "data_test")

	var rsp *pb.DataInfoRespone

	if len(valueCache) == 0 {
		var dataInfo []model.DataInfo
		var err error

		switch server.config.Sever.TypeServer.Name {

		case "server_levedb":

			temp := levedb.GetData(server.config)
			dataInfo = append(dataInfo, temp)

		case "server_mysql":
			var data []model.DataPost
			err := mysqldb.GetData(&data)
			if err != nil {
				return nil, status.Errorf(codes.Unimplemented, "get Data failed")
			}
			for i := 0; i < len(data); i++ {
				dataInfo = append(dataInfo, model.DataInfo{Name: data[i].Name, FullName: data[i].FullName})
			}
		case "server_postgressql":
			dataInfo, err = db.GetData()
			if err != nil {
				return nil, status.Errorf(codes.Unimplemented, "get Data failed")
			}
		case "server_mongodb":
			dataInfo, err = mongodb.GetAllInfo(server.config)
		default:
			return nil, status.Errorf(codes.Unimplemented, "Don't have database")
		}

		var DataInfo []model.DataInfo

		for i := 0; i < len(dataInfo); i++ {
			DataInfo = append(DataInfo, model.DataInfo{Name: dataInfo[i].Name, FullName: dataInfo[i].FullName})
		}
		rsp = &pb.DataInfoRespone{
			Data: ConvertData(DataInfo),
		}
		json1, err := json.Marshal(DataInfo)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("value Database ")
		cache.SetToRedis(ctx, "data_test", json1)
	} else {
		var temp []model.DataInfo
		err := json.Unmarshal([]byte(valueCache), &temp)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("value Cache ")

		rsp = &pb.DataInfoRespone{
			Data: ConvertData(temp),
		}

	}

	return rsp, nil
}

func (server *Server) PostData(ctx context.Context, res *pb.DataPostResqest) (*pb.DataPostRespone, error) {

	Id := int(res.GetId())
	Name := res.GetName()
	FullName := res.GetFullname()

	type Info struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
	}
	// InfoGroup, err := db.PostData(Id, Name, FullName)
	// InfoGroup := "data received : "

	// json1, err := json.Marshal(Info{Id: Id, Name: Name, FullName: FullName})

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// cache.SetToRedis(ctx, "data_test", json1)

	// valueCache := cache.GetFromRedis(ctx, "data_test")
	// fmt.Println("value Cache= ", valueCache)

	// rsp := &pb.DataPostRespone{
	// 	Notice: InfoGroup,
	// }

	// values := cache.GetAllKeys(ctx, "data_test")

	// log.Info("All values : %v \n", values)

	// time.AfterFunc(60*time.Second, func() {
	// 	ctx1 := context.TODO()
	// 	valueCache1 := cache.GetFromRedis(ctx1, "data_test")

	// 	var temp Info
	// 	err := json.Unmarshal([]byte(valueCache1), &temp)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println("value Cache= ", temp)

	// 	info, err := db.PostData(temp.Id, temp.Name, temp.FullName)
	// 	if err != nil {
	// 		log.Error("Post Data to Database failed")
	// 	}

	// 	log.Info(info)
	// })
	var noticeDb string
	switch server.config.Sever.TypeServer.Name {

	case "server_levedb":
		err := levedb.PutData(server.config, Id, Name, FullName)
		if err != nil {
			return nil, status.Errorf(codes.Unimplemented, "Post Data to Leve Database failed")
		}
		noticeDb = "Post Done"

	case "server_mysql":
		err := mysqldb.PostData(Id, Name, FullName)
		if err != nil {
			return nil, status.Errorf(codes.Unimplemented, "Post Data to MySql Database failed")
		}
		noticeDb = "Post Done"

	case "server_postgressql":

		info, err := db.PostData(Id, Name, FullName)
		if err != nil {
			return nil, status.Errorf(codes.Unimplemented, "Post Data to Postges Database failed")
		}
		noticeDb = info

	case "server_mongodb":
		err := mongodb.AddInfo(server.config, Id, Name, FullName)
		if err != nil {
			return nil, status.Errorf(codes.Unimplemented, "Post Data to Mongo Database failed")
		}
		noticeDb = "Post Done"

	default:
		return nil, status.Errorf(codes.Unimplemented, "Don't have database")
	}

	rsp := &pb.DataPostRespone{
		Notice: noticeDb,
	}
	return rsp, nil
}

func (server *Server) UpdateData(ctx context.Context, res *pb.DataUpdateResqest) (*pb.DataUpdateRespone, error) {

	oldName := res.GetOldname()
	newName := res.GetNewname()
	newFullName := res.GetNewfullname()

	InfoGroup, err := db.UpdateData(oldName, newName, newFullName)
	// InfoGroup := "data received : "

	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "Update Data failed")
	}

	rsp := &pb.DataUpdateRespone{
		Notice: InfoGroup,
	}

	return rsp, nil
}

func (server *Server) ExportData(ctx context.Context, res *pb.ExportDataResquest) (*pb.ExportDataRespone, error) {

	var dataInfo []model.DataInfo
	var err error

	switch server.config.Sever.TypeServer.Name {

	case "server_levedb":

		temp := levedb.GetData(server.config)
		dataInfo = append(dataInfo, temp)

	case "server_mysql":
		var data []model.DataPost
		err := mysqldb.GetData(&data)
		if err != nil {
			return nil, status.Errorf(codes.Unimplemented, "get Data failed")
		}
		for i := 0; i < len(data); i++ {
			dataInfo = append(dataInfo, model.DataInfo{Name: data[i].Name, FullName: data[i].FullName})
		}
	case "server_postgressql":
		dataInfo, err = db.GetData()
		if err != nil {
			return nil, status.Errorf(codes.Unimplemented, "get Data failed")
		}
	case "server_mongodb":
		dataInfo, err = mongodb.GetAllInfo(server.config)
	default:
		return nil, status.Errorf(codes.Unimplemented, "Don't have database")
	}

	var DataInfo []model.DataInfo
	var expenseData = [][]interface{}{}

	for i := 0; i < len(dataInfo); i++ {
		DataInfo = append(DataInfo, model.DataInfo{Name: dataInfo[i].Name, FullName: dataInfo[i].FullName})
		temp := []interface{}{dataInfo[i].Name, dataInfo[i].FullName}
		expenseData = append(expenseData, temp)
	}

	f := excelize.NewFile()
	index, _ := f.NewSheet("Sheet1")
	f.SetActiveSheet(index)

	err = f.SetSheetRow("Sheet1", "A1", &[]interface{}{"Name", "FullName"})
	if err != nil {

		log.Error("Error SetSheetRow ", err)
	}
	err = f.SetColWidth("Sheet1", "A", "G", 30)

	if err != nil {

		log.Error("Error SetColWidth ", err)
	}

	startRow := 2
	for i := startRow; i < (len(expenseData) + startRow); i++ {
		err = f.SetSheetRow("Sheet1", fmt.Sprintf("A%d", i), &expenseData[i-2])
		if err != nil {
			log.Error("Error SetSheetRow ", err)
		}
	}

	// Save spreadsheet by the given path.
	if err := f.SaveAs("./export/DataExportFromDB.xlsx"); err != nil {
		fmt.Println(err)
	}

	if err != nil {
		log.Error("Error SaveAs ", err)
	}

	dir, err := filepath.Abs(filepath.Dir("DataExportFromDB.xlsx"))
	if err != nil {
		log.Fatal(err)
	}
	pathExport := dir + "/export/DataExportFromDB.xlsx"

	rsp := &pb.ExportDataRespone{
		PathExport: pathExport,
	}

	return rsp, nil
}

//Version using GRPC-Gateway body using JSON

// func (server *Server) ImportData_V1(ctx context.Context, res *pb.ImportDataResquest) (*pb.ImportDataRespone, error) {

// 	var noticeDb string

// 	// pathImport := res.GetPathimport()
// 	pathImport := "/home/nhatnt/nhatnt/probationary-project/server-test/importfile/DataImportToDB.xlsx"
// 	file, err := excelize.OpenFile(pathImport)
// 	if err != nil {
// 		return nil, status.Errorf(codes.Unimplemented, "Open file err", err)
// 	}
// 	defer func() {
// 		// Close the spreadsheet.
// 		if err := file.Close(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()

// 	// for index, name := range file.GetSheetMap() {
// 	// 	fmt.Println(index, name)
// 	// }

// 	data1, err := file.GetRows("Sheet1")

// 	if err != nil {
// 		return nil, status.Errorf(codes.Unimplemented, "GetRows failed")
// 	}

// 	var data []map[string]string
// 	for _, row := range data1 {
// 		item := make(map[string]string)
// 		for i, colCell := range row {
// 			temp, _ := excelize.ColumnNumberToName(i + 1)

// 			temp1, _ := file.GetCellValue("Sheet1", temp+"1")

// 			item[temp1] = colCell
// 		}
// 		data = append(data, item)
// 	}

// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// fmt.Println("jsonData = ", string(jsonData))

// 	// Create an empty map to unmarshal JSON into
// 	var person []map[string]interface{}

// 	// Unmarshal the JSON data into the map
// 	err1 := json.Unmarshal([]byte(jsonData), &person)
// 	if err != nil {
// 		fmt.Println(err1)
// 	}

// 	infos := make([]interface{}, len(person))
// 	for i, s := range person {
// 		infos[i] = s
// 	}

// 	infos = infos[1:]

// 	var datasAdd []model.DataPost

// 	for i := 0; i < len(data); i++ {
// 		idInt, _ := strconv.Atoi(data[i][0])
// 		datasAdd = append(datasAdd, model.DataPost{Id: idInt, Name: data[i][1], FullName: data[i][2]})

// 	}

// 	// fmt.Println(infos)
// 	err = mongodb.AddManyInfoNotModel(server.config, infos)
// 	if err != nil {
// 		return nil, status.Errorf(codes.Unimplemented, "Post Data to Mongo Database failed")
// 	}
// 	noticeDb = "Post Done"

// 	rsp := &pb.ImportDataRespone{
// 		Notice: noticeDb,
// 	}
// 	return rsp, nil
// }

type InfoFile struct {
	name    string
	content []byte
}

// Handle data with form-data
func ImportDataWithHttp(w http.ResponseWriter, r *http.Request) ([]InfoFile, error) {
	file_ex, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to retrieve file from form data", http.StatusBadRequest)
	}
	defer file_ex.Close()

	// fmt.Println("file", file)

	content, err := ioutil.ReadAll(file_ex)

	xlsx, err := excelize.OpenReader(bytes.NewReader(content))
	if err != nil {
		http.Error(w, "Failed to open Excel file", http.StatusBadRequest)
		return nil, err
	}

	var info_file []InfoFile
	for _, name := range xlsx.GetSheetMap() {
		// fmt.Println(index, name)
		// nameSheet = append(nameSheet, name)
		dataRows, err := xlsx.GetRows(name)
		if err != nil {
			http.Error(w, "GetRows failed", http.StatusBadRequest)
			return nil, err
		}

		var data []map[string]string
		for _, row := range dataRows {
			item := make(map[string]string)
			for i, colCell := range row {
				temp, _ := excelize.ColumnNumberToName(i + 1)

				temp1, _ := xlsx.GetCellValue(name, temp+"1")

				item[temp1] = colCell
			}
			data = append(data, item)
		}
		jsonData, err := json.Marshal(data)

		if err != nil {
			return nil, err
		}

		info_file = append(info_file, InfoFile{name: name, content: jsonData})
	}

	return info_file, nil
}

// Version using GRPC
func (server *Server) ImportData(ctx context.Context, res *pb.ImportDataResquest) (*pb.ImportDataRespone, error) {

	var noticeDb string

	data_ex := res.GetData()
	name := res.GetName()

	// Create an empty map to unmarshal JSON into
	var person []map[string]interface{}

	// Unmarshal the JSON data into the map
	err := json.Unmarshal([]byte(data_ex), &person)
	if err != nil {
		fmt.Println(err)
	}

	infos := make([]interface{}, len(person))
	for i, s := range person {
		infos[i] = s
	}

	infos = infos[1:]

	err = mongodb.AddManyInfoNotModel(server.config, infos, name)
	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "Post Data to Mongo Database failed")
	}
	noticeDb = "Post Done"

	rsp := &pb.ImportDataRespone{
		Notice: noticeDb,
	}
	return rsp, nil
}
