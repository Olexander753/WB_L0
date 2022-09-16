package handler

import (
	"log"
	"net/http"

	"github.com/Olexander753/WB_L0/internal/schema"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getModelByID(c *gin.Context) {
	var input schema.Model
	input.Order_uid = c.Param("uid")

	output, err := h.services.SelectModel(c, input.Order_uid)
	output.Order_uid = input.Order_uid

	if err != nil {
		log.Fatal("error get model: ", err)
	}

	log.Println("Get model, order_uid = ", output.Order_uid)

	c.JSON(http.StatusOK, output)
}
