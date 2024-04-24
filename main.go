package main

import (
	"ToolWeb/conf"
	"ToolWeb/dao"
	"ToolWeb/model"
	"ToolWeb/router"
	"fmt"
)

func main() {
	//初始化配置
	config := conf.InitConf()

	fmt.Printf(config.AppConf.Port)
	//创建数据库
	//连接数据库
	err := dao.InitMYSQL()
	if err != nil {
		panic(err)
	}

	//延时关闭数据库
	defer dao.DB.Close()

	//模型绑定
	dao.DB.AutoMigrate(&model.Todo{})

	//路由注册
	r := router.SetupRouter()

	//运行服务
	r.Run(config.AppConf.Port)
}
