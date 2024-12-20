package routes

import "github.com/gin-gonic/gin"

func GetRoutes() *gin.Engine {
	r := gin.Default()
	applyHello(r.Group("/"))
	return r
}
