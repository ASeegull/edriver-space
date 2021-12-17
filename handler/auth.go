package handler

import (
	"errors"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/models"
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
			if errors.Is(err, models.ErrUserNotFound) {
				return c.JSON(http.StatusBadRequest, err.Error())
			}
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.SetCookie(&http.Cookie{
			Name:   h.cfg.Cookie.Name,
			Value:  tokens.RefreshToken,
			Path:   h.cfg.Cookie.Path,
			MaxAge: h.cfg.Cookie.MaxAge * 60,
		})

		return c.JSON(http.StatusOK, tokenResponse{
			AccessToken:  tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
		})
	}
}

func (h *AuthHandlers) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "signUp")
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

		c.SetCookie(&http.Cookie{
			Name:   h.cfg.Cookie.Name,
			MaxAge: -1,
		})

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

		c.SetCookie(&http.Cookie{
			Name:   h.cfg.Cookie.Name,
			Value:  tokens.RefreshToken,
			Path:   h.cfg.Cookie.Path,
			MaxAge: h.cfg.Cookie.MaxAge * 60,
		})

		return c.JSON(http.StatusOK, tokenResponse{
			AccessToken:  tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
		})
	}
}
