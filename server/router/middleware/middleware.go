package middleware

import (
	"errors"
	"github.com/labstack/echo/v4"
)

type Middleware struct {
}

const userIDKey = "userID"

func (m *Middleware) SettraPUserIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//traP部員のUserIDの取得
		userID := c.Request().Header.Get("X-Showcase-User")
		if userID == "" {
			return nil
		}

		c.Set(userIDKey, userID)

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
