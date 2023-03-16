package api

import (
	"context"
	"fmt"
	"server-test/server-storage/pb_storage"
)

func (serverStorage *ServerStorage) TestData(ctx context.Context, res *pb_storage.DataInfoTestResquest) (*pb_storage.DataInfoTestRespone, error) {

	temp := res.GetDataTest() + "trong nhat 1 processing"

	// temp1, err := serverStorage.client.TestData(ctx, temp)
	// fmt.Println("TestData server Storage  3= ", ctx)
	if res.GetDataTest() == "1" {
		fmt.Println("TestData server Storage  1")
		go func(temp1 string) {
			temp2, err := serverStorage.client.TestData(context.Background(), temp1)

			fmt.Println("TestData server Storage  1", temp2)
			fmt.Println("TestData server Storage  1", err)

		}(temp)
	}
	fmt.Println("TestData server Storage")
	// fmt.Println("TestData server Storage", err)
	return &pb_storage.DataInfoTestRespone{DataResp: temp}, nil
}
