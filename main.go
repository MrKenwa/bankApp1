package main

import (
	"bankApp1/config"
	"bankApp1/dbConnector"
	"bankApp1/models"
	"bankApp1/repo/userRepo"
	"bankApp1/txManager"
	userUseCase "bankApp1/usecase"
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
	usRep := userRepo.NewUserRepo(mng)

	userUC := userUseCase.NewUserUC(mng, usRep)
	//u := &models.User{
	//	Name:           "Aaaa",
	//	Lastname:       "Bbbbb",
	//	Patronymic:     "SHuuuuuu",
	//	Email:          "bbb@qq.com",
	//	Password:       "12345",
	//	PassportNumber: "666",
	//	CreatedAt:      time.Now(),
	//}
	//if id, err := userUC.Register(u); err != nil {
	//	log.Fatalf("Error registering user: %v", err)
	//} else {
	//	log.Printf("User ID: %d", id)
	//}
	f := models.UserFilter{Emails: []string{"bbb@qq.com"}}
	if id, err := userUC.Login(f, "12345"); err != nil {
		log.Fatalf("Error logging in: %v", err)
	} else {
		log.Printf("User %v logged in", id)
	}

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
