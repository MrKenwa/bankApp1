package userHttp

import (
	"bankApp1/config"
	"bankApp1/internal/models"
	"errors"
	"github.com/gofiber/fiber/v2"
	"time"
)

type UserHandlers struct {
	config *config.Config
	userUC UserUC
}

func NewUserHandlers(config *config.Config, userUC UserUC) UserHandlers {
	return UserHandlers{config: config, userUC: userUC}
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

		sessionKey, err := h.userUC.Login(c.Context(), logUser)
		if err != nil {
			return err
		}

		c.Cookie(&fiber.Cookie{
			Name:     "session_key",
			Value:    sessionKey,
			Path:     "/",
			Secure:   true,
			HTTPOnly: true,
			Domain:   h.config.Server.Domain,
			Expires:  time.Now().Add(time.Second * time.Duration(h.config.SessionSettings.SessionTTLSeconds)),
		})

		return c.JSON(fiber.Map{
			"data": "User was logged in",
		})
	}
}

func (h *UserHandlers) GetOwn() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("claims").(models.Claims)
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

func (h *UserHandlers) Hello() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": "hello! I'm alive!",
		})
	}
}
