package handlers

import (
	"net/http"
	"task-manager/internal/domain"

	"github.com/labstack/echo/v4"
)

func HandleError(err error, c echo.Context) {
	c.Logger().Error(err)

	if domainErr, ok := err.(*domain.DomainError); ok {
		c.String(domainErr.Code, domainErr.Message)
		return
	}

	c.String(http.StatusInternalServerError, err.Error())
}

func HandleJWTError(c echo.Context, err error) error {
	c.Logger().Error(err)

	return c.JSON(http.StatusUnauthorized, map[string]string{
		"message": "Unauthorized, missing or invalid token",
	})
}
