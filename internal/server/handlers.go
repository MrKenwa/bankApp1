package server

import (
	"bankApp1/internal/balances/repo/postgres"
	balanceusecase "bankApp1/internal/balances/usecase"
	"bankApp1/internal/payment/delivery"
	paymentrepo "bankApp1/internal/payment/repo/postgres"
	paymentusecase "bankApp1/internal/payment/usecase"
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
}
