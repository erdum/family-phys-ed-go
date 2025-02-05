package controllers

import (
	"go-api/services/auth"
	"go-api/requests"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthController struct {
	authService auth.AuthService
	tokenService auth.TokenService
	db *gorm.DB
}

func New(
	authService auth.AuthService,
	tokenService auth.TokenService,
	db *gorm.DB,
) *AuthController {
	return &AuthController{
		authService: authService,
		tokenService: tokenService,
		db: db,
	}
}

func (ac *AuthController) Register(c echo.Context) error {
	payload := c.Get("valid_payload").(*requests.RegisterRequest)

	response, err := ac.authService.Register(payload)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (ac *AuthController) Login(c echo.Context) error {
	payload := c.Get("valid_payload").(*requests.LoginRequest)

	response, err := ac.authService.Login(payload)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (ac *AuthController) SignOn(c echo.Context) error {
	payload := c.Get("valid_payload").(*requests.SignOnRequest)

	response, err := ac.authService.SignOn(payload)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}
