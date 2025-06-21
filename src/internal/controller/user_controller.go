package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imraushankr/gozen/src/internal/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) GetProfile(c *fiber.Ctx) error {
	return nil
}

func (uc *UserController) UpdateProfile(c *fiber.Ctx) error {
	return nil
}

func (uc *UserController) GetUsers(c *fiber.Ctx) error {
	return nil
}

func (uc *UserController) DeleteAccount(c *fiber.Ctx) error {
	return nil
}