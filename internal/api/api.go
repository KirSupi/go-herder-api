package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kirsupi/go-herder"
	"go-herder-api/internal/config"
	"strconv"
)

type API struct {
	c config.ApiConfig
	s *gin.Engine
	h *herder.Herder
}

func New(c config.Config) *API {
	api := &API{
		c.Api,
		gin.Default(),
		herder.New(c.Herder),
	}
	api.initHandlers()
	return api
}

func (api *API) Run() error {
	go api.h.Run()
	return api.s.Run(api.c.Host + ":" + strconv.Itoa(api.c.Port))
}
