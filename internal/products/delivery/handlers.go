package delivery

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type ProductHandlers struct {
	productUC ProductUC
}

func NewProductHandlers(productUC ProductUC) *ProductHandlers {
	return &ProductHandlers{productUC: productUC}
}

func (p *ProductHandlers) CreateNewCard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := CreateCardRequest{}
		c.Body()
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		createCard := req.toCreateCard()
		cid, err := p.productUC.CreateNewCard(c.Context(), createCard)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"data": fmt.Sprintf("Card created with id: %d", cid),
		})
	}
}

func (p *ProductHandlers) DeleteCard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := DeleteCardRequest{}
		c.Body()
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		if err := p.productUC.DeleteCard(c.Context(), req.CardID); err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"data": "Card deleted successfully",
		})
	}
}

func (p *ProductHandlers) CreateNewDeposit() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := CreateDepositRequest{}
		c.Body()
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		createDeposit := req.toCreateDeposit()
		did, err := p.productUC.CreateNewDeposit(c.Context(), createDeposit)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"data": fmt.Sprintf("Deposit created with id: %d", did),
		})
	}
}

func (p *ProductHandlers) DeleteDeposit() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := DeleteDepositRequest{}
		c.Body()
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		if err := p.productUC.DeleteDeposit(c.Context(), req.DepositID); err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"data": "Deposit deleted successfully",
		})
	}
}

func (p *ProductHandlers) GetCards() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := GetCardsRequest{}
		c.Body()
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		cards, err := p.productUC.GetCards(c.Context(), req.UserID)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"data": fmt.Sprintf("Card list: %v", cards),
		})
	}
}

func (p *ProductHandlers) GetDeposits() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := GetDepositsRequest{}
		c.Body()
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		deposits, err := p.productUC.GetDeposits(c.Context(), req.UserID)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"data": fmt.Sprintf("Deposit list: %v", deposits),
		})
	}
}
