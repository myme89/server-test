package api

import (
	"context"
	"encoding/json"
	"fmt"
	"server-test/cache"
	"server-test/database/db"
	"server-test/database/levedb"
	"server-test/database/mongodb"
	"server-test/database/mysqldb"
	"server-test/model"
	"server-test/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetData(ctx context.Context, res *pb.DataInfoResquest) (*pb.DataInfoRespone, error) {

	valueCache := cache.GetFromRedis(ctx, "data_test")

	var rsp *pb.DataInfoRespone

	if len(valueCache) == 0 {
		var dataInfo []model.DataInfo
		var err error

		switch server.config.Sever.TypeServer.Name {

		case "server_levedb":

			temp := levedb.GetData(server.config)
			dataInfo = append(dataInfo, temp)

		case "server_mysql":
			var data []model.DataPost
			err := mysqldb.GetData(&data)
			if err != nil {
				return nil, status.Errorf(codes.Unimplemented, "get Data failed")
			}
			for i := 0; i < len(data); i++ {
				dataInfo = append(dataInfo, model.DataInfo{Name: data[i].Name, FullName: data[i].FullName})
			}
		case "server_postgressql":
			dataInfo, err = db.GetData()
			if err != nil {
				return nil, status.Errorf(codes.Unimplemented, "get Data failed")
			}
		case "server_mongodb":
			dataInfo, err = mongodb.GetAllInfo()
		default:
			return nil, status.Errorf(codes.Unimplemented, "Don't have database")
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

	type Info struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
	}
	// InfoGroup, err := db.PostData(Id, Name, FullName)
	// InfoGroup := "data received : "

	// json1, err := json.Marshal(Info{Id: Id, Name: Name, FullName: FullName})

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// cache.SetToRedis(ctx, "data_test", json1)

	// valueCache := cache.GetFromRedis(ctx, "data_test")
	// fmt.Println("value Cache= ", valueCache)

	// rsp := &pb.DataPostRespone{
	// 	Notice: InfoGroup,
	// }

	// values := cache.GetAllKeys(ctx, "data_test")

	// log.Info("All values : %v \n", values)

	// time.AfterFunc(60*time.Second, func() {
	// 	ctx1 := context.TODO()
	// 	valueCache1 := cache.GetFromRedis(ctx1, "data_test")

	// 	var temp Info
	// 	err := json.Unmarshal([]byte(valueCache1), &temp)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println("value Cache= ", temp)

	// 	info, err := db.PostData(temp.Id, temp.Name, temp.FullName)
	// 	if err != nil {
	// 		log.Error("Post Data to Database failed")
	// 	}

	// 	log.Info(info)
	// })
	var noticeDb string
	switch server.config.Sever.TypeServer.Name {

	case "server_levedb":
		err := levedb.PutData(server.config, Id, Name, FullName)
		if err != nil {
			return nil, status.Errorf(codes.Unimplemented, "Post Data to Leve Database failed")
		}
		noticeDb = "Post Done"

	case "server_mysql":
		err := mysqldb.PostData(Id, Name, FullName)
		if err != nil {
			return nil, status.Errorf(codes.Unimplemented, "Post Data to MySql Database failed")
		}
		noticeDb = "Post Done"

	case "server_postgressql":

		info, err := db.PostData(Id, Name, FullName)
		if err != nil {
			return nil, status.Errorf(codes.Unimplemented, "Post Data to Postges Database failed")
		}
		noticeDb = info

	case "server_mongodb":
		err := mongodb.AddInfo(Id, Name, FullName)
		if err != nil {
			return nil, status.Errorf(codes.Unimplemented, "Post Data to Mongo Database failed")
		}
		noticeDb = "Post Done"

	default:
		return nil, status.Errorf(codes.Unimplemented, "Don't have database")
	}

	rsp := &pb.DataPostRespone{
		Notice: noticeDb,
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
