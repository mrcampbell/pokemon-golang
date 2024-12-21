package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s Server) applySpecies(r *gin.RouterGroup) {
	r.GET("/species/:id", s.getSpecies)
}

func (s Server) getSpecies(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id. must be an integer between 1 and 1025"})
		return
	}
	species := s.speciesService.GetSpecies(idInt)
	c.JSON(http.StatusOK, species)
}
