package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mrcampbell/pokemon-golang/pkg/app"
	"github.com/mrcampbell/pokemon-golang/pkg/http/routes"
)

type Server struct {
	router         *gin.Engine
	pokemonService *app.PokemonService
	speciesService *app.SpeciesService
	moveService    *app.MoveService
}

func NewServer(pokemonService *app.PokemonService, speciesService *app.SpeciesService, moveService *app.MoveService) *Server {
	r := routes.GetRoutes()
	return &Server{
		router:         r,
		pokemonService: pokemonService,
		speciesService: speciesService,
		moveService:    moveService,
	}
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
