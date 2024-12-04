package server

import (
	"bankApp1/config"
	"bankApp1/pkg/dbConnector/postgres"
	"fmt"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	fiber    *fiber.App
	cfg      *config.Config
	postgres postgres.PostgresDB
	redis    *redis.Client
	manager  *manager.Manager
}

func NewServer(cfg *config.Config, postgres postgres.PostgresDB, redis *redis.Client, manager *manager.Manager) *Server {
	return &Server{
		fiber:    fiber.New(),
		cfg:      cfg,
		postgres: postgres,
		redis:    redis,
		manager:  manager,
	}
}

func (s *Server) Run() error {
	s.MapHandlers()

	go func() {
		s.fiber.Post("/hello", func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"data": "hello! I'm alive!",
			})
		})
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
