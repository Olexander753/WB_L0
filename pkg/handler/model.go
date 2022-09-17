package handler

import (
	"fmt"
	"html/template"
	"log"

	"github.com/Olexander753/WB_L0/internal/schema"
	"github.com/gin-gonic/gin"
)

type ViewData struct {
	Title   string
	Message string
}

func (h *Handler) getModelByID(c *gin.Context) {
	var input schema.Model
	input.Order_uid = c.Param("uid")

	output, err := h.services.SelectModel(c, input.Order_uid)
	output.Order_uid = input.Order_uid

	if err != nil {
		log.Fatal("error get model: ", err)
	}

	log.Println("Get model, order_uid = ", output.Order_uid)

	data := ViewData{
		Title:   "Model",
		Message: output.PrintModel(),
	}
	//fmt.Println(data.Message)
	parce := fmt.Sprintf("<html><head> <meta charset=\"utf-8\"><style>li {list-style-type: none;}</style></head><body><div> <h1>%s</h1> <p>%s</p> </div>", data.Title, data.Message)
	tmpl := template.Must(template.New("data").Parse(parce))
	tmpl.Execute(c.Writer, data)

	//c.JSON(http.StatusOK, output)
}
