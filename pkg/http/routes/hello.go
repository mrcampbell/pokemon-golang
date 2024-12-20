package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func applyHello(r *gin.RouterGroup) {
	r.GET("/hello", hello)
}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "hi")
}
