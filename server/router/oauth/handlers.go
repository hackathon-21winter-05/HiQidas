package oauth

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/antihax/optional"
	"github.com/gofrs/uuid"
	"github.com/gorilla/sessions"
	"github.com/hackathon-21winter-05/HiQidas/server/protobuf/rest"
	"github.com/hackathon-21winter-05/HiQidas/server/router/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/sapphi-red/go-traq"
	"github.com/thanhpk/randstr"
)

const oauthCodeRedirect = "https://q.trap.jp/api/v3/oauth2/authorize"

// GetOauthCallbackHandler GET /oauth/callback ハンドラ
func (oh *OauthHandlerGroup) GetOauthCallbackHandler(c echo.Context) error {
	verifier := randstr.String(64)
	hash := sha256.Sum256([]byte(verifier))
	challenge := base64.RawURLEncoding.EncodeToString(hash[:])

	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	sess.Values["verifier"] = verifier
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60,
		HttpOnly: true,
	}
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var clientID string
	if strings.Contains(c.Request().Header.Get("referer"), "localhost") {
		clientID = oh.c.DevClientID
	} else {
		clientID = oh.c.ClientID
	}

	uri := fmt.Sprintf("%s?response_type=code&client_id=%s&code_challenge=%s&code_challenge_method=S256", oauthCodeRedirect, clientID, challenge)
	oauthRedirectData := &rest.GetOauthCallbackResponse{
		Uri: uri,
	}

	return utils.SendProtobuf(c, http.StatusOK, oauthRedirectData)
}

// PostOauthCodeHandler POST /oauth/code ハンドラ
func (oh *OauthHandlerGroup) PostOauthCodeHandler(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	oauthCodeData := &rest.PostOauthCodeRequest{}
	err = utils.BindProtobuf(c, oauthCodeData)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	var clientID string
	if strings.Contains(c.Request().Header.Get("referer"), "localhost") {
		clientID = oh.c.DevClientID
	} else {
		clientID = oh.c.ClientID
	}

	verifier := sess.Values["verifier"].(string)
	opts := &traq.Oauth2ApiPostOAuth2TokenOpts{
		Code:         optional.NewString(oauthCodeData.GetCode()),
		ClientId:     optional.NewString(clientID),
		CodeVerifier: optional.NewString(verifier),
	}
	token, res, err := oh.cli.Oauth2Api.PostOAuth2Token(context.Background(), "authorization_code", opts)
	if err != nil || token.AccessToken == "" || res.StatusCode >= 400 {
		return echo.NewHTTPError(res.StatusCode, err)
	}

	auth := context.WithValue(context.Background(), traq.ContextAccessToken, token.AccessToken)
	v, res, err := oh.cli.MeApi.GetMe(auth)
	if err != nil || res.StatusCode != http.StatusOK {
		return echo.NewHTTPError(res.StatusCode, err)
	}

	userUUID, err := uuid.FromString(v.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	user, _ := oh.ser.GetUserByID(c.Request().Context(), userUUID)

	if user == nil {
		iconFileUUID, err := uuid.FromString(v.IconFileId)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		_, err = oh.ser.CreateTraPUser(c.Request().Context(), userUUID, iconFileUUID, v.Name)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}

	sess.Values["userid"] = userUUID.String()
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   int(token.ExpiresIn),
		HttpOnly: true,
	}
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}
