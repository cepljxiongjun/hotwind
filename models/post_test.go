package models

import (
	"testing"

	"github.com/spf13/viper"
)

func add(t *testing.T) {
	viper.SetConfigFile("../config/in-local.toml")
	err := viper.ReadInConfig()
	if err != nil {
		t.Error(err)
	}
	Init()

}
