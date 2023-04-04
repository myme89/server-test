package api

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"server-test/server-proccess-data/database/mongodb"
	"server-test/server-proccess-data/model"
	"server-test/server-proccess-data/pb_processing"

	log "github.com/sirupsen/logrus"

	"github.com/xuri/excelize/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (serverProcessing *ServerProcessing) ProcessingFileExcel(ctx context.Context, res *pb_processing.ProcessingFileResquest) (*pb_processing.ProcessingFileRespone, error) {

	infoFileProcess := res.GetFileinfoprocess()

	fmt.Println("trongnhat 123")
	go func(idFile, fileName, linkFile, checkSum string) {
		var status string
		file, err := os.Open(linkFile)

		if err != nil {
			fmt.Println("err= ", err)
			status = "Processing Error"
			CallAPIUpdateStatus(status, infoFileProcess.Idfile)
			return
		}

		content, err := ioutil.ReadAll(file)

		if HashContentFile(content) != checkSum {
			status = "Processing Error"
			CallAPIUpdateStatus(status, infoFileProcess.Idfile)
			return
		}

		xlsx, err := excelize.OpenReader(bytes.NewReader(content))
		if err != nil {
			status = "Processing Error"
			CallAPIUpdateStatus(status, infoFileProcess.Idfile)
			return
		}
		var infoTrans []model.TemplateInfoTransaction
		// var info_file []model.InfoFile
		for _, name := range xlsx.GetSheetMap() {
			// fmt.Println(index, name)
			// nameSheet = append(nameSheet, name)
			dataRows, err := xlsx.GetRows(name)
			if err != nil {
				status = "Processing Error"
				CallAPIUpdateStatus(status, infoFileProcess.Idfile)
				return
			}

			// var data []map[string]string

			for _, row := range dataRows {

				// fmt.Println("trongnhat test", row[0])

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

		// date := "23"
		// var result []model.TemplateInfoTransaction

		infoTran := infoTrans[1:]
		// fmt.Println("data test=", t)

		// for i := 0; i < len(infoTran); i++ {
		// 	t := strings.Split(infoTran[i].TimeTrans, " ")
		// 	k := strings.Split(t[0], "-")
		// 	fmt.Println("data test=", k)
		// 	if infoTran[i].AccRec == "1234" && k[2] == date {
		// 		result = append(result, model.TemplateInfoTransaction{
		// 			IdTran:    infoTran[i].IdTran,
		// 			AccRec:    infoTran[i].AccRec,
		// 			AccSend:   infoTran[i].AccSend,
		// 			AccFee:    infoTran[i].AccFee,
		// 			Deposits:  infoTran[i].Deposits,
		// 			MoneyRec:  infoTran[i].MoneyRec,
		// 			Fee:       infoTran[i].Fee,
		// 			TimeTrans: infoTran[i].TimeTrans,
		// 		})
		// 	}
		// }

		// fmt.Println("data test=", count)

		err = mongodb.AddManyInfoTrans(serverProcessing.config, infoTran)

		if err != nil {
			status = "Processing Error"
			CallAPIUpdateStatus(status, infoFileProcess.Idfile)
			return
		}

		status = "Processing Successfully"
		CallAPIUpdateStatus(status, infoFileProcess.Idfile)

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

	}(infoFileProcess.Idfile, infoFileProcess.Filename, infoFileProcess.LinkFile, infoFileProcess.CheckSum)

	return &pb_processing.ProcessingFileRespone{Noti: "Processing Done"}, nil
}

func (serverProcessing *ServerProcessing) ExportTemplateFileUpload(ctx context.Context, res *pb_processing.ExportFileResquest) (*pb_processing.ExportFileRespone, error) {

	nameTemplate := res.GetTemplateExport()

	f := excelize.NewFile()
	index, _ := f.NewSheet("Sheet1")
	f.SetActiveSheet(index)

	err := f.SetColWidth("Sheet1", "A", "Z", 30)

	if err != nil {

		log.Error("ExportData: Error SetColWidth: ", err)
		return nil, status.Errorf(codes.InvalidArgument, "SetColWidth  failed")

	}

	if nameTemplate == "TemplateInfoPerson" {
		err = f.SetSheetRow("Sheet1", "A1", &[]interface{}{"Last Name", "First Name", "Full Name", "Phone Number", "Address"})
		if err != nil {
			log.Error("ExportData: Error SetSheetRow: ", err)
			return nil, status.Errorf(codes.InvalidArgument, "SetSheetRow  failed")
		}
	} else {

		err := f.SetSheetRow("Sheet1", "A1", &[]interface{}{"ID Trans", "Account Receive", "Account Send", "Account Fee", "Deposits", "Money Received", "Fee", "Time"})
		if err != nil {

			log.Error("ExportData: Error SetSheetRow: ", err)
			return nil, status.Errorf(codes.InvalidArgument, "SetSheetRow  failed")
		}
	}

	fileBytes, err := f.WriteToBuffer()
	if err != nil {
		log.Fatal(err)
		return nil, status.Errorf(codes.InvalidArgument, "WriteToBuffer  failed")
	}

	contentFile := fileBytes.Bytes()
	return &pb_processing.ExportFileRespone{PathExport: contentFile}, nil
}

func (serverProcessing *ServerProcessing) GetTransactionByAccount(ctx context.Context, res *pb_processing.GetTransactionByAccountResquest) (*pb_processing.GetTransactionByAccountRespone, error) {

	account := res.GetAccount()

	infoTransaction, err := mongodb.GetTransactionByAccountRec(serverProcessing.config, account)

	if err != nil {
		// Handle error
		return nil, status.Errorf(codes.InvalidArgument, "GetTransactionByAccountRec failed")
	}

	var expenseData = [][]interface{}{}

	f := excelize.NewFile()
	index, _ := f.NewSheet("Sheet1")
	f.SetActiveSheet(index)

	for i := 0; i < len(infoTransaction); i++ {
		temp := []interface{}{infoTransaction[i].IdTran, infoTransaction[i].AccRec, infoTransaction[i].AccSend, infoTransaction[i].AccFee, infoTransaction[i].Deposits, infoTransaction[i].MoneyRec, infoTransaction[i].Fee, infoTransaction[i].TimeTrans}
		expenseData = append(expenseData, temp)
	}

	err = f.SetSheetRow("Sheet1", "A1", &[]interface{}{"ID Trans", "Account Receive", "Account Send", "Account Fee", "Deposits", "Money Received", "Fee", "Time"})
	if err != nil {

		log.Error("ExportData: Error SetSheetRow: ", err)
		return nil, status.Errorf(codes.InvalidArgument, "SetSheetRow failed")

	}

	err = f.SetColWidth("Sheet1", "A", "Z", 30)

	if err != nil {

		log.Error("ExportData: Error SetColWidth: ", err)
		return nil, status.Errorf(codes.InvalidArgument, "SetColWidth failed")

	}

	startRow := 2
	for i := startRow; i < (len(expenseData) + startRow); i++ {
		err = f.SetSheetRow("Sheet1", fmt.Sprintf("A%d", i), &expenseData[i-2])
		if err != nil {
			log.Error("ExportData: Error SetSheetRow: ", err)
			return nil, status.Errorf(codes.InvalidArgument, "ExportData: Error SetSheetRow failed")

		}
	}

	fileBytes, err := f.WriteToBuffer()
	if err != nil {
		log.Fatal(err)
		return nil, status.Errorf(codes.InvalidArgument, "WriteToBuffer failed")
	}

	temp := fileBytes.Bytes()
	return &pb_processing.GetTransactionByAccountRespone{Content: temp}, nil
}

func CallAPIUpdateStatus(status, idFile string) {
	url := "http://localhost:3000/v1/status"

	fmt.Println("trongnhat 1")

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error while creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("status", status)
	req.Header.Set("idfile", idFile)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error while creating request:", err)
		return
	}
	defer resp.Body.Close()
}

func HashContentFile(content []byte) string {

	key := []byte("trongnhat99tn")

	// compute HMAC-SHA256 hash of the message with key
	hash := hmac.New(sha256.New, key)
	hash.Write(content)
	mac := hash.Sum(nil)

	macString := hex.EncodeToString(mac)
	return macString
}
