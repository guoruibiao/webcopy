package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guoruibiao/webcopy/library/localcache"
	"github.com/guoruibiao/webcopy/routers"
	"log"
)

func init() {
	localcache.Init()
}

func main() {
	router := gin.Default()
	// 静态文件路由设置
	router.Static("/statics/", "./statics")
	router.Static("/templates/", "./templates")

	// 动态注册服务路由
	routers.Init(router)

	// 服务启动
	err := router.Run(":8001")
	if err != nil {
		log.Fatal(err)
	}
}
