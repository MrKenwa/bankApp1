package delivery

import (
	"bankApp1/internal/models"
	"bankApp1/internal/payment/usecase"
	"context"
	"github.com/gofiber/fiber/v2"
)

type (
	PaymentUC interface {
		Send(context.Context, *usecase.SendData) (models.OperationID, error)
		PayIn(context.Context, *usecase.PayData) (models.OperationID, error)
		PayOut(context.Context, *usecase.PayData) (models.OperationID, error)
	}

	Handlers interface {
		Send() fiber.Handler
		PayIn() fiber.Handler
		PayOut() fiber.Handler
	}
)
