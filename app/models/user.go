package models

import (
	"loginExample/constants"
	"loginExample/util/db"
	"time"
)

/*
  CREATE TABLE `chope_users` (
	`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	`username` varchar(64) NOT NULL,
	`password` varchar(255) NOT NULL DEFAULT '',
	`create_datetime` int(11) unsigned NOT NULL,
	PRIMARY KEY (`id`),
	UNIQUE KEY `unique_username` (`username`)
  ) ENGINE=InnoDB  DEFAULT CHARSET=utf8;
*/
type User struct {
	ID             int64  `gorm:"column:id;primary_key" json:"id"`
	Username       string `gorm:"column:username;size:64;index:unique_username,unique;not null;" json:"username"`
	Password       string `gorm:"type:varchar(255);column:password;not null;" json:"password"`
	CreateDatetime int64  `gorm:"column:create_datetime;not null;" json:"create_datetime"`
}

func (m *User) TableName() string {
	return "chope_users"
}

func (m *User) Login(username string) (error, string) {
	err := db.Get(constants.DBName).Select("password").Where("username = ?", username).Limit(1).Find(m).Error
	return err, m.Password
}

func (m *User) CreateUser(username string, passowrd string) error {
	m.Username = username
	m.Password = passowrd
	m.CreateDatetime = time.Now().Unix()
	return db.Get(constants.DBName).Create(m).Error
}
