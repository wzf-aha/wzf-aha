package student

import (
	"ginDemo/app/student/controller"
	"ginDemo/common"
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine)  {
	e.GET("/get-books-list/:user/:password",controller.GetBooksList)
	e.GET("/books/books-info", controller.BooksInfo)
	//给test-middle添加局部中间件
	e.GET("/books/test-middle", common.MyTime,controller.TestMiddle)
	e.GET("/books/test-person",controller.TestPerson)
}