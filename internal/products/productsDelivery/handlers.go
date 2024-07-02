package productsDelivery

import (
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

		createCard := req.toCreateCard()
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

		createDeposit := req.toCreateDeposit()
		did, err := p.productUC.CreateNewDeposit(c.Context(), createDeposit)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"message": "Deposit created successfully",
			"cardID":  did,
		})
	}
}

func (p *ProductHandlers) GetCards() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := GetCardsRequest{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		cards, err := p.productUC.GetManyCards(c.Context(), req.UserID)
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
		req := GetDepositsRequest{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		deposits, err := p.productUC.GetManyDeposits(c.Context(), req.UserID)
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

		if err := p.productUC.DeleteCard(c.Context(), req.CardID); err != nil {
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

		if err := p.productUC.DeleteDeposit(c.Context(), req.DepositID); err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"data": "Deposit deleted successfully",
		})
	}
}
