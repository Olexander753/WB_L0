package handler

import (
	"github.com/Olexander753/WB_L0/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InutRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/", h.postModel)    //создание модели
		api.GET("/", h.getModelsByID) //вывод моделей
	}
	return router
}
