package routes

import "github.com/gin-gonic/gin"

func (s Server) GetRoutes() *gin.Engine {
	r := gin.Default()
	applyHello(r.Group("/"))
	s.applyPokemon(r.Group("/api"))
	return r
}
