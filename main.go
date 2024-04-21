package main

import (
	"JTools/dao"
	"JTools/model"
	"JTools/router"
)

func main() {

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
	r.Run()
}
