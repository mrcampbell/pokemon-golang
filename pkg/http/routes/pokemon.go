package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s Server) applyPokemon(r *gin.RouterGroup) {
	r.GET("/pokemon", s.listPokemon)
	r.GET("/pokemon/:id", s.getPokemon)
	r.POST("/pokemon", s.createPokemon)
}

type createPokemonRequest struct {
	SpeciesID int `json:"species_id"`
	Level     int `json:"level"`
}

func (s Server) listPokemon(c *gin.Context) {
	ctx := context.Background()
	p, err := s.pokemonService.ListPokemon(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, p)
}

func (s Server) getPokemon(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
	}

	p, err := s.pokemonService.GetPokemon(ctx, uid)
	if err != nil {
		// todo: wrap/type this error, instead of flaky string comparison
		if err.Error() == "error getting pokemon: no rows in result set" {
			fmt.Println("NO ROWS")
			c.JSON(http.StatusNotFound, gin.H{"error": "pokemon not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, p)
}

func (s Server) createPokemon(c *gin.Context) {
	ctx := context.Background()
	var req createPokemonRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p, err := s.pokemonService.CreatePokemon(ctx, req.SpeciesID, req.Level)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, p)
}
