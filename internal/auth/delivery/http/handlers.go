package http

import (
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/internal/auth"
	"github.com/ASeegull/edriver-space/internal/models"
	"github.com/ASeegull/edriver-space/internal/session"
	"github.com/ASeegull/edriver-space/pkg/csrf"
	"github.com/ASeegull/edriver-space/pkg/httpErrors"
	"github.com/ASeegull/edriver-space/pkg/utils"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type authHandlers struct {
	authUC auth.UseCase
	sessUC session.UCSession
	cfg    *config.Config
}

func NewAuthHandlers(authUC auth.UseCase, sessUC session.UCSession, cfg *config.Config) auth.Handlers {
	return &authHandlers{authUC: authUC, sessUC: sessUC, cfg: cfg}
}

func (h *authHandlers) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		user := &models.User{}

		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		uwt, err := h.authUC.Login(ctx, user)
		if err != nil {
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		ttl := time.Duration(h.cfg.Tokens.RefreshTokenTTL) * time.Minute
		if err := h.sessUC.CreateSession(ctx, uwt.User.Id, uwt.Tokens.RefreshToken, ttl); err != nil {
			log.Warn(err.Error())
			return err
		}

		log.Info("New access token - ", uwt.Tokens.AccessToken)

		c.SetCookie(utils.CreateCookie(uwt.Tokens.RefreshToken, h.cfg))

		return c.JSON(http.StatusOK, map[string]string{
			"accessToken":  uwt.Tokens.AccessToken,
		})
	}
}

func (h *authHandlers) Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		refreshToken, err := c.Cookie(h.cfg.Cookie.Name)
		if err != nil {
			if err == http.ErrNoCookie {
				return c.JSON(http.StatusUnauthorized, err.Error())
			}
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		if err := h.sessUC.DeleteByID(c.Request().Context(), refreshToken.Value); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.SetCookie(utils.DeleteCookie(refreshToken.Name))

		return c.JSON(http.StatusOK, "logout")
	}
}

func (h authHandlers) GetMe() echo.HandlerFunc {
	return func(c echo.Context) error {
		user, ok := c.Get("user").(map[string]string)
		if !ok {
			return c.JSON(http.StatusUnauthorized, ok)
		}

		return c.JSON(http.StatusOK, user)
	}
}

func (h authHandlers) GetCSRFToken() echo.HandlerFunc {
	return func(c echo.Context) error {
		sid, ok := c.Get("sid").(string)

		if !ok {
			return c.JSON(http.StatusUnauthorized, "do not have sid")
		}

		token := csrf.MakeToken(sid)

		log.Info(token)

		c.Request().Header.Set(csrf.CSRFHeader, token)
		// give access to a client for reading the CSRF token from the request
		c.Request().Header.Set("Access-Control-Expose-Headers", csrf.CSRFHeader)

		return c.NoContent(http.StatusOK)
	}
}

//func (h authHandlers) Welcome() echo.HandlerFunc {
//	return func(c echo.Context) error {
//		cookie, err := c.Cookie("jwt_token")
//		if err != nil {
//			if err == http.ErrNoCookie {
//				return c.JSON(http.StatusUnauthorized, "no cookie")
//			}
//			return c.JSON(http.StatusUnauthorized, err.Error())
//		}
//
//		claims, err := utils.ExtractClaimsFromJWT(cookie.Value)
//		if err != nil {
//			return c.JSON(http.StatusBadRequest, err)
//		}
//
//		return c.JSON(http.StatusOK, claims)
//	}
//}

func (h authHandlers) RefreshTokens() echo.HandlerFunc {
	return func(c echo.Context) error {

		cookie, err := c.Cookie(h.cfg.Cookie.Name)
		if err != nil {
			if err == http.ErrNoCookie {
				return c.JSON(http.StatusUnauthorized, err.Error())
			}
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		newTokens, err := h.sessUC.RefreshSession(c.Request().Context(), cookie.Value)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		c.SetCookie(utils.CreateCookie(newTokens.RefreshToken, h.cfg))

		return c.JSON(http.StatusOK, map[string]string{
			"accessToken":  newTokens.AccessToken,
		})
	}
}
