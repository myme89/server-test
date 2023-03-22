package cmd

import (
	"context"
	"fmt"
	"server-test/server-gateway/api"
	"server-test/server-gateway/cache"
	// log "github.com/sirupsen/logrus"
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

	// logs.InitLogger()

	ctx := context.TODO()
	cache.ConnectRedis(ctx)

	// values := cache.GetAllKeys(ctx, "id*")

	// logs.Logger.Info("All values : %v \n", values)

	serverAddr := "0.0.0.0:3000"
	// serverAddrhttp := "0.0.0.0:3003"

	// go api.GRPCSever(serverAddrhttp)
	api.GatewaySever(serverAddr)
}
