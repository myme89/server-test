package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server-test/server-gateway/database/mongodb"
	"server-test/server-gateway/pb"
	"server-test/server-gateway/ultils"
	"server-test/server-storage/pb_storage"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// func (server *Server) GetData(ctx context.Context, res *pb.DataInfoResquest) (*pb.DataInfoRespone, error) {

// logs.Logger.Info("GetData - API call GetData ")

// valueCache := cache.GetFromRedis(ctx, "data_test")

// var rsp *pb.DataInfoRespone

// if len(valueCache) == 0 {
// 	var dataInfo []model.DataInfo
// 	var err error

// 	switch server.config.Sever.TypeServer.Name {

// 	case "server_levedb":

// 		temp := levedb.GetData(server.config)
// 		dataInfo = append(dataInfo, temp)

// 	case "server_mysql":
// 		var data []model.DataPost
// 		err := mysqldb.GetData(&data)
// 		if err != nil {
// 			logs.Logger.Error("GetData - Get Data server_mysql failed: ", err)
// 			return nil, status.Errorf(codes.Unimplemented, "get Data failed")
// 		}
// 		for i := 0; i < len(data); i++ {
// 			dataInfo = append(dataInfo, model.DataInfo{Name: data[i].Name, FullName: data[i].FullName})
// 		}
// 	case "server_postgressql":
// 		dataInfo, err = db.GetData()
// 		if err != nil {
// 			logs.Logger.Error("GetData - Get Data server_postgressql failed: ", err)
// 			return nil, status.Errorf(codes.Unimplemented, "get Data failed")
// 		}
// 	case "server_mongodb":
// 		dataInfo, err = mongodb.GetAllInfo(server.config)
// 		if err != nil {
// 			logs.Logger.Error("GetData - Get Data server_mongodb failed: ", err)
// 			return nil, status.Errorf(codes.Unimplemented, "get Data failed")
// 		}
// 	default:
// 		logs.Logger.Error("GetData - Don't have database")
// 		return nil, status.Errorf(codes.Unimplemented, "Don't have database")
// 	}

// 	var DataInfo []model.DataInfo

// 	for i := 0; i < len(dataInfo); i++ {
// 		DataInfo = append(DataInfo, model.DataInfo{Name: dataInfo[i].Name, FullName: dataInfo[i].FullName})
// 	}
// 	rsp = &pb.DataInfoRespone{
// 		Data: ConvertData(DataInfo),
// 	}
// 	json1, err := json.Marshal(DataInfo)
// 	if err != nil {
// 		logs.Logger.Error("GetData - Error Marshal: ", err)
// 		fmt.Println(err)
// 	}

// 	fmt.Println("value Database ")
// 	logs.Logger.Info("GetData - Value Database ")
// 	cache.SetToRedis(ctx, "data_test", json1)
// } else {
// 	var temp []model.DataInfo
// 	err := json.Unmarshal([]byte(valueCache), &temp)
// 	if err != nil {
// 		logs.Logger.Error("GetData - Error Unmarshal: ", err)
// 		fmt.Println(err)
// 	}
// 	fmt.Println("value Cache ")
// 	logs.Logger.Info("GetData - Value Cache ")

// 	rsp = &pb.DataInfoRespone{
// 		Data: ConvertData(temp),
// 	}

// }

// return rsp, nil
// 	return nil, status.Errorf(codes.Unimplemented, "get Data failed")
// }

// func (server *Server) PostData(ctx context.Context, res *pb.DataPostResqest) (*pb.DataPostRespone, error) {

// 	logs.Logger.Info("PostData - API call PostData ")

// 	Id := int(res.GetId())
// 	Name := res.GetName()
// 	FullName := res.GetFullname()

// 	type Info struct {
// 		Id       int    `json:"id"`
// 		Name     string `json:"name"`
// 		FullName string `json:"full_name"`
// 	}
// 	// InfoGroup, err := db.PostData(Id, Name, FullName)
// 	// InfoGroup := "data received : "

// 	// json1, err := json.Marshal(Info{Id: Id, Name: Name, FullName: FullName})

// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	// cache.SetToRedis(ctx, "data_test", json1)

// 	// valueCache := cache.GetFromRedis(ctx, "data_test")
// 	// fmt.Println("value Cache= ", valueCache)

// 	// rsp := &pb.DataPostRespone{
// 	// 	Notice: InfoGroup,
// 	// }

// 	// values := cache.GetAllKeys(ctx, "data_test")

// 	// log.Info("All values : %v \n", values)

// 	// time.AfterFunc(60*time.Second, func() {
// 	// 	ctx1 := context.TODO()
// 	// 	valueCache1 := cache.GetFromRedis(ctx1, "data_test")

// 	// 	var temp Info
// 	// 	err := json.Unmarshal([]byte(valueCache1), &temp)
// 	// 	if err != nil {
// 	// 		fmt.Println(err)
// 	// 	}
// 	// 	fmt.Println("value Cache= ", temp)

// 	// 	info, err := db.PostData(temp.Id, temp.Name, temp.FullName)
// 	// 	if err != nil {
// 	// 		log.Error("Post Data to Database failed")
// 	// 	}

// 	// 	log.Info(info)
// 	// })
// 	var noticeDb string
// 	switch server.config.Sever.TypeServer.Name {

// 	case "server_levedb":
// 		err := levedb.PutData(server.config, Id, Name, FullName)
// 		if err != nil {
// 			logs.Logger.Error("PostData - Post Data to Leve Database failed: ", err)
// 			return nil, status.Errorf(codes.Unimplemented, "Post Data to Leve Database failed")
// 		}
// 		noticeDb = "Post Done"

// 	case "server_mysql":
// 		err := mysqldb.PostData(Id, Name, FullName)
// 		if err != nil {
// 			logs.Logger.Error("PostData - Post Data to MySql Database failed: ", err)
// 			return nil, status.Errorf(codes.Unimplemented, "Post Data to MySql Database failed")
// 		}
// 		noticeDb = "Post Done"

// 	case "server_postgressql":

// 		info, err := db.PostData(Id, Name, FullName)
// 		if err != nil {
// 			logs.Logger.Error("PostData - Post Data to Postges Database failed: ", err)
// 			return nil, status.Errorf(codes.Unimplemented, "Post Data to Postges Database failed")
// 		}
// 		noticeDb = info

// 	case "server_mongodb":
// 		err := mongodb.AddInfo(server.config, Id, Name, FullName)
// 		if err != nil {
// 			logs.Logger.Error("PostData - Post Data to Mongo Database failed: ", err)
// 			return nil, status.Errorf(codes.Unimplemented, "Post Data to Mongo Database failed")
// 		}
// 		noticeDb = "Post Done"

// 	default:
// 		logs.Logger.Error("PostData - Don't have database")
// 		return nil, status.Errorf(codes.Unimplemented, "Don't have database")
// 	}

// 	rsp := &pb.DataPostRespone{
// 		Notice: noticeDb,
// 	}
// 	return rsp, nil
// }

// func (server *Server) UpdateData(ctx context.Context, res *pb.DataUpdateResqest) (*pb.DataUpdateRespone, error) {

// 	logs.Logger.Info("UpdateData - API call UpdateData ")

// 	oldName := res.GetOldname()
// 	newName := res.GetNewname()
// 	newFullName := res.GetNewfullname()

// 	InfoGroup, err := db.UpdateData(oldName, newName, newFullName)
// 	// InfoGroup := "data received : "

// 	if err != nil {
// 		logs.Logger.Error("UpdateData - Update Data failed error: ", err)
// 		return nil, status.Errorf(codes.Unimplemented, "Update Data failed")
// 	}

// 	rsp := &pb.DataUpdateRespone{
// 		Notice: InfoGroup,
// 	}

// 	return rsp, nil
// }

// func (server *Server) ExportData(ctx context.Context, res *pb.ExportDataResquest) (*pb.ExportDataRespone, error) {

// 	logs.Logger.Info("ExportData - API call ExportData ")

// 	var dataInfo []model.DataInfo
// 	var err error

// 	switch server.config.Sever.TypeServer.Name {

// 	case "server_levedb":

// 		temp := levedb.GetData(server.config)
// 		dataInfo = append(dataInfo, temp)

// 	case "server_mysql":
// 		var data []model.DataPost
// 		err := mysqldb.GetData(&data)
// 		if err != nil {
// 			logs.Logger.Error("ExportData: server_mysql get Data failed error:", err)
// 			return nil, status.Errorf(codes.Unimplemented, "get Data failed")
// 		}
// 		for i := 0; i < len(data); i++ {
// 			dataInfo = append(dataInfo, model.DataInfo{Name: data[i].Name, FullName: data[i].FullName})
// 		}
// 	case "server_postgressql":
// 		dataInfo, err = db.GetData()
// 		if err != nil {
// 			logs.Logger.Error("ExportData: server_postgressql get Data failed error:", err)
// 			return nil, status.Errorf(codes.Unimplemented, "get Data failed")
// 		}
// 	case "server_mongodb":
// 		dataInfo, err = mongodb.GetAllInfo(server.config)
// 		if err != nil {
// 			logs.Logger.Error("ExportData: server_mongodb get Data failed error:", err)
// 			return nil, status.Errorf(codes.Unimplemented, "get Data failed")
// 		}
// 	default:
// 		logs.Logger.Error("ExportData: Don't have database")
// 		return nil, status.Errorf(codes.Unimplemented, "Don't have database")
// 	}

// 	var DataInfo []model.DataInfo
// 	var expenseData = [][]interface{}{}

// 	for i := 0; i < len(dataInfo); i++ {
// 		DataInfo = append(DataInfo, model.DataInfo{Name: dataInfo[i].Name, FullName: dataInfo[i].FullName})
// 		temp := []interface{}{dataInfo[i].Name, dataInfo[i].FullName}
// 		expenseData = append(expenseData, temp)
// 	}

// 	f := excelize.NewFile()
// 	index, _ := f.NewSheet("Sheet1")
// 	f.SetActiveSheet(index)

// 	err = f.SetSheetRow("Sheet1", "A1", &[]interface{}{"Name", "FullName"})
// 	if err != nil {

// 		logs.Logger.Error("ExportData: Error SetSheetRow: ", err)
// 	}
// 	err = f.SetColWidth("Sheet1", "A", "G", 30)

// 	if err != nil {

// 		logs.Logger.Error("ExportData: Error SetColWidth: ", err)
// 	}

// 	startRow := 2
// 	for i := startRow; i < (len(expenseData) + startRow); i++ {
// 		err = f.SetSheetRow("Sheet1", fmt.Sprintf("A%d", i), &expenseData[i-2])
// 		if err != nil {
// 			logs.Logger.Error("ExportData: Error SetSheetRow: ", err)
// 		}
// 	}

// 	// Save spreadsheet by the given path.
// 	if err := f.SaveAs("./export/DataExportFromDB.xlsx"); err != nil {
// 		fmt.Println(err)
// 	}

// 	if err != nil {
// 		logs.Logger.Error("ExportData: Error SaveAs: ", err)
// 	}

// 	dir, err := filepath.Abs(filepath.Dir("DataExportFromDB.xlsx"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	pathExport := dir + "/export/DataExportFromDB.xlsx"

// 	rsp := &pb.ExportDataRespone{
// 		PathExport: pathExport,
// 	}

// 	return rsp, nil
// }

// API Sign Up
func (server *Server) SignUp(ctx context.Context, res *pb.SignUpResquest) (*pb.SignUpRespone, error) {
	fmt.Println("ExportData - API call SignUp ")

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ultils.ErrorHandlerGRPC(codes.Unauthenticated, "Missing context metadata", http.StatusUnauthorized)
	}

	if len(md["username"]) != 1 || len(md["password"]) != 1 || len(md["lastname"]) != 1 || len(md["firstname"]) != 1 {
		return nil, ultils.ErrorHandlerGRPC(codes.Unauthenticated, "Invalid username, password, firstname or lastname", http.StatusUnauthorized)
	}

	resp, err := server.clientAuthen.SignUp(ctx, md["username"][0], md["password"][0], md["lastname"][0], md["firstname"][0])

	if err != nil {
		return nil, ultils.ErrorHandlerGRPC(codes.InvalidArgument, "Call function SignUp failed", http.StatusBadRequest)
	}
	return &pb.SignUpRespone{Message: resp, CodeHttp: http.StatusOK, Status: "success"}, nil
}

// API Login
func (server *Server) LogInAcc(ctx context.Context, res *pb.SignInResquest) (*pb.SignInRespone, error) {
	fmt.Println("ExportData - API call LogInAcc ")

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ultils.ErrorHandlerGRPC(codes.Unauthenticated, "Missing context metadata", http.StatusUnauthorized)
	}

	if len(md["username"]) != 1 || len(md["password"]) != 1 {
		return nil, ultils.ErrorHandlerGRPC(codes.Unauthenticated, "Invalid username or password", http.StatusUnauthorized)
	}

	resp, err := server.clientAuthen.SignInClient(ctx, md["username"][0], md["password"][0])

	if err != nil {
		return nil, ultils.ErrorHandlerGRPC(codes.Unauthenticated, "User or password is incorrect", http.StatusUnauthorized)
	}

	infoUser := &pb.UserAccInfo{

		Username: resp.Userinfo.Username,
		Fullname: resp.Userinfo.Firstname + " " + resp.Userinfo.Lastname,
	}

	info := &pb.SignInRespone{
		Status:      "success",
		CodeHttp:    http.StatusOK,
		UserAccInfo: infoUser,
		Token:       resp.Token,
	}
	return info, nil
}

// API Get List File Upload
func (server *Server) GetFileUploadInfo(ctx context.Context, res *pb.FileUploadInfoResquest) (*pb.FileUploadInfoRespone, error) {
	fmt.Println("ExportData - API call GetFileUploadInfo ")

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ultils.ErrorHandlerGRPC(codes.Unauthenticated, "Missing context metadata", http.StatusUnauthorized)
	}

	if len(md["token"]) != 1 {
		return nil, ultils.ErrorHandlerGRPC(codes.Unauthenticated, "Invalid token", http.StatusUnauthorized)
	}

	resp, err := server.clientAuthen.AuthenTokenClient(context.Background(), md["token"][0])
	if err != nil {

		return nil, ultils.ErrorHandlerGRPC(codes.Unauthenticated, "Authen token failed", http.StatusUnauthorized)
	}

	respDatabase, err := server.clientStogare.GetUploadFileInfoClient(context.Background(), resp.Iduser)

	if err != nil {
		errorName := strings.Split(err.Error(), "=")
		return nil, ultils.ErrorHandlerGRPC(codes.InvalidArgument, "Call function GetUploadFileInfoClient error:"+errorName[len(errorName)-1], http.StatusBadRequest)
	}

	var temp []*pb.FileUploadInfo

	for i := 0; i < len(respDatabase.Fileinfo); i++ {
		temp = append(temp, &pb.FileUploadInfo{
			FileId:       respDatabase.Fileinfo[i].Fileid,
			FileName:     respDatabase.Fileinfo[i].Filename,
			FileType:     respDatabase.Fileinfo[i].Typefile,
			SizeFile:     int64(respDatabase.Fileinfo[i].Size),
			Link:         respDatabase.Fileinfo[i].Link,
			TimeCreateAt: respDatabase.Fileinfo[i].Timecreateat,
		})
	}

	return &pb.FileUploadInfoRespone{FileInfoList: temp, Status: "success", CodeHttp: http.StatusOK}, nil
}

// API Get Short Infomation of a File Upload
func (server *Server) GetFileUploadShortInfo(ctx context.Context, res *pb.FileUploadShortInfoResquest) (*pb.FileUploadShortInfoRespone, error) {
	fmt.Println("ExportData - API call GetFileUploadShortInfo ")

	idFile := res.GetIdfile()

	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, ultils.ErrorHandlerGRPC(codes.Unauthenticated, "Missing context metadata", http.StatusUnauthorized)
	}

	if len(md["token"]) != 1 {
		return nil, ultils.ErrorHandlerGRPC(codes.Unauthenticated, "Invalid token", http.StatusUnauthorized)
	}

	_, err := server.clientAuthen.AuthenTokenClient(context.Background(), md["token"][0])

	if err != nil {
		return nil, ultils.ErrorHandlerGRPC(codes.Unauthenticated, "Authen token failed", http.StatusUnauthorized)
	}

	respDatabase, err := server.clientStogare.GetUploadFileShortInfoClient(ctx, idFile)

	if err != nil {
		errorName := strings.Split(err.Error(), "=")
		return nil, ultils.ErrorHandlerGRPC(codes.InvalidArgument, "Call function GetUploadFileInfoClient error:"+errorName[len(errorName)-1], http.StatusBadRequest)
	}

	linkSplit := strings.Split(respDatabase.Fileinfo.Link, "/")

	linkStogare := strings.Join(linkSplit[:len(linkSplit)-1], "/")

	temp := &pb.FileUploadInfo{
		FileId:       idFile,
		FileName:     respDatabase.Fileinfo.Filename,
		FileType:     respDatabase.Fileinfo.Typefile,
		SizeFile:     int64(respDatabase.Fileinfo.Size),
		Link:         linkStogare,
		TimeCreateAt: respDatabase.Fileinfo.Timecreateat,
	}

	return &pb.FileUploadShortInfoRespone{FileShortInfo: temp, Status: "success", CodeHttp: http.StatusOK}, nil
}

// API Upload File
func (server *Server) ImportDataWithHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ExportData - API call ImportDataWithHttp ")
	token := r.Header.Get("token")
	idProcsessService := r.Header.Get("id_process_service")
	idUploadService := r.Header.Get("id_upload_service")
	idFunctionProcess := r.Header.Get("id_function_process")

	var temp string

	maxFileSize := int64(3 * 1024 * 1024)

	if len(token) == 0 || len(idUploadService) == 0 {
		ultils.ErrorHandler(w, r, http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "Get token, idUploadService  failed",
		})
		return
	}

	resp, err := server.clientAuthen.AuthenTokenClient(context.Background(), token)
	if err != nil {
		ultils.ErrorHandler(w, r, http.StatusUnauthorized, map[string]interface{}{
			"code":    http.StatusUnauthorized,
			"status":  "error",
			"message": "Authen token failed",
		})
		return
	}

	file_ex, a, err := r.FormFile("file")

	if err != nil {
		ultils.ErrorHandler(w, r, http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "Failed to retrieve file from form data",
		})
		return
	}
	defer file_ex.Close()

	if a.Header.Get("Content-Type") != "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
		ultils.ErrorHandler(w, r, http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "Format file error (xlsx)",
		})
		return
	}

	if a.Size > maxFileSize {
		ultils.ErrorHandler(w, r, http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "Size file > 3 Mb",
		})
		return
	}

	var infoFileUpLoad *pb_storage.FileInfoRespone
	content, err := ioutil.ReadAll(file_ex)

	if err != nil {
		ultils.ErrorHandler(w, r, http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "ReadAll Failed",
		})
		return
	}

	switch idUploadService {
	case "1":
		infoFileUpLoad, err = server.clientStogare.UploadFile(context.Background(), a.Filename, a.Header.Get("Content-Type"), resp.Iduser, a.Size, content)
		if err != nil {
			errorName := strings.Split(err.Error(), "=")
			ultils.ErrorHandler(w, r, http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"status":  "error",
				"message": "Funtion UploadFile error:" + errorName[len(errorName)-1],
			})
			return
		}
	default:
		ultils.ErrorHandler(w, r, http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "Id Upload Service incorrect",
		})
		return
	}

	if idProcsessService == "1" && idFunctionProcess == "1" {
		_, err := server.clientProcessing.ProcessingDataClient(context.Background(), infoFileUpLoad.Id, a.Filename, infoFileUpLoad.Link, infoFileUpLoad.CheckSum)

		if err != nil {
			ultils.ErrorHandler(w, r, http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"status":  "error",
				"message": "Funtion ProcessingDataClient Failed",
			})
			return
		}
		temp = "Upload Done and Processing Done"
	} else {
		temp = "Upload Done and Not Processing"
	}

	type Respone struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Status  string `json:"status"`
	}

	noticeDb := Respone{
		Code:    http.StatusOK,
		Status:  "success",
		Message: temp,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(noticeDb)
}

// API Export by Template
func (server *Server) ExportDataHttp(w http.ResponseWriter, r *http.Request) {

	fmt.Println("ExportData - API call ExportDataHttp ")

	values := r.URL.Query()

	template := values.Get("template")

	token := r.Header.Get("token")

	if len(template) == 0 {
		ultils.ErrorHandler(w, r, http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "Get template or token failed",
		})
		return
	}

	if len(token) == 0 {
		ultils.ErrorHandler(w, r, http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "Get token or token failed",
		})
		return
	}

	_, err := server.clientAuthen.AuthenTokenClient(context.Background(), token)
	if err != nil {
		ultils.ErrorHandler(w, r, http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "Authen token failed",
		})
	}

	resp, err := server.clientProcessing.ExportFileTemplateExcelClient(context.Background(), template)

	if err != nil {
		errorName := strings.Split(err.Error(), "=")
		ultils.ErrorHandler(w, r, http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "Call Function ExportFileTemplateExcelClient error:" + errorName[len(errorName)-1],
		})
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=test.xlsx")
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Write([]byte(resp.PathExport))
}

// API Export by Funtion
func (server *Server) ExportFuntionWithHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API call ExportFuntionWithHttp ")
	values := r.URL.Query()

	account := values.Get("account")

	token := r.Header.Get("token")

	if len(account) == 0 || len(token) == 0 {
		ultils.ErrorHandler(w, r, http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "Get account or token  failed",
		})
		return
	}

	_, err := server.clientAuthen.AuthenTokenClient(context.Background(), token)

	if err != nil {
		ultils.ErrorHandler(w, r, http.StatusUnauthorized, map[string]interface{}{
			"code":    http.StatusUnauthorized,
			"status":  "error",
			"message": "Authen token failed",
		})
		return
	}

	resp, err := server.clientProcessing.ExportFuntionClient(context.Background(), account)

	if err != nil {
		errorName := strings.Split(err.Error(), "=")
		ultils.ErrorHandler(w, r, http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "Call funtion ExportFuntionClient error:" + errorName[len(errorName)-1],
		})
		return
	}

	type Respone struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Status  string `json:"status"`
	}

	w.Header().Set("Content-Disposition", "attachment; filename=InfoTransaction.xlsx")
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Write(resp.Content)

}

// API Download File by ID File
func (server *Server) DowloadLinkWithHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API call DowloadLinkWithHttp ")
	values := r.URL.Query()

	idFile := values.Get("idfile")

	token := r.Header.Get("token")

	if len(idFile) == 0 || len(token) == 0 {
		ultils.ErrorHandler(w, r, http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "Get idFile or token  failed",
		})
		return
	}

	_, err := server.clientAuthen.AuthenTokenClient(context.Background(), token)

	if err != nil {
		ultils.ErrorHandler(w, r, http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "Authen token failed",
		})
		return
	}

	resp, err := server.clientStogare.DownloadFileClient(context.Background(), idFile)

	if err != nil {
		errorName := strings.Split(err.Error(), "=")
		ultils.ErrorHandler(w, r, http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "Call Function DownloadFileClient error:" + errorName[len(errorName)-1],
		})
		return
	}

	name := strings.Split(resp.Name, "/")

	w.Header().Set("Content-Disposition", "attachment; filename="+name[len(name)-1])
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Write(resp.Content)
}

// API Update Status Proccess File
func (server *Server) UpdateStatusWithHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API call UpdateStatusWithHttp ")
	status := r.Header.Get("status")
	idFile := r.Header.Get("idfile")

	if len(status) == 0 || len(idFile) == 0 {
		ultils.ErrorHandler(w, r, http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": " Status or idfile invalid",
		})
		return
	}
	resp, err := server.clientStogare.UpdateStatusFileClient(context.Background(), idFile, status)

	if err != nil {
		ultils.ErrorHandler(w, r, http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "Call funtion UpdateStatusFileClient failed",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Noti)
}

// API???
func (server *Server) GetListServiceUpload(ctx context.Context, res *pb.GetListServiceUploadResquest) (*pb.GetListServiceUploadRespone, error) {

	md, _ := metadata.FromIncomingContext(ctx)
	// log.Info("nhatnt md: ", md)

	_, err := server.clientAuthen.AuthenTokenClient(context.Background(), md["token"][0])

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Authen token failed")
	}

	resp, err := mongodb.GetListServiceUpload(server.config)
	fmt.Println(resp)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "GetListServiceUpload failed")
	}
	var temp []*pb.ServiceInfo

	for i := 0; i < len(resp); i++ {
		temp = append(temp, &pb.ServiceInfo{
			Id:   resp[i].Id,
			Name: resp[i].Name,
		})
	}

	return &pb.GetListServiceUploadRespone{ServiceInfo: temp}, nil
}

func (server *Server) GetListServiceProcess(ctx context.Context, res *pb.GetListServiceProcessResquest) (*pb.GetListServiceProcessRespone, error) {

	IdServiceUpload := res.GetIdServiceUpload()

	md, _ := metadata.FromIncomingContext(ctx)
	// log.Info("nhatnt md: ", md)

	_, err := server.clientAuthen.AuthenTokenClient(context.Background(), md["token"][0])

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Authen token failed")
	}

	resp, err := mongodb.GetListServiceProcess(server.config, IdServiceUpload)
	fmt.Println(resp)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "GetListServiceUpload failed")
	}
	var temp []*pb.ServiceInfo

	for i := 0; i < len(resp); i++ {
		temp = append(temp, &pb.ServiceInfo{
			Id:   resp[i].Id,
			Name: resp[i].Name,
		})
	}

	return &pb.GetListServiceProcessRespone{ServiceInfoProcess: temp}, nil
}
func (server *Server) GetListFunctionProcess(ctx context.Context, res *pb.GetListFunctionProcessResquest) (*pb.GetListFunctionProcessRespone, error) {

	IdServiceProcess := res.GetIdServiceProcess()

	md, _ := metadata.FromIncomingContext(ctx)
	// log.Info("nhatnt md: ", md)

	_, err := server.clientAuthen.AuthenTokenClient(context.Background(), md["token"][0])

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Authen token failed")
	}

	resp, err := mongodb.GetListFunctionProcess(server.config, IdServiceProcess)
	fmt.Println(resp)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "GetListServiceUpload failed")
	}
	var temp []*pb.ServiceInfo

	for i := 0; i < len(resp); i++ {
		temp = append(temp, &pb.ServiceInfo{
			Id:   resp[i].Id,
			Name: resp[i].Name,
		})
	}

	return &pb.GetListFunctionProcessRespone{FunctionInfoProcess: temp}, nil
}
