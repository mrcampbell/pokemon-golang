package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrcampbell/pokemon-golang/pkg/app"
)

type Server struct {
	router         *gin.Engine
	pokemonService app.PokemonService
	speciesService app.SpeciesService
	moveService    app.MoveService
}

func NewServer(pokemonService app.PokemonService, speciesService app.SpeciesService, moveService app.MoveService) *Server {
	s := &Server{
		pokemonService: pokemonService,
		speciesService: speciesService,
		moveService:    moveService,
	}

	s.router = s.GetRoutes()
	return s
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
