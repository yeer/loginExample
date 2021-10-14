package main

import (
	"fmt"
	"loginExample/cmd"
)

const (
	buildDate  = "2021-10-13"
	commitHash = "HEAD"
)

func main() {
	//globle panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic:", err)
			return
		}
	}()
	//start
	cmd.Execute()
}
