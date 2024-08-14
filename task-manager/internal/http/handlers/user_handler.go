package handlers

import (
	"net/http"
	"task-manager/configs"
	"task-manager/internal/domain/user"
	"task-manager/internal/http/middlewares"
	"task-manager/internal/http/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service user.UserService
}

func NewUserHandler(service user.UserService) *UserHandler {
	return &UserHandler{service}
}

func (handler *UserHandler) RegisterRoutes(group *echo.Group, routePrefix string) {
	group.POST(routePrefix+"/login", handler.loginHandler)
	group.POST(routePrefix+"/register", handler.registerHandler)
	group.GET(routePrefix, handler.getUserHandler, middlewares.JWTMiddleware)
}

// @Router /api/v1/users/login [post]
// @Summary Login with existing user
// @Description Returns token for valid authentication
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param email formData string true "User email"
// @Param password formData string true "User password" format(password)
// @Success 201 {object} map[string]interface{} "access_token"
// @Failure 400
// @Failure 404
// @Failure 500
func (handler *UserHandler) loginHandler(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if len(email) == 0 || len(password) == 0 {
		return c.JSON(http.StatusBadRequest, nil)
	}

	registeredUser, err := handler.service.GetUser(email)
	if err != nil {
		return err
	}

	if email != registeredUser.Email || password != registeredUser.Password {
		return c.JSON(http.StatusUnauthorized, nil)
	}

	claims := &models.JWTClaims{
		Email: registeredUser.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        registeredUser.ID.String(),
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
// @Summary Register new user
// @Description Registers new user
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param email formData string true "User email"
// @Param password formData string true "User password" format(password)
// @Success 201 {object} user.User
// @Failure 400
// @Failure 500
func (handler *UserHandler) registerHandler(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user, err := handler.service.CreateUser(email, password)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"user": user,
	})
}

// @Router /api/v1/users [get]
// @Summary Get user info
// @Description Returns authenticated user
// @Security BearerAuth
// @Produce json
// @Success 200 {object} user.User
// @Failure 500
func (handler *UserHandler) getUserHandler(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*models.JWTClaims)

	user, err := handler.service.GetUser(claims.Email)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
