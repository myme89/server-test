package api

import (
	"server-test/server-database/model"
	"server-test/server-database/pb_database"
)

func ConvertData(data []model.FileInfoUpload) []*pb_database.FileUploadInfo {
	var temp []*pb_database.FileUploadInfo
	for i := 0; i < len(data); i++ {
		temp = append(temp, &pb_database.FileUploadInfo{

			Filename: data[i].FileName,
			Typefile: data[i].TypeFile,
			Size:     float32(data[i].Size),
			Link:     data[i].Link,
			Createat: data[i].CreateAt,
		})
	}
	// log.Info("trongnhat temp = ", temp)
	return temp
}
