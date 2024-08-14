package middlewares

import (
	"net/http"
	"task-manager/configs"
	"task-manager/internal/http/models"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var JWTMiddleware = echojwt.WithConfig(echojwt.Config{
	SigningKey: []byte(configs.Environment.JWTSecret),
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return new(models.JWTClaims)
	},
	ErrorHandler: func(c echo.Context, err error) error {
		c.Logger().Error(err)

		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Unauthorized, missing or invalid token",
		})
	},
})
