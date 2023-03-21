package mysqldb

import (
	"log"
	"server-test/server-database/config"
	"server-test/server-database/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	path := "./../../config/config.yml"
	config, err := config.NewConfig(path)
	if err != nil {
		log.Fatal("Get congfig err")
	}
	_, err = InitMySqlDb(config)

	require.NoError(t, err)
}

func TestGetData(t *testing.T) {
	path := "./../../config/config.yml"
	config, err := config.NewConfig(path)
	_, err = InitMySqlDb(config)
	listData := []model.DataPost{
		{Id: 1, Name: "Billy", FullName: "nhat24"},
		{Id: 35, Name: "Billy2", FullName: "nhat2423"},
	}
	err = GetData(&listData)
	require.NoError(t, err)
}

func TestPostData(t *testing.T) {
	path := "./../../config/config.yml"
	config, err := config.NewConfig(path)
	_, err = InitMySqlDb(config)
	err = PostData(1934, "nhat432", "NTN4441")
	require.NoError(t, err)
}
