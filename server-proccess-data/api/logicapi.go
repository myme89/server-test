package api

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"server-test/server-proccess-data/database/mongodb"
	"server-test/server-proccess-data/model"
	"server-test/server-proccess-data/pb_processing"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/xuri/excelize/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"server-test/server-proccess-data/model"
// 	"server-test/server-proccess-data/pb_processing"
// 	"time"

// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"

// 	"github.com/xuri/excelize/v2"
// )

// func (serverProcessing *ServerProcessing) TestData2(ctx context.Context, res *pb_processing.DataInfoTestResquest1) (*pb_processing.DataInfoTestRespone1, error) {

// 	fmt.Println("TestData server Storage start")
// 	temp := res.GetDataTest() + " TestData2" + " TestData2"

// 	time.Sleep(5 * time.Second)

// 	fmt.Println("TestData server Storage end")

// 	return &pb_processing.DataInfoTestRespone1{DataResp: temp}, nil
// }

func (serverProcessing *ServerProcessing) ProcessingFileExcel(ctx context.Context, res *pb_processing.ProcessingFileResquest) (*pb_processing.ProcessingFileRespone, error) {

	infoFileProcess := res.GetFileinfoprocess()

	xlsx, err := excelize.OpenReader(bytes.NewReader(infoFileProcess.Content))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "OpenReader file excel failed")
	}
	var infoTrans []model.TemplateInfoTransaction
	// var info_file []model.InfoFile
	for _, name := range xlsx.GetSheetMap() {
		// fmt.Println(index, name)
		// nameSheet = append(nameSheet, name)
		dataRows, err := xlsx.GetRows(name)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "GetRows file excel failed")
		}

		// var data []map[string]string

		for _, row := range dataRows {

			fmt.Println("trongnhat test", row[0])

			// item := make(map[string]string)
			// for i, colCell := range row {
			// 	temp, _ := excelize.ColumnNumberToName(i + 1)

			// 	temp1, _ := xlsx.GetCellValue(name, temp+"1")
			// 	// infoTran := model.TemplateInfoTransaction{
			// 	// 	IdTran: colCell,
			// 	// }
			// 	item[temp1] = colCell
			// 	fmt.Println("trongnhat test", item[temp1])
			// }
			// // fmt.Println("trongnhat test", item)
			// data = append(data, item)

			infoTrans = append(infoTrans, model.TemplateInfoTransaction{
				IdTran:    row[0],
				AccRec:    row[1],
				AccSend:   row[2],
				AccFee:    row[3],
				Deposits:  row[4],
				MoneyRec:  row[5],
				Fee:       row[6],
				TimeTrans: row[7],
			})

		}
		// jsonData, err := json.Marshal(data)

		// if err != nil {
		// 	return nil, status.Errorf(codes.InvalidArgument, "Marshal data failed")

		// }

		// info_file = append(info_file, model.InfoFile{Name: name, Content: jsonData})
	}

	date := "23"
	var result []model.TemplateInfoTransaction

	infoTran := infoTrans[1:]
	// fmt.Println("data test=", t)

	for i := 0; i < len(infoTran); i++ {
		t := strings.Split(infoTran[i].TimeTrans, " ")
		k := strings.Split(t[0], "-")
		fmt.Println("data test=", k)
		if infoTran[i].AccRec == "1234" && k[2] == date {
			result = append(result, model.TemplateInfoTransaction{
				IdTran:    infoTran[i].IdTran,
				AccRec:    infoTran[i].AccRec,
				AccSend:   infoTran[i].AccSend,
				AccFee:    infoTran[i].AccFee,
				Deposits:  infoTran[i].Deposits,
				MoneyRec:  infoTran[i].MoneyRec,
				Fee:       infoTran[i].Fee,
				TimeTrans: infoTran[i].TimeTrans,
			})
		}
	}

	// fmt.Println("data test=", count)

	err = mongodb.AddManyInfoTrans(serverProcessing.config, result)

	// nameFileExcel := infoFileProcess.Filename + " " + infoFileProcess.Idfile
	// resp, err := serverProcessing.clientDatabase.UploadDataFileExcelClient(ctx, info_file, nameFileExcel)

	// for i := 0; i < len(info_file); i++ {
	// 	// Create an empty map to unmarshal JSON into
	// 	var person []map[string]interface{}

	// 	// Unmarshal the JSON data into the map
	// 	err := json.Unmarshal([]byte(info_file[i].Content), &person)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return nil, status.Errorf(codes.Unimplemented, "Unmarshal failed")
	// 	}

	// 	infos := make([]interface{}, len(person))
	// 	for i, s := range person {
	// 		infos[i] = s
	// 	}

	// 	infos = infos[1:]
	// 	fmt.Println("trongnhat test1= ", infos)
	// 	// var datasAdd []model.TemplateInfoTransaction
	// 	// for i := 0; i < len(infos); i++ {
	// 	// 	datasAdd = append(datasAdd, model.TemplateInfoTransaction{AccRec: , Name: data[i][1], FullName: data[i][2]})

	// 	// }

	// 	fmt.Println("nhatnt test", serverProcessing.config)
	// 	err = mongodb.AddManyInfoNotModel(serverProcessing.config, infos, nameFileExcel)
	// 	if err != nil {
	// 		log.Error(" Post Data to Mongo Database failed:", err)
	// 		return nil, status.Errorf(codes.Unimplemented, "AddManyInfoNotModel failed")
	// 	}

	// }

	url := "http://localhost:3000/v1/status"

	fmt.Println("trongnhat 1")

	// Replace with the data you want to send in the request body
	// data := map[string]string{"status": "Done", "idfile": infoFileProcess.Idfile}
	// payload, err := json.Marshal(data)
	// if err != nil {
	// 	// Handle error
	// 	return nil, status.Errorf(codes.InvalidArgument, "Error while encoding data to JSON")
	// 	// fmt.Println("Error while encoding data to JSON:", err)
	// 	// return
	// }

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		// Handle error
		return nil, status.Errorf(codes.InvalidArgument, "Error while creating request:")

		// fmt.Println("Error while creating request:", err)
		// return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("status", "Done")
	req.Header.Set("idfile", infoFileProcess.Idfile)

	client := &http.Client{}
	resp, err := client.Do(req)

	// fmt.Println("1231212312= ", resp)
	if err != nil {
		// Handle error
		return nil, status.Errorf(codes.InvalidArgument, "Error making request")

	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Handle error
		return nil, status.Errorf(codes.InvalidArgument, "Error reading response body")

		// fmt.Println("Error reading response body:", err)
		// return
	}

	// Use the content of the response as needed
	fmt.Println(string(body))
	// if err != nil {
	// 	return nil, status.Errorf(codes.InvalidArgument, "UploadDataFileExcelClient failed")
	// }

	// noti, err := serverProcessing.clientDatabase.UpdateStatusProcessingClient(ctx, "Done", infoFileProcess.Idfile)

	// if err != nil {
	// 	return nil, status.Errorf(codes.InvalidArgument, "UpdateStatusProcessingClient failed")

	// }
	// fmt.Println("noti ", resp)
	return &pb_processing.ProcessingFileRespone{Noti: "Processing Done"}, nil
}

func (serverProcessing *ServerProcessing) ExportTemplateFileUpload(ctx context.Context, res *pb_processing.ExportFileResquest) (*pb_processing.ExportFileRespone, error) {

	nameTemplate := res.GetTemplateExport()

	var expenseData = [][]interface{}{}

	f := excelize.NewFile()
	index, _ := f.NewSheet("Sheet1")
	f.SetActiveSheet(index)

	if nameTemplate == "TemplateInfoPerson" {

		dataInfo, err := mongodb.GetAllInfo(serverProcessing.config, nameTemplate)
		// resp, err := serverStorage.clientDatabase.ExportFileTemplateExcelClient(ctx, nameTemplate)

		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "ExportTemplateFileUpload  failed")
		}

		for i := 0; i < len(dataInfo); i++ {
			temp := []interface{}{dataInfo[i].LastName, dataInfo[i].FistName, dataInfo[i].FullName, dataInfo[i].PhoneNumber, dataInfo[i].Address}
			expenseData = append(expenseData, temp)
		}

		err = f.SetSheetRow("Sheet1", "A1", &[]interface{}{"Last Name", "First Name", "Full Name", "Phone Number", "Address"})
		if err != nil {

			log.Error("ExportData: Error SetSheetRow: ", err)
		}
	} else {

		dataInfo, err := mongodb.GetAllInfoTrans(serverProcessing.config, nameTemplate)

		// resp, err := serverStorage.clientDatabase.ExportFileTemplateExcelClient(ctx, nameTemplate)

		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "ExportTemplateFileUpload  failed")
		}

		for i := 0; i < len(dataInfo); i++ {
			temp := []interface{}{dataInfo[i].IdTran, dataInfo[i].AccRec, dataInfo[i].AccSend, dataInfo[i].AccFee, dataInfo[i].Deposits, dataInfo[i].MoneyRec, dataInfo[i].Fee, dataInfo[i].TimeTrans}
			expenseData = append(expenseData, temp)
		}

		err = f.SetSheetRow("Sheet1", "A1", &[]interface{}{"ID Trans", "Account Receive", "Account Send", "Account Fee", "Deposits", "Money Received", "Fee", "Time"})
		if err != nil {

			log.Error("ExportData: Error SetSheetRow: ", err)
		}
	}

	err := f.SetColWidth("Sheet1", "A", "G", 30)

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
	// if err := f.SaveAs("./storage-export/Test.xlsx"); err != nil {
	// 	fmt.Println(err)
	// }

	// if err != nil {
	// 	log.Error("ExportData: Error SaveAs: ", err)
	// }

	fileBytes, err := f.WriteToBuffer()
	if err != nil {
		log.Fatal(err)
	}

	// dir, err := filepath.Abs(filepath.Dir("TemplateInfoPerson.xlsx"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// pathExport := dir + "/stogare-export/DataExportFromDB.xlsx"
	// pathExport := "localhost:3000/v1/downloadlink?dir=" + "storage-export/Test.xlsx"

	temp := fileBytes.Bytes()
	return &pb_processing.ExportFileRespone{PathExport: temp}, nil
}

func (serverProcessing *ServerProcessing) DownloafFileProcess(ctx context.Context, res *pb_processing.DownloadFileProcessResquest) (*pb_processing.DownloadFileProcessRespone, error) {

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

	return &pb_processing.DownloadFileProcessRespone{NameFile: name, ContentFile: content}, nil
}
