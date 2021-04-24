package teacher

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers(e *gin.Engine) {
	e.GET("/v1/goods/list", goodsHandler)
	e.GET("/checkout", checkoutHandler)
}

func goodsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello goodsHandler",
	})
}

func checkoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello checkoutHandler",
	})
}