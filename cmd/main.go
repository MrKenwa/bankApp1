package main

import (
	"bankApp1/config"
	"bankApp1/internal/server"
	"bankApp1/pkg/dbConnector/postgres"
	"bankApp1/pkg/dbConnector/redis"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	} else {
		log.Println("Loaded config")
	}

	sqlDB, err := postgres.ConnectToDB(cfg)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	} else {
		log.Println("Connected to DB")
	}
	db := postgres.PostgresDB{DB: sqlDB}

	redisClient := redis.NewRedisClient(cfg)
	log.Println("Connected to Redis")

	mngr := manager.Must(trmsqlx.NewDefaultFactory(sqlDB))

	s := server.NewServer(cfg, db, redisClient, mngr)

	if err := s.Run(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
