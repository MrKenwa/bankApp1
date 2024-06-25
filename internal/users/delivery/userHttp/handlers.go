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

		uid, err := h.userUC.Register(regUser)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"data": fmt.Sprintf("User was registered with id=%d", uid),
		})
	}
}
