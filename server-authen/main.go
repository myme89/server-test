package main

import (
	"fmt"
	"server-test/server-authen/api-authen"
	"server-test/server-authen/config"
	"server-test/server-authen/database/db"
	"server-test/server-authen/database/mongodb"
	"server-test/server-authen/database/mysqldb"
)

func main() {

	config := config.GetConfig()

	// fmt.Println("log: ", config)
	switch config.Sever.TypeServer.Name {

	case "server_levedb":
		// leveldb, err := levedb.InitLeveDb(config)

		// if err != nil {
		// 	fmt.Println("Error connecting to database : error=% ", err, leveldb)
		// }
	case "server_mysql":
		_, err := mysqldb.InitMySqlDb(config)
		if err != nil {
			fmt.Println("Error connecting to database : error=% ", err)
			// logs.Logger.Error("Error connecting to mysql database : error=% ", err)
		}
	case "server_postgressql":

		db, err := db.Init(config)

		if err != nil {
			// logs.Logger.Error("Error Init to postgres database : error=% ", err)
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			// logs.Logger.Error("Error Ping to postgres database : error=% ", err)
			panic(err)
		}
	case "server_mongodb":
		mongodb.InitMongoDB(config)
	default:
		fmt.Println("Don't have sever  " + config.Sever.TypeServer.Name)
		// logs.Logger.Info("Don't have sever  " + config.Sever.TypeServer.Name)
	}

	//----------------------------------------------------------------------------//
	fmt.Println("Successfully " + config.Sever.TypeServer.Name + " connected")
	api.GRPCSeverAuthen("0.0.0.0:3002", config)
}
