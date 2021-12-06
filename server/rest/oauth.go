package rest

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/hackathon-21winter-05/HiQidas/server/protobuf/rest"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
)

type oauthHandlers interface {
	GetOauthCallbackHandler(c echo.Context) error
}

const oauthCodeRedirect = "https://q.trap.jp/api/v3/oauth2/authorize"

func (r *restHandlersImpl) GetOauthCallbackHandler(c echo.Context) error {
	verifier := randstr.String(64)
	hash := sha256.Sum256([]byte(verifier))
	challenge := base64.RawURLEncoding.EncodeToString(hash[:])

	sess, _ := session.Get("session", c)
	sess.Values["verifier"] = verifier
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60,
		HttpOnly: true,
	}
	err := sess.Save(c.Request(), c.Response())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var clientID string
	if strings.Contains(c.Request().Header.Get("referer"), "localhost") {
		clientID = r.c.Client_ID_Dev
	} else {
		clientID = r.c.Client_ID
	}

	uri := fmt.Sprintf("%s?response_type=code&client_id=%s&code_challenge=%s&code_challenge_method=S256", oauthCodeRedirect, clientID, challenge)
	redirectData := &rest.GetOauthCallbackResponse{
		Uri: uri,
	}

	return sendProtobuf(c, http.StatusOK, redirectData)
}
