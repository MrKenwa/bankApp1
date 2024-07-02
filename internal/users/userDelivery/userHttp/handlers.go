package userHttp

import (
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

		regUser := req.toRegisterUser()

		uid, err := h.userUC.Register(c.Context(), regUser)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"message": "User was registered",
			"userID":  uid,
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

		uid, err := h.userUC.Login(c.Context(), logUser)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"message": "User was logged in",
			"userID":  uid,
		})
	}
}
