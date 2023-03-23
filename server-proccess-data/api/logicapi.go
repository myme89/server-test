package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"server-test/server-proccess-data/model"
	"server-test/server-proccess-data/pb_processing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/xuri/excelize/v2"
)

func (serverProcessing *ServerProcessing) TestData2(ctx context.Context, res *pb_processing.DataInfoTestResquest1) (*pb_processing.DataInfoTestRespone1, error) {

	fmt.Println("TestData server Storage start")
	temp := res.GetDataTest() + " TestData2" + " TestData2"

	time.Sleep(5 * time.Second)

	fmt.Println("TestData server Storage end")

	return &pb_processing.DataInfoTestRespone1{DataResp: temp}, nil
}

func (serverProcessing *ServerProcessing) ProcessingFileExcel(ctx context.Context, res *pb_processing.ProcessingFileResquest) (*pb_processing.ProcessingFileRespone, error) {

	infoFileProcess := res.GetFileinfoprocess()

	xlsx, err := excelize.OpenReader(bytes.NewReader(infoFileProcess.Content))
	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "OpenReader file excel failed")
	}

	var info_file []model.InfoFile
	for _, name := range xlsx.GetSheetMap() {
		// fmt.Println(index, name)
		// nameSheet = append(nameSheet, name)
		dataRows, err := xlsx.GetRows(name)
		if err != nil {
			return nil, status.Errorf(codes.Unimplemented, "GetRows file excel failed")
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
			return nil, status.Errorf(codes.Unimplemented, "Marshal data failed")

		}

		info_file = append(info_file, model.InfoFile{Name: name, Content: jsonData})
	}

	nameFileExcel := infoFileProcess.Filename + " " + infoFileProcess.Idfile
	resp, err := serverProcessing.clientDatabase.UploadDataFileExcelClient(ctx, info_file, nameFileExcel)

	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "UploadDataFileExcelClient failed")
	}

	noti, err := serverProcessing.clientDatabase.UpdateStatusProcessingClient(ctx, "Done", infoFileProcess.Idfile)
	fmt.Println("noti ", resp)
	return &pb_processing.ProcessingFileRespone{Noti: noti}, nil
}
