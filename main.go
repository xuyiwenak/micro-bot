package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	"user-web/handler"
)

var (
	dockerMode string
)

func main() {
	// 创建新服务
	service := web.NewService(
		// 后面两个web，第一个是指是web类型的服务，第二个是服务自身的名字
		web.Name("bambooRat.micro.web.bot"),
		web.Version("latest"),
		web.Address(":8088"),
	)
	// 初始化服务
	if err := service.Init(
		web.Action(
			func(c *cli.Context) {
				// 初始化handler
				handler.Init()
			}),
	); err != nil {
		log.Fatal(err)
	}

	// 注册登录接口
	service.HandleFunc("/bot/push", handler.Login)

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
