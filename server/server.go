package server

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"loginExample/util/cache"
	"loginExample/util/db"
	"loginExample/util/logger"
	// "loginExample/app/controllers"
)

var server *gin.Engine

var registers = [...]func() error{
	logger.Register,
	db.Register,
	cache.Register,
}

func Recover(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("panic:", err)
			logger.Error("stack:", string(debug.Stack()))
			ctx.String(http.StatusBadGateway, "internal error")
		}
	}()
	ctx.Next()
}

func Logger(ctx *gin.Context) {
	start := time.Now()
	path := ctx.Request.URL.Path
	query := ctx.Request.URL.RawQuery
	ctx.Next()
	end := time.Now()
	latency := end.Sub(start)
	if len(ctx.Errors) > 0 {
		for _, e := range ctx.Errors.Errors() {
			logger.Error(e)
		}
	}
	content := fmt.Sprintf("%d %s %s %s %s %dms",
		ctx.Writer.Status(),
		ctx.Request.Method,
		path,
		query,
		ctx.ClientIP(),
		latency.Milliseconds(),
	)
	logger.Debug(content)
}

func Init() (err error) {
	for _, register := range registers {
		err = register()
		if err != nil {
			return
		}
	}
	if viper.GetString("server.mode") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	if viper.GetString("server.domain") != "" {
		constants.DOMAIN = viper.GetString("server.domain")
	}
	server = gin.New()
	server.Use(Logger)
	server.Use(Recover)
	return nil
}

func setupRouter() {
	server.GET("/", util.Handle(controllers.MainIndex))
	server.GET("/:path/:url", util.Handle(controllers.UUrlMainGet))
	server.POST("/u-url/:path/create", util.Handle(controllers.UUrlMainCreate))
}

func Run() {
	setupRouter()
	addr := fmt.Sprintf(":%s", viper.GetString("server.port"))
	err := server.Run(addr)
	if err != nil {
		panic("The web service failed to start!")
	}
}
