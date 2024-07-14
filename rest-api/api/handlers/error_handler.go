package handlers

import (
	"net/http"
	"rest-api/internal/domain"

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
