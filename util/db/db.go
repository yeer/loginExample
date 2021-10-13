package db

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/gorm/logger"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbs = map[string]*gorm.DB{}
var logLogger = log.New(os.Stdout, "\r\n", log.LstdFlags)

var RowNULL = fmt.Errorf("empty result")
var LogLevelList = map[string]logger.LogLevel{
	"silent": logger.Silent,
	"error":  logger.Error,
	"warn":   logger.Warn,
	"info":   logger.Info,
}

func Register() error {
	configs := viper.GetStringMap("component.db")
	for key := range configs {
		config := viper.GetStringMapString("component.db." + key)
		connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config["user"],
			config["password"],
			config["host"],
			config["port"],
			config["database"])

		newLogger := GetDBLogger(config)
		db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{
			Logger: newLogger,
		})
		if err != nil {
			return err
		}
		if config["debug"] == "true" {
			db = db.Debug()
		}
		dbs[key] = db
	}
	return nil
}

func Get(key string) *gorm.DB {
	db, ok := dbs[key]
	if !ok {
		panic("db配置不存在:" + key)
	}
	return db
}

func GetDBLogger(config map[string]string) logger.Interface {
	slowThreshold := 200 //默认慢sql阈值, 200毫秒
	if _, ok := config["slow_threshold"]; ok {
		slowThreshold, _ = strconv.Atoi(config["slow_threshold"])
	}
	logLevel := logger.Warn //默认报警级别
	if _, ok1 := config["log_level"]; ok1 {
		if _, ok2 := LogLevelList[config["log_level"]]; ok2 {
			logLevel = LogLevelList[config["log_level"]]
		}
	}
	return logger.New(
		logLogger, // io writer
		logger.Config{
			SlowThreshold: time.Duration(slowThreshold) * time.Millisecond, // 慢 SQL 阈值
			LogLevel:      logLevel,                                        // Log level
			Colorful:      true,                                            // 禁用彩色打印
		},
	)
}

func SetLogger(Log *log.Logger) {
	logLogger = Log
}
