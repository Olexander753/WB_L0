package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Olexander753/WB_L0/internal/schema"
	"github.com/gin-gonic/gin"
)

func (h *Handler) postModel(c *gin.Context) {
	var input schema.Model

	if err := c.BindJSON(&input); err != nil {
		log.Fatal("error input 1: ", err)
	}

	uid, err := h.services.InsertModel(c, input)
	if err != nil {
		log.Fatal("error insert: ", err)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"uid":     uid,
		"message": "model успешно создана",
	})
}

func (h *Handler) getModelByID(c *gin.Context) {
	var input schema.Model
	input.Order_uid = c.Param("uid")

	output, err := h.services.SelectModel(c, input.Order_uid)
	output.Order_uid = input.Order_uid

	if err != nil {
		log.Fatal("error get list 1: ", err)
	}

	fmt.Println(output)

	c.JSON(http.StatusOK, output)
}
