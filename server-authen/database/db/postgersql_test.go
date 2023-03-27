package db

import (
	"log"
	"server-test/server-authen/config"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	path := "./../../config/config.yml"
	config, err := config.NewConfig(path)
	if err != nil {
		log.Fatal("Get congfig err")
	}
	_, err = Init(config)

	require.NoError(t, err)
}
