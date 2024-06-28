package server

import (
	"bankApp1/internal/balances/repo/postgres"
	balanceusecase "bankApp1/internal/balances/usecase"
	"bankApp1/internal/payment/delivery"
	paymentrepo "bankApp1/internal/payment/repo/postgres"
	paymentusecase "bankApp1/internal/payment/usecase"
	productsDelivery "bankApp1/internal/products/delivery"
	productsrepo "bankApp1/internal/products/repo/postgres"
	productsusecase "bankApp1/internal/products/usecase"
	"bankApp1/internal/users/delivery/userHttp"
	userrepo "bankApp1/internal/users/repo/postgres"
	"bankApp1/internal/users/usecase"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
)

func (s *Server) MapHandlers() {
	userRepo := userrepo.NewUserRepo(trmsqlx.DefaultCtxGetter, &s.postgres)
	userUC := usecase.NewUserUC(s.manager, userRepo)
	userHandlers := userHttp.NewUserHandlers(userUC)

	userGroup := s.fiber.Group("users")
	userHttp.MapUserRoutes(userGroup, userHandlers)

	balanceRepo := postgres.NewBalanceRepo(trmsqlx.DefaultCtxGetter, &s.postgres)
	operationRepo := paymentrepo.NewOperationRepo(trmsqlx.DefaultCtxGetter, &s.postgres)
	balanceUC := balanceusecase.NewBalanceUsecase(s.manager, balanceRepo)
	paymentUC := paymentusecase.NewPaymentUC(s.manager, balanceUC, operationRepo)
	paymentHandlers := delivery.NewPaymentHandlers(paymentUC)

	paymentGroup := s.fiber.Group("payment")
	delivery.MapPaymentRoutes(paymentGroup, paymentHandlers)

	cardsRepo := productsrepo.NewCardRepo(trmsqlx.DefaultCtxGetter, &s.postgres)
	depositsRepo := productsrepo.NewDepositRepo(trmsqlx.DefaultCtxGetter, &s.postgres)
	productUC := productsusecase.NewProductsUC(s.manager, cardsRepo, depositsRepo, balanceUC)
	productHandlers := productsDelivery.NewProductHandlers(productUC)

	productGroup := s.fiber.Group("product")
	productsDelivery.MapProductsRoutes(productGroup, productHandlers)
}
