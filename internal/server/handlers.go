package server

import (
	"bankApp1/internal/users/delivery/userHttp"
	userrepo "bankApp1/internal/users/repo"
	"bankApp1/internal/users/usecase"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
)

func (s *Server) MapHandlers() {
	userRepo := userrepo.NewUserRepo(trmsqlx.DefaultCtxGetter, &s.postgres)
	userUC := usecase.NewUserUC(s.manager, userRepo)
	userHandlers := userHttp.NewUserHandlers(userUC)

	group := s.fiber.Group("users")
	userHttp.MapUserRoutes(group, userHandlers)
}
