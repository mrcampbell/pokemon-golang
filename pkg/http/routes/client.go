package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func applyClient(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")
	r.GET("/", index)
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}
