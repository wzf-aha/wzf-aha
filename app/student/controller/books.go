package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// 定义接收数据的结构体
type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func GetBooksList(c *gin.Context)  {
	// 声明接收的变量
	var json Login
	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindUri(&json); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 判断用户名密码是否正确
	if json.User != "root" || json.Pssword != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

func BooksInfo(c *gin.Context)  {
	copyContext := c.Copy()
	log.Println("开始")
	//异步处理
	go func() {
		time.Sleep(10 * time.Second)
		log.Println("异步执行10S：" + copyContext.Request.URL.Path)
	}()
	log.Println("结束")
}

func TestMiddle(c *gin.Context)  {
	// 取值
	req, _ := c.Get("request")
	log.Println("request:", req)
	// 页面接收
	c.JSON(200, gin.H{"request": req})
}

//结构体数据验证
type Person struct {
	//不能为空并且大于10
	Age      int       `form:"age" binding:"required,gt=10"`
	Name     string    `form:"name" binding:"required"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func TestPerson(c *gin.Context)  {
	var person Person
	if err := c.ShouldBind(&person); err != nil {
		c.String(500, fmt.Sprint(err))
		return
	}
	c.String(200, fmt.Sprintf("%#v", person))
}