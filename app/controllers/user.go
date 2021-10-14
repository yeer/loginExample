package controllers

import (
	"loginExample/app/service"
	"loginExample/constants"
	"loginExample/util/response"
)

func Index(ctx *response.Context) {
	ctx.JsonError(constants.StatusOk, "This is Login Example")
}

// user login
func Login(ctx *response.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	if username == "" || password == "" {
		ctx.JsonError(constants.StatusParamsError, "")
		return
	}

	service := new(service.User)
	if service.Login(username, password) != nil {
		ctx.JsonError(constants.StatusLoginFailure, constants.StatusText(constants.StatusLoginFailure))
		return
	}
	ctx.JsonSuccess(constants.StatusOk)
}

// register user
func Register(ctx *response.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	if username == "" || password == "" {
		ctx.JsonError(constants.StatusParamsError, "")
		return
	}

	service := new(service.User)
	if service.CreateUser(username, password) != nil {
		ctx.JsonError(constants.StatusCreateFailure, constants.StatusText(constants.StatusCreateFailure))
		return
	}
	ctx.JsonSuccess(constants.StatusOk)
}
