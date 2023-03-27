package levedb

import (
	"fmt"
	"log"
	"server-test/server-authen/config"
	"testing"
)

func TestGetData(t *testing.T) {
	path := "./../../config/config.yml"
	config, err := config.NewConfig(path)
	if err != nil {
		log.Fatal("Get congfig err")
	}

	temp := GetData(config)
	fmt.Println(temp)
}

func TestPutData(t *testing.T) {
	Id := 2
	Name := "nhat1"
	FullName := "NTN2"

	path := "./../../config/config.yml"
	config, err := config.NewConfig(path)
	if err != nil {
		log.Fatal("Get congfig err")
	}

	err = PutData(config, Id, Name, FullName)
	if err != nil {
		log.Fatal("Put data failed")
	}
}
