package server

import (
	"bankApp1/internal/balances/balancesRepo/postgres"
	"bankApp1/internal/balances/balancesUsecase"
	"bankApp1/internal/cards/cardsRepo"
	"bankApp1/internal/cards/cardsUsecase"
	"bankApp1/internal/deposits/depositsRepo"
	"bankApp1/internal/deposits/depositsUsecase"
	"bankApp1/internal/middleware"
	"bankApp1/internal/payment/paymentDelivery"
	operationsRepo "bankApp1/internal/payment/paymentRepo/postgres"
	"bankApp1/internal/payment/paymentUsecase"
	"bankApp1/internal/products/productsDelivery"
	"bankApp1/internal/products/productsUsecase"
	"bankApp1/internal/users/userDelivery/userHttp"
	userrepo "bankApp1/internal/users/userRepo/postgres"
	"bankApp1/internal/users/userRepo/redis"
	"bankApp1/internal/users/userUsecase"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
)

func (s *Server) MapHandlers() {
	userRepo := userrepo.NewUserRepo(trmsqlx.DefaultCtxGetter, &s.postgres)
	userRedisRepo := redis.NewUserRedisRepo(s.cfg, s.redis)
	userUC := userUsecase.NewUserUC(s.manager, &userRepo, userRedisRepo)
	userHandlers := userHttp.NewUserHandlers(s.cfg, &userUC)
	cardRepo := cardsRepo.NewCardRepo(trmsqlx.DefaultCtxGetter, &s.postgres)
	depositRepo := depositsRepo.NewDepositRepo(trmsqlx.DefaultCtxGetter, &s.postgres)
	cardsUC := cardsUsecase.NewCardUC(s.manager, &cardRepo)
	depositsUC := depositsUsecase.NewDepositsUC(s.manager, depositRepo)

	mw := middleware.NewMDWManager(userRedisRepo)
	userGroup := s.fiber.Group("users")
	userHttp.MapUserRoutes(userGroup, &userHandlers, mw)

	balanceRepo := postgres.NewBalanceRepo(trmsqlx.DefaultCtxGetter, &s.postgres)
	operationRepo := operationsRepo.NewOperationRepo(trmsqlx.DefaultCtxGetter, &s.postgres)
	balanceUC := balancesUsecase.NewBalanceUC(s.manager, &balanceRepo)
	paymentUC := paymentUsecase.NewPaymentUC(s.manager, &balanceUC, &operationRepo, &cardsUC, &depositsUC)
	paymentHandlers := paymentDelivery.NewPaymentHandlers(paymentUC)

	paymentGroup := s.fiber.Group("payment")
	paymentDelivery.MapPaymentRoutes(paymentGroup, paymentHandlers, mw)

	productUC := productsUsecase.NewProductsUC(s.manager, &cardsUC, &depositsUC, &balanceUC)
	productHandlers := productsDelivery.NewProductHandlers(&productUC)

	productGroup := s.fiber.Group("product")
	productsDelivery.MapProductsRoutes(productGroup, &productHandlers, mw)
}
