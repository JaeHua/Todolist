package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//取todolist.html
	r.LoadHTMLFiles("template/homepage.html", "template/todolist.html")

	//告诉gin去找静态资源
	//r.Static("/static", "src/TodoList/static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "homepage.html", nil)
	})
	r.GET("/todolist", func(c *gin.Context) {
		c.HTML(http.StatusOK, "todolist.html", nil)

	})
	r.Run()
}
