package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

/**
中间件的流转
	gin提供了两个函数Abort()和Next()，配合着return关键字用来跳转或者终止存在着业务逻辑关系的中间件
	abort()就是终止该中间件的流程，如果不return的话会继续执行后面的逻辑，但不再执行其他的中间件。next()跳过当前的中间件，执行下一个中间件，待下一个中间件执行完后再回到当前next位置，直接后面的逻辑
 */

//定义中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		log.Println("中间件开始执行")
		c.Set("request","中间件")
		status := c.Writer.Status()
		log.Println("中间件执行结束：",status)
		t2 := time.Since(t)
		log.Println("time:",t2)
	}
}

//定义中间件
func MiddleWare2(c *gin.Context)  {
	t := time.Now()
	log.Println("中间件开始执行")
	// 设置变量到Context的key中，可以通过Get()取
	c.Set("request","中间件")
	// 跳过当前中间件，执行函数
	c.Next()
	// 中间件执行完后续的一些事情
	status := c.Writer.Status()
	log.Println("中间件执行结束：",status)
	t2 := time.Since(t)
	log.Println("time:",t2)
}

//计算程序用时
func MyTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	// 统计时间
	since := time.Since(start)
	fmt.Println("程序用时：", since)
}
