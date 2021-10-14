package main

import (
	"fmt"
	"loginExample/app/service"
	"loginExample/util/cache"
	"loginExample/util/db"
	"loginExample/util/logger"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("./conf/conf.yaml")
	//初始化全部的配置
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//初始化全部的配置
	_ = logger.Register()
	_ = db.Register()
	_ = cache.Register()
}

func TestRegister(t *testing.T) {
	//curl -X POST -d 'username=foo&password=xx22@4' "http://127.0.0.1:10010/v1/users/create"
	service := new(service.User)
	result := service.CreateUser("a4", "xx22@4")
	assert.Equal(t, result, nil)

	result = service.CreateUser("user1111111111111111111211", "xx8*38737823@@**#")
	assert.Equal(t, result, nil)
}
func TestLogin(t *testing.T) {
	//curl -X POST -d 'username=foo&password=xx22@4' "http://127.0.0.1:10010/v1/login"
	service := new(service.User)
	result := service.Login("a4", "xx22@4")
	assert.Equal(t, result, nil)
}
