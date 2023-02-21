package main

import (
	"context"
	"fmt"
	"server-test/api"
	"server-test/cache"
	"server-test/db"

	log "github.com/sirupsen/logrus"
)

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("Hello Trong Nhat")

	db.Init()

	ctx := context.TODO()
	cache.ConnectRedis(ctx)

	values := cache.GetAllKeys(ctx, "id*")

	log.Info("All values : %v \n", values)

	serverAddr := "0.0.0.0:3000"
	// serverAddrhttp := "0.0.0.0:3003"

	// go api.GRPCSever(serverAddrhttp)
	api.GatewaySever(serverAddr)

}
