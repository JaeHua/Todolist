package controller

import (
	"ToolWeb/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

/*
	控制模块
*/

func HomepageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "homepage.html", nil)
}
func TodolistHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "todolist.html", nil)
}
func CreateTodo(c *gin.Context) {
	//前端页面填写待办事项 点击提交 发请求到此
	//1.从请求中把数据拿出来
	var todo model.Todo
	c.BindJSON(&todo)

	//2.存入数据库
	err := model.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}

}
func GetAllTodo(c *gin.Context) {

	todoList, err := model.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateTodo(c *gin.Context) {
	id := strings.TrimPrefix(c.Param("id"), ":") //奇奇怪怪去除冒号

	todo, err := model.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}
	// 切换状态(主要是点击可以切换完成与未完成)
	if todo.Status {
		todo.Status = false
	} else {
		todo.Status = true
	}
	if err := model.UpdateTodo(todo); err != nil {
		c.JSON(http.StatusOK, err.Error())
	} else {
		c.JSON(http.StatusOK, todo)
	}

}
func DeleteATodo(c *gin.Context) {
	id := strings.TrimPrefix(c.Param("id"), ":") //奇奇怪怪去除冒号

	if err := model.DeleteTodo(id); err != nil {
		c.JSON(http.StatusOK, err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}

}
