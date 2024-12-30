package main

import (
	"github.com/kelin428/hengo"
	"github.com/kelin428/hengo/log"
)

type Name struct {
	Id string `json:"id"`
}

func main() {
	hen := hengo.New()
	hen.Get("/hello", func(c *hengo.Context) {
		var name Name
		c.BindJSON(&name)

		// 创建日志配置
		logConfig := log.NewConfig(false)
		logger := logConfig.LogHelp("log/log.go:30", "server/LocalHttpRequestFilter")
		logger.Info("test")

		// 返回json数据
		c.JsonResponse(hengo.StatusOK, "Hello World", hengo.H{
			"name":    "Hen",
			"version": "0.1",
		})
	})

	hen.Run(":8000")
}
