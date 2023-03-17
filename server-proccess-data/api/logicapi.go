package api

import (
	"context"
	"fmt"
	"server-test/server-proccess-data/pb_processing"
	"time"
)

func (serverProcessing *ServerProcessing) TestData2(ctx context.Context, res *pb_processing.DataInfoTestResquest1) (*pb_processing.DataInfoTestRespone1, error) {

	fmt.Println("TestData server Storage start")
	temp := res.GetDataTest() + " TestData2" + " TestData2"

	time.Sleep(5 * time.Second)

	fmt.Println("TestData server Storage end")

	return &pb_processing.DataInfoTestRespone1{DataResp: temp}, nil
}
