package server

import (
	"bankApp1/internal/users/delivery/userHttp"
	userrepo "bankApp1/internal/users/repo"
	"bankApp1/internal/users/usecase"
)

func (s *Server) MapHandlers() {
	userRepo := userrepo.NewUserRepo(s.manager)
	userUC := usecase.NewUserUC(s.manager, userRepo)
	userHandlers := userHttp.NewUserHandlers(userUC)

	group := s.fiber.Group("users")
	userHttp.MapUserRoutes(group, userHandlers)
}
