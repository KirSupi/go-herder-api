package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kirsupi/go-herder"
	"net/http"
	"strconv"
)

func (api *API) initHandlers() {
	apiGroup := api.s.Group("/api")
	{
		herderGroup := apiGroup.Group("/herder")
		herderGroup.GET("/", api.herderGet)
		{
			queueGroup := herderGroup.Group("/queue")
			queueGroup.GET("/", api.herderQueueGet)
			queueGroup.POST("/", api.herderQueuePost)
			queueGroup.DELETE("/", api.herderQueueDelete)
			queueGroup.DELETE("/:id", api.herderQueueDeleteWithId)
		}
		{
			activeGroup := herderGroup.Group("/active")
			activeGroup.GET("/", api.herderActiveGet)
			activeGroup.DELETE("/:id", api.herderActiveDeleteWithId)
		}
		{
			queueGroup := herderGroup.Group("/finished")
			queueGroup.GET("/", api.herderFinishedGet)
		}
	}
}

func (api *API) herderGet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Ok.WithData(api.h.GetAllStates()))
}

func (api *API) herderQueueGet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Ok.WithData(api.h.GetQueue()))
}
func (api *API) herderQueuePost(ctx *gin.Context) {
	var tc herder.TaskConfig
	if err := ctx.BindJSON(&tc); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorBadRequest)
	}
	ctx.JSON(http.StatusOK, Ok.WithData(api.h.AddToQueue(tc)))
}
func (api *API) herderQueueDelete(ctx *gin.Context) {
	api.h.ClearQueue()
	ctx.JSON(http.StatusOK, Ok)
}
func (api *API) herderQueueDeleteWithId(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorBadRequest)
		return
	}
	err = api.h.RemoveFromQueue(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorInternalServer)
		return
	}
	ctx.JSON(http.StatusOK, Ok)
}

func (api *API) herderActiveGet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Ok.WithData(api.h.GetActive()))
}
func (api *API) herderActiveDeleteWithId(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorBadRequest)
		return
	}
	err = api.h.Kill(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorInternalServer)
		return
	}
	ctx.JSON(http.StatusOK, Ok)
}

func (api *API) herderFinishedGet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Ok.WithData(api.h.GetFinished()))
}
