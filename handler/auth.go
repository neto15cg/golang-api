package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// AuthData is struct to login request data
type AuthData struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validade:"required"`
}

// AuthHandler service handler authentication
type AuthHandler struct {
	signin func(email, password string) (string, error)
}

func (handler *AuthHandler) login(c echo.Context) error {
	req := AuthData{}

	err := c.Bind(&req)
	if err != nil {
		return echo.ErrUnauthorized
	}

	r, err := handler.signin(req.Username, req.Password)
	if err != nil {
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": r,
	})
}
