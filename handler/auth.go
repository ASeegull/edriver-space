package handler

import (
	"errors"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/model"
	"github.com/ASeegull/edriver-space/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthHandlers struct {
	AuthService service.Auth
	cfg         *config.Config
}

func NewAuthHandlers(authService service.Auth, cfg *config.Config) *AuthHandlers {
	return &AuthHandlers{
		AuthService: authService,
		cfg:         cfg,
	}
}

type singInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type tokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (h *AuthHandlers) SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input singInInput

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "invalid input body")
		}

		tokens, err := h.AuthService.SignIn(c.Request().Context(), service.UserSignInInput{
			Login:    input.Email,
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
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (h *AuthHandlers) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input signUpInput

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "invalid input body")
		}

		tokens, err := h.AuthService.SignUp(c.Request().Context(), service.UserSignUpInput{
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

func (h *AuthHandlers) SignOut() echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(h.cfg.Cookie.Name)
		if err != nil {
			if err != http.ErrNoCookie {
				return c.JSON(http.StatusBadRequest, err.Error())
			}

			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		if err := h.AuthService.DeleteSession(c.Request().Context(), cookie.Value); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.SetCookie(h.deleteCookie())

		return c.JSON(http.StatusOK, "signOut")
	}
}

func (h *AuthHandlers) RefreshTokens() echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(h.cfg.Cookie.Name)
		if err != nil {
			if err != http.ErrNoCookie {
				return c.JSON(http.StatusBadRequest, err.Error())
			}

			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		tokens, err := h.AuthService.RefreshTokens(c.Request().Context(), cookie.Value)
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

func (h *AuthHandlers) createCookie(refreshToken string) *http.Cookie {
	return &http.Cookie{
		Name:     h.cfg.Cookie.Name,
		Value:    refreshToken,
		MaxAge:   h.cfg.Cookie.MaxAge * 60,
		Path:     h.cfg.Cookie.Path,
		HttpOnly: h.cfg.Cookie.HTTPOnly,
		Secure:   h.cfg.Cookie.Secure,
	}
}

func (h *AuthHandlers) deleteCookie() *http.Cookie {
	return &http.Cookie{
		Name:     h.cfg.Cookie.Name,
		Value:    "",
		MaxAge:   -1,
		Path:     h.cfg.Cookie.Path,
		HttpOnly: h.cfg.Cookie.HTTPOnly,
		Secure:   h.cfg.Cookie.Secure,
	}
}
