package main

import (
	"device_service/api"
	"device_service/config"
	"device_service/mq"
	"device_service/repository"
	"device_service/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	db, err := sqlx.Connect("postgres", cfg.GetSQLxDSN())
	if err != nil {
		log.Fatal(err)
	}

	conn, err := amqp.Dial(cfg.GetAMqpDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	repo := repository.NewDeviceRepository(db)
	svc := service.NewDeviceService(repo)

	producer, err := mq.NewProducer(conn)
	if err != nil {
		log.Fatal(err)
	}

	handler := api.NewDeviceHandler(svc, producer)

	r := gin.Default()
	handler.RegisterRoutes(r)

	log.Println("Device service started on :8090")
	r.Run(cfg.GetServerDSN())
}
