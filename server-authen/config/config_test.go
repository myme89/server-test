package config

import (
	"log"
	"testing"
)

func TestGetConfig(t *testing.T) {
	path := "./config.yml"
	_, err := NewConfig(path)

	if err != nil {
		log.Fatal("Get congfig err")
	}

}
