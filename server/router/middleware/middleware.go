package middleware

import (
	"errors"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/service"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type Middleware struct {
	s *service.Service
}

func NewMiddleware(s *service.Service) *Middleware {
	return &Middleware{s: s}
}

const userIDKey = "userID"

func (m *Middleware) CheckLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		userID := sess.Values["userid"]
		if userID == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Please Login")
		}

		return next(c)
	}
}

func (m *Middleware) GetUserID(c echo.Context) (string, error) {
	rowUserID := c.Get(userIDKey)
	userID, ok := rowUserID.(string)
	if !ok {
		return "", errors.New("invalid context userID")
	}

	return userID, nil
}

func (m *Middleware) DeleteHeyaCheckAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		heyaIDStr := c.Param("heyaID")
		heyaID, err := uuid.FromString(heyaIDStr)
		if err != nil {
			c.Logger().Info(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		sess, err := session.Get("session", c)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		userIDstr := sess.Values["userid"].(string)
		if userIDstr == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Please Login")
		}

		userID, err := uuid.FromString(userIDstr)
		if err != nil {
			c.Logger().Info(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		heya, err := m.s.GetHeyaByID(c.Request().Context(), heyaID)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		if heya.CreatorID != userID {
			return echo.NewHTTPError(http.StatusForbidden, "You are not a creator")
		}

		return next(c)
	}
}
