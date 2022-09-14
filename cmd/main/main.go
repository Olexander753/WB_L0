package main

import (
	"fmt"
	"log"

	"github.com/Olexander753/WB_L0/internal/config"
	"github.com/Olexander753/WB_L0/internal/event"
	"github.com/Olexander753/WB_L0/internal/server"
	"github.com/Olexander753/WB_L0/pkg/handler"
	"github.com/Olexander753/WB_L0/pkg/repository"
	"github.com/Olexander753/WB_L0/pkg/service"
)

func main() {

	cfg := config.GetConfig()

	//connect to postgresql db
	db, err := repository.NewPostgres(cfg)
	if err != nil {
		log.Println("Failed connect to postgres, error: ", err)
	}

	//connect to nats
	es, err := event.NewNats(fmt.Sprintf("nats://%s", cfg.Nats.Address))
	if err != nil {
		log.Println("Failed connect to nats, error: ", err)
	}
	event.SetEventStore(es)
	defer event.Close()

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	//create http router
	srv := new(server.Server)

	err = srv.Run(cfg.Server.Port, handlers.InutRoutes())
	if err != nil {
		log.Fatal("error http server: ", err)
	}
}
