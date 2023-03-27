package mongodb

import (
	"fmt"
	"server-test/server-storage/config"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

// var test *mongo.Client

func TestInit(t *testing.T) {
	path := "./../../config/config.yml"
	config, err := config.NewConfig(path)
	if err != nil {
		log.Fatal("Get congfig err")
	}
	clientMongo := InitMongoDB(config)
	fmt.Println(clientMongo)
}

func TestGetAllData(t *testing.T) {
	path := "./../../config/config.yml"
	config, err := config.NewConfig(path)
	if err != nil {
		log.Fatal("Get congfig err")
	}
	clientMongo := InitMongoDB(config)
	fmt.Println(clientMongo)

	temp, err := GetAllInfo(config, "")
	if err != nil {
		log.Fatal("Get data  failed ", err)
	}
	fmt.Println(temp)
}

func TestAddinfo(t *testing.T) {
	path := "./../../config/config.yml"
	config, err := config.NewConfig(path)
	if err != nil {
		log.Fatal("Get congfig err")
	}
	clientMongo := InitMongoDB(config)
	fmt.Println(clientMongo)

	err = AddInfo(config, 2, "nhat", "NTN")
	if err != nil {
		log.Fatal("Get data  failed ", err)
	}
	require.NoError(t, err)
}
