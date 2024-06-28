package userHttp

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	userUC UserUC
}

func NewUserHandlers(userUC UserUC) *UserHandlers {
	return &UserHandlers{userUC}
}

func (h *UserHandlers) Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := RegisterRequest{}
		c.Body()
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		regUser := req.toRegisterUser()

		uid, err := h.userUC.Register(c.Context(), regUser)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"data": fmt.Sprintf("User was registered with id=%d", uid),
		})
	}
}

func (h *UserHandlers) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := LoginRequest{}
		c.Body()
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		logUser := req.toLoginUser()

		uid, err := h.userUC.Login(c.Context(), logUser)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"data": fmt.Sprintf("User was logged in with id=%d", uid),
		})
	}
}
