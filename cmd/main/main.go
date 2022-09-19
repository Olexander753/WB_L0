package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Olexander753/WB_L0/internal/config"
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
		log.Fatal("Failed connect to postgres, error: ", err)
	}

	//create cach
	ce := repository.NewCach(db)
	ce.SelectModels()

	repo := repository.NewRepository(db, ce)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	//connect to nats
	retry.ForeverSleep(2*time.Second, func(_ int) error {
		es, err := service.NewNats(fmt.Sprintf("nats://%s", cfg.Nats.Address), repo)
		if err != nil {
			log.Println("Failed connect to nats, error: ", err)
			return err
		}
		err = es.CreateModel()
		if err != nil {
			log.Println(err)
			return err
		}
		services.SetEventStore(es)
		return nil
	})
	defer services.Close()

	//create http router
	srv := new(server.Server)

	err = srv.Run(cfg.Server.Port, handlers.InutRoutes())
	if err != nil {
		log.Fatal("error http server: ", err)
	}

}
