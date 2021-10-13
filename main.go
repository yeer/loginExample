package main

import (
	"fmt"
)

const (
	buildDate  = "2021-10-13"
	commitHash = "HEAD"
)

func main() {
	//捕获全局panic异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic异常——", err)
			return
		}
	}()
	//命令初始化
	cmd.Execute()
}
