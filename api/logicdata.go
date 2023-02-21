package api

import (
	"context"
	"encoding/json"
	"fmt"
	"server-test/cache"
	"server-test/db"
	"server-test/model"
	"server-test/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetData(ctx context.Context, res *pb.DataInfoResquest) (*pb.DataInfoRespone, error) {

	valueCache := cache.GetFromRedis(ctx, "data_test")

	var rsp *pb.DataInfoRespone
	if len(valueCache) == 0 {

		dataInfo, err := db.GetData()

		if err != nil {
			return nil, status.Errorf(codes.Unimplemented, "get Data failed")
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
