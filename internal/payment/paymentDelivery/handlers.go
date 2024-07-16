package paymentDelivery

import (
	"bankApp1/internal/models"
	"errors"
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
		claims, ok := c.Locals("claims").(models.Claims)
		if !ok {
			return errors.New("cannot get claims")
		}

		req := SendRequest{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		err := req.checkData()
		if err != nil {
			return err
		}

		sendData := req.toSendData()
		sendData.UserID = claims.UserID
		opid, err := p.paymentUC.Send(c.Context(), sendData)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"message": "Money sent successfully",
			"data":    opid,
		})
	}
}

func (p *PaymentHandlers) PayIn() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("claims").(models.Claims)
		if !ok {
			return errors.New("cannot get claims")
		}

		req := PayRequest{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		err := req.checkData()
		if err != nil {
			return err
		}

		payData := req.toPayData()
		payData.UserID = claims.UserID
		payData.OpType = "pay in"
		opid, err := p.paymentUC.PayIn(c.Context(), payData)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"message": "Money sent successfully",
			"data":    opid,
		})
	}
}

func (p *PaymentHandlers) PayOut() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("claims").(models.Claims)
		if !ok {
			return errors.New("cannot get claims")
		}

		req := PayRequest{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		err := req.checkData()
		if err != nil {
			return err
		}

		payData := req.toPayData()
		payData.UserID = claims.UserID
		payData.OpType = "pay out"
		opid, err := p.paymentUC.PayOut(c.Context(), payData)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"message": "Money sent successfully",
			"data":    opid,
		})
	}
}
