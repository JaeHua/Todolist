package router

import (
	"JTools/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//取todolist.html
	r.LoadHTMLFiles("template/homepage.html", "template/todolist.html")

	//告诉gin去找静态资源
	r.Static("/static", "./static")

	r.GET("/", controller.HomepageHandler)
	r.GET("/todolist", controller.TodolistHandler)

	//v1
	v1Group := r.Group("v1")
	{
		//代办事项
		//添加
		v1Group.POST("/todo", controller.CreateTodo)
		//查看所有的代办事项和某一个代办事项
		v1Group.GET("/todo", controller.GetAllTodo)
		//修改
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
