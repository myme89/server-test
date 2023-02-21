package api

import (
	"server-test/model"
	"server-test/pb"
)

func ConvertData(data []model.DataInfo) []*pb.DataInfo {
	var temp []*pb.DataInfo
	for i := 0; i < len(data); i++ {
		temp = append(temp, &pb.DataInfo{
			Name:     data[i].Name,
			Fullname: data[i].FullName,
		})
	}
	// log.Info("trongnhat temp = ", temp)
	return temp
}

// func ConvertDataCache(data []model.DataInfo) []*pb.DataInfo {
// 	var temp []*pb.DataInfo
// 	for i := 0; i < len(data); i++ {
// 		temp = append(temp, &pb.DataInfo{
// 			Name:     data[i].Name,
// 			Fullname: data[i].FullName,
// 		})
// 	}
// 	// log.Info("trongnhat temp = ", temp)
// 	return temp
// }

// func ConvertData1(data string) []*pb.DataInfo {
// 	var temp []*pb.DataInfo
// 	for i := 0; i < len(data); i++ {
// 		temp = append(temp, &pb.DataInfo{
// 			Name:     data[i].Name,
// 			Fullname: data[i].FullName,
// 		})
// 	}
// 	// log.Info("trongnhat temp = ", temp)
// 	return temp
// }
