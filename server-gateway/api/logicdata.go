package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server-test/server-gateway/pb"
	"strings"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"server-test/server-gateway/pb"
// 	"strings"

// 	log "github.com/sirupsen/logrus"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/metadata"
// 	"google.golang.org/grpc/status"
// )

// // import (
// // 	"context"
// // 	"encoding/json"
// // 	"fmt"
// // 	"io/ioutil"
// // 	"net/http"
// // 	"server-test/database/mongodb"
// // 	"server-test/logs"
// // 	"server-test/pb"

// // 	log "github.com/sirupsen/logrus"
// // 	"google.golang.org/grpc/codes"
// // 	"google.golang.org/grpc/metadata"
// // 	"google.golang.org/grpc/status"
// // )

// // func (server *Server) GetData(ctx context.Context, res *pb.DataInfoResquest) (*pb.DataInfoRespone, error) {

// // 	// logs.Logger.Info("GetData - API call GetData ")

// // 	// valueCache := cache.GetFromRedis(ctx, "data_test")

// // 	// var rsp *pb.DataInfoRespone

// // 	// if len(valueCache) == 0 {
// // 	// 	var dataInfo []model.DataInfo
// // 	// 	var err error

// // 	// 	switch server.config.Sever.TypeServer.Name {

// // 	// 	case "server_levedb":

// // 	// 		temp := levedb.GetData(server.config)
// // 	// 		dataInfo = append(dataInfo, temp)

// // 	// 	case "server_mysql":
// // 	// 		var data []model.DataPost
// // 	// 		err := mysqldb.GetData(&data)
// // 	// 		if err != nil {
// // 	// 			logs.Logger.Error("GetData - Get Data server_mysql failed: ", err)
// // 	// 			return nil, status.Errorf(codes.Unimplemented, "get Data failed")
// // 	// 		}
// // 	// 		for i := 0; i < len(data); i++ {
// // 	// 			dataInfo = append(dataInfo, model.DataInfo{Name: data[i].Name, FullName: data[i].FullName})
// // 	// 		}
// // 	// 	case "server_postgressql":
// // 	// 		dataInfo, err = db.GetData()
// // 	// 		if err != nil {
// // 	// 			logs.Logger.Error("GetData - Get Data server_postgressql failed: ", err)
// // 	// 			return nil, status.Errorf(codes.Unimplemented, "get Data failed")
// // 	// 		}
// // 	// 	case "server_mongodb":
// // 	// 		dataInfo, err = mongodb.GetAllInfo(server.config)
// // 	// 		if err != nil {
// // 	// 			logs.Logger.Error("GetData - Get Data server_mongodb failed: ", err)
// // 	// 			return nil, status.Errorf(codes.Unimplemented, "get Data failed")
// // 	// 		}
// // 	// 	default:
// // 	// 		logs.Logger.Error("GetData - Don't have database")
// // 	// 		return nil, status.Errorf(codes.Unimplemented, "Don't have database")
// // 	// 	}

// // 	// 	var DataInfo []model.DataInfo

// // 	// 	for i := 0; i < len(dataInfo); i++ {
// // 	// 		DataInfo = append(DataInfo, model.DataInfo{Name: dataInfo[i].Name, FullName: dataInfo[i].FullName})
// // 	// 	}
// // 	// 	rsp = &pb.DataInfoRespone{
// // 	// 		Data: ConvertData(DataInfo),
// // 	// 	}
// // 	// 	json1, err := json.Marshal(DataInfo)
// // 	// 	if err != nil {
// // 	// 		logs.Logger.Error("GetData - Error Marshal: ", err)
// // 	// 		fmt.Println(err)
// // 	// 	}

// // 	// 	fmt.Println("value Database ")
// // 	// 	logs.Logger.Info("GetData - Value Database ")
// // 	// 	cache.SetToRedis(ctx, "data_test", json1)
// // 	// } else {
// // 	// 	var temp []model.DataInfo
// // 	// 	err := json.Unmarshal([]byte(valueCache), &temp)
// // 	// 	if err != nil {
// // 	// 		logs.Logger.Error("GetData - Error Unmarshal: ", err)
// // 	// 		fmt.Println(err)
// // 	// 	}
// // 	// 	fmt.Println("value Cache ")
// // 	// 	logs.Logger.Info("GetData - Value Cache ")

// // 	// 	rsp = &pb.DataInfoRespone{
// // 	// 		Data: ConvertData(temp),
// // 	// 	}

// // 	// }

// // 	// return rsp, nil
// // 	return nil, status.Errorf(codes.Unimplemented, "get Data failed")
// // }

// // // func (server *Server) PostData(ctx context.Context, res *pb.DataPostResqest) (*pb.DataPostRespone, error) {

// // // 	logs.Logger.Info("PostData - API call PostData ")

// // // 	Id := int(res.GetId())
// // // 	Name := res.GetName()
// // // 	FullName := res.GetFullname()

// // // 	type Info struct {
// // // 		Id       int    `json:"id"`
// // // 		Name     string `json:"name"`
// // // 		FullName string `json:"full_name"`
// // // 	}
// // // 	// InfoGroup, err := db.PostData(Id, Name, FullName)
// // // 	// InfoGroup := "data received : "

// // // 	// json1, err := json.Marshal(Info{Id: Id, Name: Name, FullName: FullName})

// // // 	// if err != nil {
// // // 	// 	fmt.Println(err)
// // // 	// }
// // // 	// cache.SetToRedis(ctx, "data_test", json1)

// // // 	// valueCache := cache.GetFromRedis(ctx, "data_test")
// // // 	// fmt.Println("value Cache= ", valueCache)

// // // 	// rsp := &pb.DataPostRespone{
// // // 	// 	Notice: InfoGroup,
// // // 	// }

// // // 	// values := cache.GetAllKeys(ctx, "data_test")

// // // 	// log.Info("All values : %v \n", values)

// // // 	// time.AfterFunc(60*time.Second, func() {
// // // 	// 	ctx1 := context.TODO()
// // // 	// 	valueCache1 := cache.GetFromRedis(ctx1, "data_test")

// // // 	// 	var temp Info
// // // 	// 	err := json.Unmarshal([]byte(valueCache1), &temp)
// // // 	// 	if err != nil {
// // // 	// 		fmt.Println(err)
// // // 	// 	}
// // // 	// 	fmt.Println("value Cache= ", temp)

// // // 	// 	info, err := db.PostData(temp.Id, temp.Name, temp.FullName)
// // // 	// 	if err != nil {
// // // 	// 		log.Error("Post Data to Database failed")
// // // 	// 	}

// // // 	// 	log.Info(info)
// // // 	// })
// // // 	var noticeDb string
// // // 	switch server.config.Sever.TypeServer.Name {

// // // 	case "server_levedb":
// // // 		err := levedb.PutData(server.config, Id, Name, FullName)
// // // 		if err != nil {
// // // 			logs.Logger.Error("PostData - Post Data to Leve Database failed: ", err)
// // // 			return nil, status.Errorf(codes.Unimplemented, "Post Data to Leve Database failed")
// // // 		}
// // // 		noticeDb = "Post Done"

// // // 	case "server_mysql":
// // // 		err := mysqldb.PostData(Id, Name, FullName)
// // // 		if err != nil {
// // // 			logs.Logger.Error("PostData - Post Data to MySql Database failed: ", err)
// // // 			return nil, status.Errorf(codes.Unimplemented, "Post Data to MySql Database failed")
// // // 		}
// // // 		noticeDb = "Post Done"

// // // 	case "server_postgressql":

// // // 		info, err := db.PostData(Id, Name, FullName)
// // // 		if err != nil {
// // // 			logs.Logger.Error("PostData - Post Data to Postges Database failed: ", err)
// // // 			return nil, status.Errorf(codes.Unimplemented, "Post Data to Postges Database failed")
// // // 		}
// // // 		noticeDb = info

// // // 	case "server_mongodb":
// // // 		err := mongodb.AddInfo(server.config, Id, Name, FullName)
// // // 		if err != nil {
// // // 			logs.Logger.Error("PostData - Post Data to Mongo Database failed: ", err)
// // // 			return nil, status.Errorf(codes.Unimplemented, "Post Data to Mongo Database failed")
// // // 		}
// // // 		noticeDb = "Post Done"

// // // 	default:
// // // 		logs.Logger.Error("PostData - Don't have database")
// // // 		return nil, status.Errorf(codes.Unimplemented, "Don't have database")
// // // 	}

// // // 	rsp := &pb.DataPostRespone{
// // // 		Notice: noticeDb,
// // // 	}
// // // 	return rsp, nil
// // // }

// // // func (server *Server) UpdateData(ctx context.Context, res *pb.DataUpdateResqest) (*pb.DataUpdateRespone, error) {

// // // 	logs.Logger.Info("UpdateData - API call UpdateData ")

// // // 	oldName := res.GetOldname()
// // // 	newName := res.GetNewname()
// // // 	newFullName := res.GetNewfullname()

// // // 	InfoGroup, err := db.UpdateData(oldName, newName, newFullName)
// // // 	// InfoGroup := "data received : "

// // // 	if err != nil {
// // // 		logs.Logger.Error("UpdateData - Update Data failed error: ", err)
// // // 		return nil, status.Errorf(codes.Unimplemented, "Update Data failed")
// // // 	}

// // // 	rsp := &pb.DataUpdateRespone{
// // // 		Notice: InfoGroup,
// // // 	}

// // // 	return rsp, nil
// // // }

// // func (server *Server) ExportData(ctx context.Context, res *pb.ExportDataResquest) (*pb.ExportDataRespone, error) {

// // 	logs.Logger.Info("ExportData - API call ExportData ")

// // 	var dataInfo []model.DataInfo
// // 	var err error

// // 	switch server.config.Sever.TypeServer.Name {

// // 	case "server_levedb":

// // 		temp := levedb.GetData(server.config)
// // 		dataInfo = append(dataInfo, temp)

// // 	case "server_mysql":
// // 		var data []model.DataPost
// // 		err := mysqldb.GetData(&data)
// // 		if err != nil {
// // 			logs.Logger.Error("ExportData: server_mysql get Data failed error:", err)
// // 			return nil, status.Errorf(codes.Unimplemented, "get Data failed")
// // 		}
// // 		for i := 0; i < len(data); i++ {
// // 			dataInfo = append(dataInfo, model.DataInfo{Name: data[i].Name, FullName: data[i].FullName})
// // 		}
// // 	case "server_postgressql":
// // 		dataInfo, err = db.GetData()
// // 		if err != nil {
// // 			logs.Logger.Error("ExportData: server_postgressql get Data failed error:", err)
// // 			return nil, status.Errorf(codes.Unimplemented, "get Data failed")
// // 		}
// // 	case "server_mongodb":
// // 		dataInfo, err = mongodb.GetAllInfo(server.config)
// // 		if err != nil {
// // 			logs.Logger.Error("ExportData: server_mongodb get Data failed error:", err)
// // 			return nil, status.Errorf(codes.Unimplemented, "get Data failed")
// // 		}
// // 	default:
// // 		logs.Logger.Error("ExportData: Don't have database")
// // 		return nil, status.Errorf(codes.Unimplemented, "Don't have database")
// // 	}

// // 	var DataInfo []model.DataInfo
// // 	var expenseData = [][]interface{}{}

// // 	for i := 0; i < len(dataInfo); i++ {
// // 		DataInfo = append(DataInfo, model.DataInfo{Name: dataInfo[i].Name, FullName: dataInfo[i].FullName})
// // 		temp := []interface{}{dataInfo[i].Name, dataInfo[i].FullName}
// // 		expenseData = append(expenseData, temp)
// // 	}

// // 	f := excelize.NewFile()
// // 	index, _ := f.NewSheet("Sheet1")
// // 	f.SetActiveSheet(index)

// // 	err = f.SetSheetRow("Sheet1", "A1", &[]interface{}{"Name", "FullName"})
// // 	if err != nil {

// // 		logs.Logger.Error("ExportData: Error SetSheetRow: ", err)
// // 	}
// // 	err = f.SetColWidth("Sheet1", "A", "G", 30)

// // 	if err != nil {

// // 		logs.Logger.Error("ExportData: Error SetColWidth: ", err)
// // 	}

// // 	startRow := 2
// // 	for i := startRow; i < (len(expenseData) + startRow); i++ {
// // 		err = f.SetSheetRow("Sheet1", fmt.Sprintf("A%d", i), &expenseData[i-2])
// // 		if err != nil {
// // 			logs.Logger.Error("ExportData: Error SetSheetRow: ", err)
// // 		}
// // 	}

// // 	// Save spreadsheet by the given path.
// // 	if err := f.SaveAs("./export/DataExportFromDB.xlsx"); err != nil {
// // 		fmt.Println(err)
// // 	}

// // 	if err != nil {
// // 		logs.Logger.Error("ExportData: Error SaveAs: ", err)
// // 	}

// // 	dir, err := filepath.Abs(filepath.Dir("DataExportFromDB.xlsx"))
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // 	pathExport := dir + "/export/DataExportFromDB.xlsx"

// // 	rsp := &pb.ExportDataRespone{
// // 		PathExport: pathExport,
// // 	}

// // 	return rsp, nil
// // }

// // //Version using GRPC-Gateway body using JSON

// // // func (server *Server) ImportData_V1(ctx context.Context, res *pb.ImportDataResquest) (*pb.ImportDataRespone, error) {

// // // 	var noticeDb string

// // // 	// pathImport := res.GetPathimport()
// // // 	pathImport := "/home/nhatnt/nhatnt/probationary-project/server-test/importfile/DataImportToDB.xlsx"
// // // 	file, err := excelize.OpenFile(pathImport)
// // // 	if err != nil {
// // // 		return nil, status.Errorf(codes.Unimplemented, "Open file err", err)
// // // 	}
// // // 	defer func() {
// // // 		// Close the spreadsheet.
// // // 		if err := file.Close(); err != nil {
// // // 			fmt.Println(err)
// // // 		}
// // // 	}()

// // // 	// for index, name := range file.GetSheetMap() {
// // // 	// 	fmt.Println(index, name)
// // // 	// }

// // // 	data1, err := file.GetRows("Sheet1")

// // // 	if err != nil {
// // // 		return nil, status.Errorf(codes.Unimplemented, "GetRows failed")
// // // 	}

// // // 	var data []map[string]string
// // // 	for _, row := range data1 {
// // // 		item := make(map[string]string)
// // // 		for i, colCell := range row {
// // // 			temp, _ := excelize.ColumnNumberToName(i + 1)

// // // 			temp1, _ := file.GetCellValue("Sheet1", temp+"1")

// // // 			item[temp1] = colCell
// // // 		}
// // // 		data = append(data, item)
// // // 	}

// // // 	jsonData, err := json.Marshal(data)
// // // 	if err != nil {
// // // 		fmt.Println(err)
// // // 	}

// // // 	// fmt.Println("jsonData = ", string(jsonData))

// // // 	// Create an empty map to unmarshal JSON into
// // // 	var person []map[string]interface{}

// // // 	// Unmarshal the JSON data into the map
// // // 	err1 := json.Unmarshal([]byte(jsonData), &person)
// // // 	if err != nil {
// // // 		fmt.Println(err1)
// // // 	}

// // // 	infos := make([]interface{}, len(person))
// // // 	for i, s := range person {
// // // 		infos[i] = s
// // // 	}

// // // 	infos = infos[1:]

// // // 	var datasAdd []model.DataPost

// // // 	for i := 0; i < len(data); i++ {
// // // 		idInt, _ := strconv.Atoi(data[i][0])
// // // 		datasAdd = append(datasAdd, model.DataPost{Id: idInt, Name: data[i][1], FullName: data[i][2]})

// // // 	}

// // // 	// fmt.Println(infos)
// // // 	err = mongodb.AddManyInfoNotModel(server.config, infos)
// // // 	if err != nil {
// // // 		return nil, status.Errorf(codes.Unimplemented, "Post Data to Mongo Database failed")
// // // 	}
// // // 	noticeDb = "Post Done"

// // // 	rsp := &pb.ImportDataRespone{
// // // 		Notice: noticeDb,
// // // 	}
// // // 	return rsp, nil
// // // }

// // type InfoFile struct {
// // 	name    string
// // 	content []byte
// // }

// Handle data with form-data
func (server *Server) ImportDataWithHttp(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("token")

	resp, err := server.clientAuthen.AuthenTokenClient(context.Background(), token)

	if err != nil {
		http.Error(w, "Authen token failed", http.StatusUnauthorized)
	}

	fmt.Println(resp.Iduser)

	file_ex, a, err := r.FormFile("file")

	if err != nil {
		// logs.Logger.Error("ImportDataWithHttp: Failed to retrieve file from form data")
		http.Error(w, "Failed to retrieve file from form data", http.StatusBadRequest)
		return
	}
	defer file_ex.Close()

	if a.Header.Get("Content-Type") != "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
		// logs.Logger.Error("ImportDataWithHttp: Format file error")
		http.Error(w, "Format file error (xlsx)", http.StatusBadRequest)
		return
	}

	fmt.Println("Gateway server")

	content, err := ioutil.ReadAll(file_ex)

	idFileUpLoad, err := server.clientStogare.UploadFile(context.Background(), a.Filename, a.Header.Get("Content-Type"), resp.Iduser, a.Size, content)

	// fmt.Println("check id File: ", idFile)
	if r.Header.Get("run") == "run" {
		go func(idFile, fileName string, fileContent []byte) {
			_, err := server.clientProcessing.ProcessingDataClient(context.Background(), idFile, fileName, fileContent)

			if err != nil {
				http.Error(w, "ProcessingDataClient Failed ", http.StatusBadRequest)
			}
		}(idFileUpLoad.Link, a.Filename, content)
	}

	// if a.Size > 1024*1024 {
	// 	logs.Logger.Error("ImportDataWithHttp: File too lagre")
	// 	http.Error(w, "File too lagre (<=1 Mb)", http.StatusBadRequest)
	// 	return
	// }

	// xlsx, err := excelize.OpenReader(bytes.NewReader(content))
	// xlsx, err := excelize.OpenReader(file_ex)

	// if err != nil {
	// 	logs.Logger.Error("ImportDataWithHttp: Failed to open Excel file ", err)
	// 	http.Error(w, "Failed to open Excel file", http.StatusBadRequest)
	// 	return
	// }

	// var info_file []InfoFile
	// for _, name := range xlsx.GetSheetMap() {
	// 	// fmt.Println(index, name)
	// 	// nameSheet = append(nameSheet, name)
	// 	dataRows, err := xlsx.GetRows(name)
	// 	if err != nil {
	// 		logs.Logger.Error("ImportDataWithHttp: GetRows failed")
	// 		http.Error(w, "GetRows failed", http.StatusBadRequest)
	// 		return
	// 	}

	// 	var data []map[string]string
	// 	for _, row := range dataRows {
	// 		item := make(map[string]string)
	// 		for i, colCell := range row {
	// 			temp, _ := excelize.ColumnNumberToName(i + 1)

	// 			temp1, _ := xlsx.GetCellValue(name, temp+"1")

	// 			item[temp1] = colCell
	// 		}
	// 		data = append(data, item)
	// 	}
	// 	jsonData, err := json.Marshal(data)

	// 	if err != nil {
	// 		logs.Logger.Error("ImportDataWithHttp: Masrhal Data error:", err)
	// 		return
	// 	}

	// 	info_file = append(info_file, InfoFile{name: name, content: jsonData})
	// }

	// for i := 0; i < len(info_file); i++ {
	// 	// Create an empty map to unmarshal JSON into
	// 	var person []map[string]interface{}

	// 	// Unmarshal the JSON data into the map
	// 	err = json.Unmarshal([]byte(info_file[i].content), &person)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	infos := make([]interface{}, len(person))
	// 	for i, s := range person {
	// 		infos[i] = s
	// 	}

	// 	infos = infos[1:]

	// 	err = mongodb.AddManyInfoNotModel(server.config, infos, info_file[i].name)
	// 	if err != nil {
	// 		logs.Logger.Error("ImportDataWithHttp: Post Data to Mongo Database failed:", err)
	// 	}

	// }

	type Respone struct {
		Notice string
	}

	noticeDb := Respone{Notice: "Upload Done"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(noticeDb)
}

// // // Version using GRPC
// // // func (server *Server) ImportData(ctx context.Context, res *pb.ImportDataResquest) (*pb.ImportDataRespone, error) {

// // // 	logs.Logger.Info("ImportData: API call ImportData")
// // // 	var noticeDb string

// // // 	data_ex := res.GetData()
// // // 	name := res.GetName()

// // // 	// Create an empty map to unmarshal JSON into
// // // 	var person []map[string]interface{}

// // // 	// Unmarshal the JSON data into the map
// // // 	err := json.Unmarshal([]byte(data_ex), &person)
// // // 	if err != nil {
// // // 		fmt.Println(err)
// // // 	}

// // // 	infos := make([]interface{}, len(person))
// // // 	for i, s := range person {
// // // 		infos[i] = s
// // // 	}

// // // 	infos = infos[1:]

// // // 	err = mongodb.AddManyInfoNotModel(server.config, infos, name)
// // // 	if err != nil {
// // // 		return nil, status.Errorf(codes.Unimplemented, "Post Data to Mongo Database failed")
// // // 	}
// // // 	noticeDb = "Post Done"

// // // 	rsp := &pb.ImportDataRespone{
// // // 		Notice: noticeDb,
// // // 	}
// // // 	return rsp, nil
// // // }

// // func (server *Server) TestData1(ctx context.Context, res *pb.TestResquest) (*pb.TestRespone, error) {

// // 	temp := res.GetTestdata()

// // 	temp1, err := server.clientStogare.TestData(ctx, temp)

// // 	fmt.Println(err)

// // 	return &pb.TestRespone{Notice: temp1}, nil
// // }

func (server *Server) SignUp(ctx context.Context, res *pb.SignUpResquest) (*pb.SignUpRespone, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	log.Info("nhatnt md: ", md)
	temp := md["username"][0]

	log.Info("nhatnt UserName: ", temp)

	log.Info("nhatnt Password: ", md["password"])

	resp, err := server.clientAuthen.SignUp(ctx, md["username"][0], md["password"][0], md["lastname"][0], md["firstname"][0])

	log.Info("nhatnt: ", err)
	log.Info("nhatntq: ", ok)
	return &pb.SignUpRespone{Notice: resp}, nil
}

func (server *Server) LogInAcc(ctx context.Context, res *pb.SignInResquest) (*pb.SignInRespone, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	log.Info("nhatnt md: ", md)
	temp := md["username"][0]

	log.Info("nhatnt UserName: ", temp)

	log.Info("nhatnt Password: ", md["password"])

	resp, err := server.clientAuthen.SignInClient(ctx, md["username"][0], md["password"][0])

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "User or password is incorrect")
	}

	log.Info("nhatntq: ", ok)

	infoUser := &pb.UserAccInfo{

		Username: resp.Userinfo.Username,
		Fullname: resp.Userinfo.Firstname + " " + resp.Userinfo.Lastname,
	}

	info := &pb.SignInRespone{
		Useraccinfo: infoUser,
		Token:       resp.Token,
	}
	return info, nil
}

func (server *Server) GetFileUploadInfo(ctx context.Context, res *pb.FileUploadInfoResquest) (*pb.FileUploadInfoRespone, error) {

	md, _ := metadata.FromIncomingContext(ctx)
	log.Info("nhatnt md: ", md)

	resp, err := server.clientAuthen.AuthenTokenClient(context.Background(), md["token"][0])

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Authen token failed")
	}
	respDatabase, err := server.clientStogare.GetUploadFileInfoClient(context.Background(), resp.Iduser)

	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "Get list file failed")
	}

	fmt.Println(respDatabase)

	var temp []*pb.FileUploadInfo

	for i := 0; i < len(respDatabase.Fileinfo); i++ {
		temp = append(temp, &pb.FileUploadInfo{
			Filename:     respDatabase.Fileinfo[0].Filename,
			Filetype:     respDatabase.Fileinfo[0].Typefile,
			Sizefile:     int64(respDatabase.Fileinfo[0].Size),
			Link:         respDatabase.Fileinfo[0].Link,
			Timecreateat: respDatabase.Fileinfo[0].Timecreateat,
		})
	}

	return &pb.FileUploadInfoRespone{Fileinfo: temp}, nil
}

func (server *Server) ExportData(ctx context.Context, res *pb.ExportDataResquest) (*pb.ExportDataRespone, error) {

	fmt.Println("ExportData - API call ExportData ")
	md, _ := metadata.FromIncomingContext(ctx)
	log.Info("nhatnt md: ", md)

	_, err := server.clientAuthen.AuthenTokenClient(context.Background(), md["token"][0])

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Authen token failed")
	}
	fmt.Println("trongnhat test", server.clientProcessing)
	resp, err := server.clientProcessing.ExportFileTemplateExcelClient(ctx, md["template"][0])

	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "ExportData failed")
	}

	return &pb.ExportDataRespone{PathExport: resp.PathExport}, nil
}

func (server *Server) DowloadLinkWithHttp(w http.ResponseWriter, r *http.Request) {

	values := r.URL.Query()

	dir := values.Get("dir")
	idFile := values.Get("idfile")

	if len(dir) != 0 {

		resp, err := server.clientProcessing.DownloadFileClient(context.Background(), dir)

		if err != nil {
			fmt.Println(err)
			http.Error(w, "DownloadFileClient failded", http.StatusBadRequest)
			return
		}
		name := strings.Split(resp.NameFile, "/")
		fmt.Println(name)

		w.Header().Set("Content-Disposition", "attachment; filename="+name[len(name)-1])
		w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		w.Write(resp.ContentFile)
	} else {
		resp, err := server.clientStogare.DownloadFileClient(context.Background(), idFile)

		if err != nil {
			fmt.Println(err)
			http.Error(w, "DownloadFileClient failded", http.StatusBadRequest)
			return
		}
		name := strings.Split(resp.Name, "/")
		fmt.Println(name)

		w.Header().Set("Content-Disposition", "attachment; filename="+name[len(name)-1])
		w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		w.Write(resp.Content)
	}

	// http.ServeContent(w, r, "DataImportToDB.xlsx", currentTime, file)
}
