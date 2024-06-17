package main

import (
	"bankApp1/config"
	"bankApp1/dbConnector"
	"bankApp1/runtime"
	"bankApp1/txManager"
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
	runtime.Run(mng)
	//ctx := context.Background()
	//cardRep := cardRepo.NewCardRepo(mng)
	//depRepo := depositRepo.NewDepositRepo(mng)
	//balRepo := balanceRepo.NewBalanceRepo(mng)
	//opRepo := operationRepo.NewOperationRepo(mng)
	//b := models.Balance{
	//	CardID:    models.CardID(10),
	//	Amount:    100000,
	//	CreatedAt: time.Now(),
	//}
	//if bid, err := balRepo.Create(ctx, b); err != nil {
	//	log.Fatalf("Error creating balance: %v", err)
	//} else {
	//	log.Printf("Created new balance: %v", bid)
	//}
	//sf := models.BalanceFilter{CardIDs: []models.CardID{14}}
	//rf := models.BalanceFilter{DepositIDs: []models.DepositID{5}}
	//pUC := paymentUsecase.NewPaymentUC(mng, balRepo, opRepo)
	//if oid, err := pUC.Send(sf, rf, 5000, "c2d"); err != nil {
	//	log.Fatalf("Error sending transfer: %v", err)
	//} else {
	//	log.Printf("Successfully sent transfer: %v", oid)
	//}
	//if _, err := pUC.PayIn(rf, 100000, "pay in"); err != nil {
	//	log.Fatalf("Error pay in: %v", err)
	//} else {
	//	log.Println("Pay in")
	//}

	//prUC := productsUsecase.NewProductsUsecase(mng, cardRep, depRepo, balRepo)
	//if id, err := prUC.CreateNewDeposit(4, "depchik", 1.25); err != nil {
	//	log.Fatalf("Error creating deposit: %v", err)
	//} else {
	//	log.Printf("Successfully create deposit %v", id)
	//}
	//usRep := userRepo.NewUserRepo(mng)

	//userUC := userUseCase.NewUserUC(mng, usRep)
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
	//f := models.UserFilter{Emails: []string{"bbb@qq.com"}}
	//if id, err := userUC.Login(f, "12345"); err != nil {
	//	log.Fatalf("Error logging in: %v", err)
	//} else {
	//	log.Printf("User %v logged in", id)
	//}

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
