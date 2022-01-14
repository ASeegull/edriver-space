package handler

import (
	"errors"
	"net/http"

	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/middleware"
	"github.com/ASeegull/edriver-space/model"
	"github.com/ASeegull/edriver-space/service"
	"github.com/labstack/echo/v4"
)

type UsersHandlers struct {
	usersService service.Users
	cfg          *config.Config
}

func NewUsersHandlers(usersService service.Users, cfg *config.Config) *UsersHandlers {
	return &UsersHandlers{
		usersService: usersService,
		cfg:          cfg,
	}
}

func (h *UsersHandlers) InitUsersRoutes(e *echo.Group, mw middleware.Middleware) {
	users := e.Group("/users")

	users.POST("/sign-in", h.SignIn())
	users.POST("/sign-out", h.SignOut())
	users.POST("/sign-up", h.SignUp())
	users.GET("/refresh-tokens", h.RefreshTokens())

	authenticated := users.Group("/", mw.UserIdentity())

	authenticated.POST("add-driver-licence", h.AddDriverLicence())
}

type singInInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type tokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (h *UsersHandlers) SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input singInInput

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "input body has not json format")
		}

		if err := c.Validate(input); err != nil {
			return c.JSON(http.StatusBadRequest, "invalid input body")
		}

		tokens, err := h.usersService.SignIn(c.Request().Context(), service.UserSignInInput{
			Email:    input.Email,
			Password: input.Password,
		})

		if err != nil {
			if errors.Is(err, model.ErrUserNotFound) {
				return c.JSON(http.StatusBadRequest, err.Error())
			}
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.SetCookie(h.createCookie(tokens.RefreshToken))

		return c.JSON(http.StatusOK, tokenResponse{
			AccessToken:  tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
		})
	}
}

type signUpInput struct {
	Firstname string `json:"firstname" validate:"required,max=64"`
	Lastname  string `json:"lastname" validate:"required,max=64"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required,min=8,max=32"`
}

func (h *UsersHandlers) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input signUpInput

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "input body has not json format")
		}

		if err := c.Validate(input); err != nil {
			return c.JSON(http.StatusBadRequest, "invalid input body")
		}

		tokens, err := h.usersService.SignUp(c.Request().Context(), service.UserSignUpInput{
			Firstname: input.Firstname,
			Lastname:  input.Lastname,
			Email:     input.Email,
			Password:  input.Password,
		})

		if err != nil {
			if errors.Is(err, model.ErrUserNotFound) {
				return c.JSON(http.StatusBadRequest, err.Error())
			}
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.SetCookie(h.createCookie(tokens.RefreshToken))

		return c.JSON(http.StatusOK, tokenResponse{
			AccessToken:  tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
		})
	}
}

func (h *UsersHandlers) SignOut() echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(h.cfg.Cookie.Name)
		if err != nil {
			if err != http.ErrNoCookie {
				return c.JSON(http.StatusBadRequest, err.Error())
			}

			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		if err := h.usersService.DeleteSession(c.Request().Context(), cookie.Value); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.SetCookie(h.deleteCookie())

		return c.JSON(http.StatusOK, "signOut")
	}
}

func (h *UsersHandlers) RefreshTokens() echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(h.cfg.Cookie.Name)
		if err != nil {
			if err != http.ErrNoCookie {
				return c.JSON(http.StatusBadRequest, err.Error())
			}

			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		tokens, err := h.usersService.RefreshTokens(c.Request().Context(), cookie.Value)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		c.SetCookie(h.createCookie(tokens.RefreshToken))

		return c.JSON(http.StatusOK, tokenResponse{
			AccessToken:  tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
		})
	}
}

type addDriverLicenceInput struct {
	IndividualTaxNumber string `json:"individual_tax_number" validate:"required"`
}

func (h UsersHandlers) AddDriverLicence() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input addDriverLicenceInput

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "input body has not json format")
		}

		if err := c.Validate(input); err != nil {
			return c.JSON(http.StatusBadRequest, "invalid input body")
		}

		userId, ok := c.Get("userId").(string)
		if !ok {
			return c.JSON(http.StatusUnauthorized, "user id is not present in context")
		}

		if err := h.usersService.AddDriverLicence(c.Request().Context(), service.AddDriverLicenceInput{
			IndividualTaxNumber: input.IndividualTaxNumber,
		}, userId); err != nil {
			if errors.Is(err, model.ErrDriverLicenceNotFound) {
				return c.JSON(http.StatusBadRequest, err.Error())
			}
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, "successfully added")
	}
}

func (h *UsersHandlers) createCookie(refreshToken string) *http.Cookie {
	return &http.Cookie{
		Name:     h.cfg.Cookie.Name,
		Value:    refreshToken,
		MaxAge:   h.cfg.Cookie.MaxAge * 60,
		Path:     h.cfg.Cookie.Path,
		HttpOnly: h.cfg.Cookie.HTTPOnly,
		Secure:   h.cfg.Cookie.Secure,
	}
}

func (h *UsersHandlers) deleteCookie() *http.Cookie {
	return &http.Cookie{
		Name:     h.cfg.Cookie.Name,
		Value:    "",
		MaxAge:   -1,
		Path:     h.cfg.Cookie.Path,
		HttpOnly: h.cfg.Cookie.HTTPOnly,
		Secure:   h.cfg.Cookie.Secure,
	}
}
