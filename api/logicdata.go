package api

import (
	"context"
	"encoding/json"
	"fmt"
	"server-test/cache"
	"server-test/db"
	"server-test/model"
	"server-test/pb"

	log "github.com/sirupsen/logrus"
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

func (server *Server) PostData(ctx context.Context, res *pb.DataPostResqest) (*pb.DataPostRespone, error) {

	Id := int(res.GetId())
	Name := res.GetName()
	FullName := res.GetFullname()

	log.Info("nhatnt", Id)
	log.Info("nhatnt", Name)
	log.Info("nhatnt", FullName)

	InfoGroup, err := db.PostData(Id, Name, FullName)
	// InfoGroup := "data received : "

	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, " Post Data failed")
	}

	rsp := &pb.DataPostRespone{
		Notice: InfoGroup,
	}
	return rsp, nil
}

func (server *Server) UpdateData(ctx context.Context, res *pb.DataUpdateResqest) (*pb.DataUpdateRespone, error) {

	oldName := res.GetOldname()
	newName := res.GetNewname()
	newFullName := res.GetNewfullname()

	InfoGroup, err := db.UpdateData(oldName, newName, newFullName)
	// InfoGroup := "data received : "

	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "Update Data failed")
	}

	rsp := &pb.DataUpdateRespone{
		Notice: InfoGroup,
	}

	return rsp, nil
}
