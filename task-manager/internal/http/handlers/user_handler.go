package handlers

import (
	"net/http"
	"task-manager/configs"
	"task-manager/internal/domain/user"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type jwtClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type UserHandler struct {
	service user.UserService
}

func NewUserHandler(service user.UserService) *UserHandler {
	return &UserHandler{service}
}

func (handler *UserHandler) RegisterRoutes(group *echo.Group, routePrefix string) {
	group.POST(routePrefix+"/login", handler.loginHandler)
	group.POST(routePrefix+"/register", handler.registerHandler)
	group.GET(routePrefix, handler.getUserHandler)
}

// @Router /api/v1/users/login [post]
// @Summary Login user
func (handler *UserHandler) loginHandler(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user, err := handler.service.GetUser(email)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	if email != user.Email || password != user.Password {
		return echo.ErrUnauthorized
	}

	claims := &jwtClaims{
		user.Email,
		jwt.RegisteredClaims{
			ID:        user.ID.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(configs.Environment.JWTSecret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"access_token": signedToken,
	})
}

// @Router /api/v1/users/register [post]
// @Summary Registers user
func (handler *UserHandler) registerHandler(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user, err := handler.service.CreateUser(email, password)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user": user,
	})
}

// @Router /api/v1/users [get]
// @Summary Login user
func (handler *UserHandler) getUserHandler(c echo.Context) error {
	email := c.FormValue("email")
	user, err := handler.service.GetUser(email)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
