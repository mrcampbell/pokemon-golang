package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kr/pretty"
)

func (s Server) applyPokemon(r *gin.RouterGroup) {
	r.GET("/pokemon", s.listPokemon)
}

func (s Server) listPokemon(c *gin.Context) {
	ctx := context.Background()
	p, err := s.pokemonService.ListPokemon(ctx)
	pretty.Println(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, p)
}
