package levedb

import (
	"encoding/json"
	"fmt"
	"log"
	"server-test/config"
	"server-test/logs"
	"server-test/model"

	"github.com/syndtr/goleveldb/leveldb"
)

type Info struct {
	Name     string `json:"name"`
	FullName string `json:"fullname"`
}

var db *leveldb.DB
var err error

func PutData(config *config.Config, Id int, Name, FullName string) error {

	db, err = leveldb.OpenFile(config.Sever.ServerLevelDB.PathFile, nil)
	if err != nil {
		logs.Logger.Fatal("PutData - Open file leveDb error: ", err)
		log.Fatal("Yikes!")
	}
	defer db.Close()

	byteArray, err := json.Marshal(model.DataPost{Id: Id, Name: Name, FullName: FullName})
	if err != nil {
		logs.Logger.Error("PutData - Marshal data leveDb error: ", err)
		fmt.Println(err)
	}

	err = db.Put([]byte("data_test1"), byteArray, nil)
	if err != nil {
		logs.Logger.Error("PutData - Put data leveDb error: ", err)
		fmt.Println(err)
	}

	fmt.Println(byteArray)
	logs.Logger.Info("PutData - data leveDb: ", byteArray)

	return err
}

func GetData(config *config.Config) model.DataInfo {

	db, err = leveldb.OpenFile(config.Sever.ServerLevelDB.PathFile, nil)
	if err != nil {
		logs.Logger.Fatal("GetData - Open file leveDb error: ", err)
		log.Fatal("Yikes!")
	}
	defer db.Close()

	var temp model.DataPost
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		// key := iter.Key()
		value := iter.Value()

		err := json.Unmarshal([]byte(value), &temp)
		if err != nil {
			logs.Logger.Panic("GetData - Unmarshal Data leveDb error: ", err)
			panic(err)
		}

	}

	return model.DataInfo{Name: temp.Name, FullName: temp.FullName}
}
