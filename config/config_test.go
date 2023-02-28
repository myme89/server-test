package config

import (
	"log"
	"testing"
)

func TestGetConfig(t *testing.T) {
	path := "/home/nhatnt/nhatnt/probationary-project/server-test/config/config.yml"
	_, err := NewConfig(path)

	if err != nil {
		log.Fatal("Get congfig err")
	}

}
