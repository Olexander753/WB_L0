package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Olexander753/WB_L0/internal/cach"
	"github.com/Olexander753/WB_L0/internal/config"
	"github.com/Olexander753/WB_L0/internal/event"
	"github.com/Olexander753/WB_L0/internal/server"
	"github.com/Olexander753/WB_L0/pkg/handler"
	"github.com/Olexander753/WB_L0/pkg/repository"
	"github.com/Olexander753/WB_L0/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/tinrab/retry"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	cfg := config.GetConfig()

	//connect to postgresql db
	db, err := repository.NewPostgres(cfg)
	if err != nil {
		log.Println("Failed connect to postgres, error: ", err)
	}

	//connect to nats
	retry.ForeverSleep(2*time.Second, func(_ int) error {
		es, err := event.NewNats(fmt.Sprintf("nats://%s", cfg.Nats.Address))
		if err != nil {
			log.Println("Failed connect to nats, error: ", err)
			return err
		}
		err = es.OnModelCreated(event.ModelCreated)
		if err != nil {
			log.Println(err)
			return err
		}
		event.SetEventStore(es)
		return nil
	})
	defer event.Close()

	//create cach
	ce := cach.NewCach(db)
	ce.SelectModels()

	repo := repository.NewRepository(db, ce)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	//create http router
	srv := new(server.Server)

	err = srv.Run(cfg.Server.Port, handlers.InutRoutes())
	if err != nil {
		log.Fatal("error http server: ", err)
	}

}
