package main

import (
	"bankApp1/config"
	"bankApp1/internal/server"
	"bankApp1/pkg/dbConnector"
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

	sqlDB, err := dbConnector.ConnectToDB(cfg)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	} else {
		log.Println("Connected to DB")
	}
	db := dbConnector.PostgresDB{DB: sqlDB}
	mngr := manager.Must(trmsqlx.NewDefaultFactory(sqlDB))

	s := server.NewServer(cfg, db, mngr)

	if err := s.Run(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
