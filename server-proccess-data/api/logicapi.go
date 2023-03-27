package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"server-test/server-proccess-data/database/mongodb"
	"server-test/server-proccess-data/model"
	"server-test/server-proccess-data/pb_processing"

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

	var info_file []model.InfoFile
	for _, name := range xlsx.GetSheetMap() {
		// fmt.Println(index, name)
		// nameSheet = append(nameSheet, name)
		dataRows, err := xlsx.GetRows(name)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "GetRows file excel failed")
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
			return nil, status.Errorf(codes.InvalidArgument, "Marshal data failed")

		}

		info_file = append(info_file, model.InfoFile{Name: name, Content: jsonData})
	}

	nameFileExcel := infoFileProcess.Filename + " " + infoFileProcess.Idfile
	// resp, err := serverProcessing.clientDatabase.UploadDataFileExcelClient(ctx, info_file, nameFileExcel)

	for i := 0; i < len(info_file); i++ {
		// Create an empty map to unmarshal JSON into
		var person []map[string]interface{}

		// Unmarshal the JSON data into the map
		err := json.Unmarshal([]byte(info_file[i].Content), &person)
		if err != nil {
			fmt.Println(err)
			return nil, status.Errorf(codes.Unimplemented, "Unmarshal failed")
		}

		infos := make([]interface{}, len(person))
		for i, s := range person {
			infos[i] = s
		}

		infos = infos[1:]

		fmt.Println("nhatnt test", serverProcessing.config)
		err = mongodb.AddManyInfoNotModel(serverProcessing.config, infos, nameFileExcel)
		if err != nil {
			log.Error(" Post Data to Mongo Database failed:", err)
			return nil, status.Errorf(codes.Unimplemented, "AddManyInfoNotModel failed")
		}

	}

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
