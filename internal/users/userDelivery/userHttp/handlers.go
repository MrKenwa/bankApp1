package userHttp

import (
	"bankApp1/internal/models"
	"errors"
	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	userUC UserUC
}

func NewUserHandlers(userUC UserUC) UserHandlers {
	return UserHandlers{userUC}
}

func (h *UserHandlers) Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := RegisterRequest{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		if !req.checkData() {
			return errors.New("invalid data")
		}

		regUser := req.toRegisterUser()

		uid, err := h.userUC.Register(c.Context(), regUser)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"message": "User was registered",
			"data":    uid,
		})
	}
}

func (h *UserHandlers) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := LoginRequest{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		logUser := req.toLoginUser()

		token, err := h.userUC.Login(c.Context(), logUser)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"message": "User was logged in",
			"data":    token,
		})
	}
}

func (h *UserHandlers) GetOwn() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("claims").(*models.Claims)
		if !ok {
			return errors.New("cannot get claims")
		}

		user, err := h.userUC.GetUser(c.Context(), claims.UserID)
		if err != nil {
			return err
		}

		userResp := getUserResponse{
			UserID:     user.UserID,
			Name:       user.Name,
			Lastname:   user.Lastname,
			Patronymic: user.Patronymic,
			Email:      user.Email,
		}

		return c.JSON(fiber.Map{
			"data": userResp,
		})
	}
}

func (h *UserHandlers) RefreshToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("claims").(*models.Claims)
		if !ok {
			return errors.New("cannot get claims")
		}

		token, err := h.userUC.RefreshToken(c.Context(), claims.UserID)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"data": token,
		})
	}
}
