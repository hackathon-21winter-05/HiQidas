package ws

import "github.com/labstack/echo/v4"

func (wh *WSHandlerGroup) ConnectHeyaWS(c echo.Context) error {
	return wh.s.ConnectHeyaWS(c)
}

func (wh *WSHandlerGroup) ConnectYjsWS(c echo.Context) error {
	return wh.s.ConnectYjsWS(c)
}

func (wh *WSHandlerGroup) ConnectParserWS(c echo.Context) error {
	return wh.s.ConnectParserWS(c)
}
