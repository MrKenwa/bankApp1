package main

import (
	"bankApp1/config"
	"bankApp1/dbConnector"
	"bankApp1/txManager"
	"context"
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
	mng := txManager.NewTxManager(db)
	ctx := context.Background()

	// ТЕСТЫ

	//userRep := userRepo.NewUserRepo(mng)
	//user := &models.User{
	//	Name:           "Vasya",
	//	Lastname:       "Pupkin",
	//	Patronymic:     "Vtoroi",
	//	Email:          "www@www.com",
	//	Password:       "1111",
	//	PassportNumber: "1111",
	//	CreatedAt:      time.Time{},
	//	DeletedAt:      nil,
	//}
	//if id, err := userRep.Create(ctx, user); err != nil {
	//	log.Fatalf("Error creating user: %v", err)
	//} else {
	//	fmt.Println(id)
	//}

	//Get User
	//filter := models.UserFilter{IDs: []models.UserID{3}}
	//if user, err := userRep.Get(ctx, filter); err != nil {
	//	log.Fatalf("Error getting users: %v", err)
	//} else {
	//	log.Printf("Found user: %v", user)
	//}
}
