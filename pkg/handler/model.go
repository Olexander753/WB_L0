package handler

import (
	"log"
	"net/http"

	"github.com/Olexander753/WB_L0/internal/schema"
	"github.com/gin-gonic/gin"
)

type getResponse struct {
	Data []schema.Model `json:"data"`
}

func (h *Handler) postModel(c *gin.Context) {
	var input schema.Model

	if err := c.BindJSON(&input); err != nil {
		log.Fatal("error input 1: ", err)
	}

	// id, err := h.db.InsertModel(input)
	// if err != nil {
	// 	log.Fatal("error insert: ", err)
	// }

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "model успешно создана",
	})
}

func (h *Handler) getModelsByID(c *gin.Context) {
	var input schema.Model
	var list []schema.Model

	if err := c.BindJSON(&input); err != nil {
		log.Fatal("error input 3: ", err)
	}

	// list, err := h.services.TransactionsList.GetListByID(input.User_id)

	// if err != nil {
	// 	log.Fatal("error get list 1: ", err)
	// }

	c.JSON(http.StatusOK, getResponse{
		Data: list,
	})
}
