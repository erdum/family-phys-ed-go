package controllers

import (
	"go-api/services/auth"
	"go-api/requests"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authSevice auth.AuthService
}

func NewAuthController(authSevice auth.AuthService) *AuthController {
	return &AuthController{authSevice: authSevice}
}

func (ac *AuthController) Login(context echo.Context) error {
	payload := context.Get("valid_payload").(*requests.LoginRequest)

	data, err := ac.authSevice.AuthenticateWithThirdParty(payload.IdToken)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, data)
}
