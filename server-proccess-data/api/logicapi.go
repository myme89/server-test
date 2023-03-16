package api

import (
	"context"
	"fmt"
	"server-test/server-proccess-data/pb_processing"
)

// func (serverStorage *ServerStorage) TestData(ctx context.Context, res *pb_storage.DataInfoTestResquest) (*pb_storage.DataInfoTestRespone, error) {

// 	temp := res.GetDataTest()

// 	fmt.Println("TestData server Storage")

// 	return &pb_storage.DataInfoTestRespone{DataResp: temp}, nil
// }

func (serverProcessing *ServerProcessing) TestData2(ctx context.Context, res *pb_processing.DataInfoTestResquest1) (*pb_processing.DataInfoTestRespone1, error) {

	temp := res.GetDataTest() + " TestData2" + " TestData2"

	fmt.Println("TestData server Storage")

	return &pb_processing.DataInfoTestRespone1{DataResp: temp}, nil
}
