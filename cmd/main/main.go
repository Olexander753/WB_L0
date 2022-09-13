package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/Olexander753/WB_L0/internal/db"
	"github.com/Olexander753/WB_L0/internal/event"
	"github.com/Olexander753/WB_L0/internal/schema"
	"github.com/Olexander753/WB_L0/pkg/config"
	"github.com/Olexander753/WB_L0/pkg/util"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Read config")
	cfg := config.GetConfig()
	fmt.Println(cfg)

	//connect to postgresql db
	log.Println("connect to postgresql db")
	repo, err := db.NewPostgres(cfg)
	if err != nil {
		log.Panicln("Failed connect to db, error: ", err)
		return
	}
	db.SetRepository(repo)
	defer db.Close()

	//connect to nats
	log.Println("connect to nats")
	es, err := event.NewNats(fmt.Sprintf("nats://%s", cfg.Nats.Address))
	if err != nil {
		log.Println("Failed connect to nats, error: ", err)
	}
	event.SetEventStore(es)
	defer event.Close()

	//create http router
	log.Println("create http router")
	router := newRouter()
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Failed create http router, error: ", err)
	}
}

func newRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/models", createModelHandler).
		Methods("POST").
		Queries("body", "{body}")
	return
}

func createModelHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		ID string `json:"id"`
	}
	ctx := r.Context()

	//Read parametrs
	body := template.HTMLEscapeString(r.FormValue("body"))
	if len(body) < 1 {
		util.ResponseError(w, http.StatusBadRequest, "Invalid body")
		return
	}

	//Create model
	model := schema.Model{
		Order_uid: "1ewqeqxa132",
	}

	if err := db.InsertModel(ctx, model); err != nil {
		log.Println(err)
		util.ResponseError(w, http.StatusInternalServerError, "Failed to create meow")
		return
	}

	//Publish event
	if err := event.PublishModelCreated(model); err != nil {
		log.Println(err)
	}

	//Return new model
	util.ResponseOk(w, response{ID: model.Order_uid})

}
