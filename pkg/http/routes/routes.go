package routes

import "github.com/gin-gonic/gin"

func (s Server) GetRoutes() *gin.Engine {
	r := gin.Default()
	applyHello(r.Group("/"))
	applyClient(r)
	s.applyPokemon(r.Group("/api"))
	s.applySpecies(r.Group("/api"))
	return r
}
