package main

import (
	"fmt"
	"log"

	"github.com/Olexander753/WB_L0/internal/db"
	"github.com/Olexander753/WB_L0/internal/event"
	"github.com/Olexander753/WB_L0/pkg/util/config"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Read config")
	cfg := config.GetConfig()

	//connect to postgresql db
	url := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DB)
	repo, err := db.NewPostgres(url)
	if err != nil {
		log.Panicln(err)
		return
	}
	db.SetRepository(repo)
	defer db.Close()

	//connect to nats
	es, err := event.NewNats(fmt.Sprint("nats://%s", cfg.Nats.Address))
	if err != nil {
		log.Println(err)
	}
	event.SetEventStore(es)
	defer event.Close()

	//create http router

}

func newRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/meows", createModelHandler).
		Methods("POST").
		Queries("body", "{body}")
	return
}
