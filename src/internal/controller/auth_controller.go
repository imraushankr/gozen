package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imraushankr/gozen/src/internal/service"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) Register(c *fiber.Ctx) error {
	return nil
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	return nil
}

func (ac *AuthController) RefreshToken(c *fiber.Ctx) error {
	return nil
}

func (ac *AuthController) ForgotPassword(c *fiber.Ctx) error {
	return nil
}

func (ac *AuthController) ResetPassword(c *fiber.Ctx) error {
	return nil
}

func (ac *AuthController) ChangePassword(c *fiber.Ctx) error {
	return nil
}