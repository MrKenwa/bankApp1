package paymentDelivery

import (
	"github.com/gofiber/fiber/v2"
)

type PaymentHandlers struct {
	paymentUC PaymentUC
}

func NewPaymentHandlers(paymentUC PaymentUC) *PaymentHandlers {
	return &PaymentHandlers{paymentUC: paymentUC}
}

func (p *PaymentHandlers) Send() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := SendRequest{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		err := req.checkData()
		if err != nil {
			return err
		}

		sendData := req.toSendData()
		opid, err := p.paymentUC.Send(c.Context(), sendData)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"message":     "Money sent successfully",
			"operationID": opid,
		})
	}
}

func (p *PaymentHandlers) PayIn() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := PayRequest{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		err := req.checkData()
		if err != nil {
			return err
		}

		payData := req.toPayData()
		opid, err := p.paymentUC.PayIn(c.Context(), payData)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"message":     "Money sent successfully",
			"operationID": opid,
		})
	}
}

func (p *PaymentHandlers) PayOut() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := PayRequest{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		err := req.checkData()
		if err != nil {
			return err
		}

		payData := req.toPayData()
		opid, err := p.paymentUC.PayOut(c.Context(), payData)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"message":     "Money sent successfully",
			"operationID": opid,
		})
	}
}
