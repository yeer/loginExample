package service

import (
	"fmt"
	"loginExample/app/models"
	"loginExample/constants"
	"loginExample/util/logger"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	error error
}

func (u *User) GetErrors() error {
	return u.error
}

func (u *User) CreateUser(username string, password string) error {
	var err error
	models := new(models.User)
	hash_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		u.error = err
		logger.Error(constants.StatusText(constants.StatusCreateFailure), err.Error())
	}

	err = models.CreateUser(username, string(hash_password))
	if err != nil {
		u.error = err
		logger.Error(constants.StatusText(constants.StatusCreateFailure), err.Error())
	}
	return err
}

func (u *User) Login(username string, password string) error {
	var err error
	models := new(models.User)
	err, hash_password := models.Login(username)
	fmt.Println(hash_password)
	err = bcrypt.CompareHashAndPassword([]byte(hash_password), []byte(password)) //验证（对比）
	if err != nil {
		u.error = err
		logger.Error(constants.StatusText(constants.StatusDbNoResult), err.Error())
	}
	return err
}
