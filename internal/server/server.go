package server

import (
	"bankApp1/config"
	"bankApp1/pkg/dbConnector"
	"bankApp1/txManager"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	fiber    *fiber.App
	cfg      *config.Config
	postgres dbConnector.PostgresDB
	manager  *txManager.TxManager
}

func NewServer(cfg *config.Config, postgres dbConnector.PostgresDB, manager *txManager.TxManager) *Server {
	return &Server{
		fiber:    fiber.New(),
		cfg:      cfg,
		postgres: postgres,
		manager:  manager,
	}
}

func (s *Server) Run() error {
	s.MapHandlers()

	go func() {
		if err := s.fiber.Listen(s.cfg.Server.Host); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	err := s.fiber.Shutdown()
	if err != nil {
		fmt.Println("HTTP server shutdown with panic: ", err)
		panic(err)
	} else {
		fmt.Println("HTTP server closed properly")
	}

	return nil
}
