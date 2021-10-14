package main

import (
	"loginExample/app/service"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestRegister(t *testing.T) {
	//curl -H "Content-Type: application/json" -X POST -d 'username=foo&password=xx22@4' "http://127.0.0.1:10010/v1/users/create"
	service := new(service.User)
	result := service.CreateUser("a1", "xx22@4")
	assert.Equal(t, result, nil)
}
func TestLogin(t *testing.T) {
	//curl -X POST -d 'username=foo&password=xx22@4' "http://127.0.0.1:10010/v1/login"
	service := new(service.User)
	result := service.Login("a1", "xx22@4")
	assert.Equal(t, result, nil)
}
