package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "c", "./config.toml", "配置文件路径")
	flag.Parse()
}

func main() {
	engine := gin.Default()

	engine.GET("/echo", func(ctx *gin.Context) {
		ctx.String(200, "Hello ChatBot Echo!\nv2\n\nConfig Path: %s\n", configPath)
	})

	if err := engine.Run(":8080"); err != nil {
		fmt.Println(err.Error())
	}
}
