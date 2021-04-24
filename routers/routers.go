package routers

import (
	"ginDemo/common"
	"github.com/gin-gonic/gin"
)

type Opetion func(engine *gin.Engine)

var options = []Opetion{}

//注册APP路由配置
func Include(opts ...Opetion)  {
	options = append(options,opts...)
}

//初始化
func Init() *gin.Engine  {
	//创建一个无中间件路由
	r := gin.New()

	//注册全局中间件
	r.Use(gin.Logger())
	r.Use(common.MiddleWare2)

	for _, opt := range options {
		opt(r)
	}
	return r
}