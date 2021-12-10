package http

import (
	"fmt"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/internal/auth"
	"github.com/ASeegull/edriver-space/internal/models"
	"github.com/ASeegull/edriver-space/internal/session"
	"github.com/ASeegull/edriver-space/pkg/csrf"
	"github.com/ASeegull/edriver-space/pkg/httpErrors"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
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

		userWithToken, err := h.authUC.Login(ctx, user)
		if err != nil {
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		newSession := &models.Session{
			UserID: userWithToken.User.ID,
		}

		sessionID, err := h.sessUC.CreateSession(ctx, newSession, 300)
		if err != nil {
			log.Warn(err.Error())
			return err
		}

		c.SetCookie(&http.Cookie{
			Name:     h.cfg.Session.Name,
			Value:    sessionID,
			MaxAge:   h.cfg.Session.Expire,
			HttpOnly: h.cfg.Cookie.HTTPOnly,
			SameSite: 0,
		})

		log.Info("session_id=", sessionID)

		return c.JSON(http.StatusOK, userWithToken)
	}
}

func (h *authHandlers) Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println(c.Get("user"))

		cookie, err := c.Cookie("session_id")
		if err != nil {
			if err == http.ErrNoCookie {
				return c.JSON(http.StatusUnauthorized, err.Error())
			}
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		if err := h.sessUC.DeleteByID(c.Request().Context(), cookie.Value); err != nil {
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		c.SetCookie(&http.Cookie{
			Name:   cookie.Name,
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})

		return c.JSON(http.StatusOK, "ok")
	}
}

func (h authHandlers) GetMe() echo.HandlerFunc {
	return func(c echo.Context) error {
		user, ok := c.Get("user").(*models.User)
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
