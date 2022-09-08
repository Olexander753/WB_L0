package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/Olexander753/WB_L0/internal/db"
	"github.com/Olexander753/WB_L0/internal/event"
	"github.com/Olexander753/WB_L0/internal/schema"
	"github.com/Olexander753/WB_L0/pkg/util"
)

func createModelHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		ID string `json:"id"`
	}
	ctx := r.Context()

	body := template.HTMLEscapeString(r.FormValue("body"))
	if len(body) < 1 || len(body) > 140 {
		util.ResponseError(w, http.StatusBadRequest, "Invalid body")
		return
	}

	//create model
	model := schema.Model{
		Order_uid: "dasdasdasd", //TODO
	}

	if err := db.InsertModel(ctx, model); err != nil {
		log.Println(err)
		util.ResponseError(w, http.StatusInternalServerError, "Failed to create meow")
		return
	}

	// Publish event
	if err := event.PublishModelCreated(model); err != nil {
		log.Println(err)
	}

	// Return new meow
	util.ResponseOk(w, response{ID: model.Order_uid})
}
