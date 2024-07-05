package productsDelivery

import (
	"bankApp1/internal/models"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type ProductHandlers struct {
	productUC ProductsUC
}

func NewProductHandlers(productUC ProductsUC) ProductHandlers {
	return ProductHandlers{productUC: productUC}
}

func (p *ProductHandlers) CreateNewCard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := CreateCardRequest{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		if err := req.checkData(); err != nil {
			return err
		}

		claims, ok := c.Locals("claims").(*models.Claims)
		if !ok {
			return errors.New("cannot get claims")
		}

		createCard := req.toCreateCard()
		createCard.UserID = claims.UserID
		cid, err := p.productUC.CreateNewCard(c.Context(), createCard)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"message": "Card created successfully",
			"cardID":  cid,
		})
	}
}

func (p *ProductHandlers) CreateNewDeposit() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := CreateDepositRequest{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		if err := req.checkData(); err != nil {
			return err
		}

		claims, ok := c.Locals("claims").(*models.Claims)
		if !ok {
			return errors.New("cannot get claims")
		}

		createDeposit := req.toCreateDeposit()
		createDeposit.UserID = claims.UserID
		did, err := p.productUC.CreateNewDeposit(c.Context(), createDeposit)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"message":   "Deposit created successfully",
			"depositID": did,
		})
	}
}

func (p *ProductHandlers) GetCards() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("claims").(*models.Claims)
		if !ok {
			return errors.New("cannot get claims")
		}

		cards, err := p.productUC.GetManyCards(c.Context(), claims.UserID)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"data": fmt.Sprintf("%v", cards),
		})
	}
}

func (p *ProductHandlers) GetDeposits() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("claims").(*models.Claims)
		if !ok {
			return errors.New("cannot get claims")
		}

		deposits, err := p.productUC.GetManyDeposits(c.Context(), claims.UserID)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"data": fmt.Sprintf("%v", deposits),
		})
	}
}

func (p *ProductHandlers) DeleteCard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := DeleteCardRequest{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		claims, ok := c.Locals("claims").(*models.Claims)
		if !ok {
			return errors.New("cannot get claims")
		}

		if err := p.productUC.DeleteCard(c.Context(), req.CardID, claims.UserID); err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"data": "Card deleted successfully",
		})
	}
}

func (p *ProductHandlers) DeleteDeposit() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := DeleteDepositRequest{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		claims, ok := c.Locals("claims").(*models.Claims)
		if !ok {
			return errors.New("cannot get claims")
		}

		if err := p.productUC.DeleteDeposit(c.Context(), req.DepositID, claims.UserID); err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"data": "Deposit deleted successfully",
		})
	}
}
