package server

import (
	"bankApp1/internal/balances/balancesRepo/postgres"
	"bankApp1/internal/balances/balancesUsecase"
	"bankApp1/internal/cards/cardsRepo"
	"bankApp1/internal/cards/cardsUsecase"
	"bankApp1/internal/deposits/depositsRepo"
	"bankApp1/internal/deposits/depositsUsecase"
	"bankApp1/internal/payment/paymentDelivery"
	operationsRepo "bankApp1/internal/payment/paymentRepo/postgres"
	"bankApp1/internal/payment/paymentUsecase"
	"bankApp1/internal/products/productsDelivery"
	"bankApp1/internal/products/productsUsecase"
	"bankApp1/internal/users/userDelivery/userHttp"
	userrepo "bankApp1/internal/users/userRepo/postgres"
	"bankApp1/internal/users/userUsecase"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
)

func (s *Server) MapHandlers() {
	userRepo := userrepo.NewUserRepo(trmsqlx.DefaultCtxGetter, &s.postgres)
	userUC := userUsecase.NewUserUC(s.manager, &userRepo)
	userHandlers := userHttp.NewUserHandlers(&userUC)

	userGroup := s.fiber.Group("users")
	userHttp.MapUserRoutes(userGroup, &userHandlers)

	balanceRepo := postgres.NewBalanceRepo(trmsqlx.DefaultCtxGetter, &s.postgres)
	operationRepo := operationsRepo.NewOperationRepo(trmsqlx.DefaultCtxGetter, &s.postgres)
	balanceUC := balancesUsecase.NewBalanceUC(s.manager, &balanceRepo)
	paymentUC := paymentUsecase.NewPaymentUC(s.manager, &balanceUC, &operationRepo)
	paymentHandlers := paymentDelivery.NewPaymentHandlers(paymentUC)

	paymentGroup := s.fiber.Group("payment")
	paymentDelivery.MapPaymentRoutes(paymentGroup, paymentHandlers)

	cardRepo := cardsRepo.NewCardRepo(trmsqlx.DefaultCtxGetter, &s.postgres)
	depositRepo := depositsRepo.NewDepositRepo(trmsqlx.DefaultCtxGetter, &s.postgres)
	cardsUC := cardsUsecase.NewCardUC(s.manager, &cardRepo)
	depositsUC := depositsUsecase.NewDepositsUC(s.manager, depositRepo)
	productUC := productsUsecase.NewProductsUC(s.manager, &cardsUC, &depositsUC, &balanceUC)
	productHandlers := productsDelivery.NewProductHandlers(&productUC)

	productGroup := s.fiber.Group("product")
	productsDelivery.MapProductsRoutes(productGroup, &productHandlers)
}
