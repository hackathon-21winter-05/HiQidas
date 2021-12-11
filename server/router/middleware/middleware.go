package middleware

import (
	"errors"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type Middleware struct{}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

const userIDKey = "userID"

func (m *Middleware) CheckLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
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
