package main

import (
	"bankApp1/config"
	"bankApp1/dbConnector"
	"fmt"
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
	db := &dbConnector.DataBase{DB: sqlDB}
	fmt.Printf("%#v\n", db)
}
