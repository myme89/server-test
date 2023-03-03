package cmd

import (
	"context"
	"fmt"
	"server-test/api"
	"server-test/cache"
	"server-test/config"
	"server-test/database/db"
	"server-test/database/mongodb"
	"server-test/database/mysqldb"

	log "github.com/sirupsen/logrus"
)

func Version(version string) {

	fmt.Println("Version: ", version)
}

func Help() {
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("			./server-test  <command>						")
	fmt.Println()

	fmt.Println("The commands are:")
	fmt.Println("\t\t\t version		:		Get version					")
	fmt.Println("\t\t\t run			:		Run server  				")
}

func Run() {
	config := config.GetConfig()

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
		}
	case "server_postgressql":

		db, err := db.Init(config)

		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			panic(err)
		}
	case "server_mongodb":
		mongodb.InitMongoDB(config)
	default:
		fmt.Println("Don't have sever  " + config.Sever.TypeServer.Name)
	}

	//----------------------------------------------------------------------------//
	fmt.Println("Successfully " + config.Sever.TypeServer.Name + " connected")
	ctx := context.TODO()
	cache.ConnectRedis(ctx)

	values := cache.GetAllKeys(ctx, "id*")

	log.Info("All values : %v \n", values)

	serverAddr := "0.0.0.0:3000"
	// serverAddrhttp := "0.0.0.0:3003"

	// go api.GRPCSever(serverAddrhttp)
	api.GatewaySever(serverAddr, config)
}
