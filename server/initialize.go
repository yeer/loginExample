package server

import (
	"fmt"
	"loginExample/app/models"
	"loginExample/constants"
	"loginExample/util/db"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDatabase() {
	createSQL := fmt.Sprintf(
		"CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8;",
		constants.DBName,
	)
	config := viper.GetStringMapString("component.db.user")
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=True&loc=Local",
		config["user"],
		config["password"],
		config["host"],
		config["port"])
	dbadmin, err := gorm.Open(mysql.Open(connStr))
	if err != nil {
		panic("conn database error.")
	}
	err = dbadmin.Exec(createSQL).Error
	if err != nil {
		panic("create database error.")
	}
}

func InitializeTables() {
	if db.Get(constants.DBName).AutoMigrate(&models.User{}) != nil {
		panic("iInitialize database error.")
	} else {
		fmt.Println("iInitialize database")
	}
}
