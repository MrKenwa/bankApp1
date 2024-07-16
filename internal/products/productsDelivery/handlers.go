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

func (h *ProductHandlers) CreateNewCard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("claims").(models.Claims)
		if !ok {
			return errors.New("cannot get claims")
		}

		req := CreateCardRequest{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		if err := req.checkData(); err != nil {
			return err
		}

		createCard := req.toCreateCard()
		createCard.UserID = claims.UserID
		cid, err := h.productUC.CreateNewCard(c.Context(), createCard)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"message": "Card created successfully",
			"cardID":  cid,
		})
	}
}

func (h *ProductHandlers) CreateNewDeposit() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("claims").(models.Claims)
		if !ok {
			return errors.New("cannot get claims")
		}

		req := CreateDepositRequest{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		if err := req.checkData(); err != nil {
			return err
		}

		createDeposit := req.toCreateDeposit()
		createDeposit.UserID = claims.UserID
		did, err := h.productUC.CreateNewDeposit(c.Context(), createDeposit)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"message": "Deposit created successfully",
			"cardID":  did,
		})
	}
}

func (h *ProductHandlers) GetCards() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("claims").(models.Claims)
		if !ok {
			return errors.New("cannot get claims")
		}

		cards, err := h.productUC.GetManyCards(c.Context(), claims.UserID)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"data": fmt.Sprintf("%v", cards),
		})
	}
}

func (h *ProductHandlers) GetDeposits() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("claims").(models.Claims)
		if !ok {
			return errors.New("cannot get claims")
		}

		deposits, err := h.productUC.GetManyDeposits(c.Context(), claims.UserID)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"data": fmt.Sprintf("%v", deposits),
		})
	}
}

func (h *ProductHandlers) DeleteCard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("claims").(models.Claims)
		if !ok {
			return errors.New("cannot get claims")
		}

		req := DeleteCardRequest{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		if err := h.productUC.DeleteCard(c.Context(), req.CardID, claims.UserID); err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"data": "Card deleted successfully",
		})
	}
}

func (h *ProductHandlers) DeleteDeposit() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("claims").(models.Claims)
		if !ok {
			return errors.New("cannot get claims")
		}

		req := DeleteDepositRequest{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		if err := h.productUC.DeleteDeposit(c.Context(), req.DepositID, claims.UserID); err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"data": "Deposit deleted successfully",
		})
	}
}
