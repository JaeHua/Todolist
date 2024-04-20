package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var (
	DB *gorm.DB
)

// Todo module
type Todo struct {
	ID     int    `gorm:"primaryKey"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// initMySQL
func initMYSQL() (err error) {
	DB, err = gorm.Open("mysql", "root:jh529529@(localhost)/todolist?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return

	}
	return DB.DB().Ping()
}

func main() {
	//创建数据库
	//连接数据库
	err := initMYSQL()
	if err != nil {
		panic(err)
	}

	//延时关闭数据库
	defer DB.Close()

	//模型绑定
	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	//取todolist.html
	r.LoadHTMLFiles("template/homepage.html", "template/todolist.html")

	//告诉gin去找静态资源
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "homepage.html", nil)
	})
	r.GET("/todolist", func(c *gin.Context) {
		c.HTML(http.StatusOK, "todolist.html", nil)

	})

	//v1
	v1Group := r.Group("v1")
	{
		//代办事项
		//添加
		v1Group.POST("/todo", func(c *gin.Context) {
			//前端页面填写待办事项 点击提交 发请求到此
			//1.从请求中把数据拿出来
			var todo Todo
			c.BindJSON(&todo)

			//2.存入数据库
			if err = DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todo)
			}

		})
		//查看所有的代办事项和某一个代办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			var todoList []Todo
			if err = DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, err.Error())
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})

		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		//修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {

		})
		//删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {

		})
	}
	r.Run()
}
