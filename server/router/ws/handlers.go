package ws

import "github.com/labstack/echo/v4"

func (wh *WSHandlerGroup) ConnectHeyaWS(c echo.Context) error {
	return wh.s.ConnectHeyaWS(c)
}
