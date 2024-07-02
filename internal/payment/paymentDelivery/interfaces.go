package paymentDelivery

import (
	"bankApp1/internal/models"
	"bankApp1/internal/payment/paymentUsecase"
	"context"
	"github.com/gofiber/fiber/v2"
)

type (
	PaymentUC interface {
		Send(context.Context, *paymentUsecase.SendData) (models.OperationID, error)
		PayIn(context.Context, *paymentUsecase.PayData) (models.OperationID, error)
		PayOut(context.Context, *paymentUsecase.PayData) (models.OperationID, error)
	}

	Handlers interface {
		Send() fiber.Handler
		PayIn() fiber.Handler
		PayOut() fiber.Handler
	}
)
