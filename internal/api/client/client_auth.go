package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/kaatinga/plantbook/internal/api/middleware"
	"github.com/kaatinga/plantbook/internal/api/models"

	"github.com/pkg/errors"
)

const (
	cliAuthURL string = "/user/login"
)

// Auth request session cookie from api by pass loginPassword credentials,
// returns cookie with plantbook_token and httpCode.
func (a *API) Auth(ctx context.Context, params models.UserLoginPassword) (*http.Cookie, int, error) {
	// make body request
	paramsBts, err := json.Marshal(params)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.WithMessage(err, "json.Marshal error")
	}
	bodyRequest := bytes.NewReader(paramsBts)
	r, err := http.NewRequest("POST", a.urlWithQuery(cliAuthURL, ""), bodyRequest)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.WithMessage(err, "make request error")
	}
	r.Header.Add("Content-Type", "application/json")
	// send request
	resp, err := a.handler.Do(r)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.WithMessage(err, "request error")
	}
	defer resp.Body.Close()
	// process response
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			body = []byte{}
		}
		return nil, resp.StatusCode, errors.Errorf("error, http code=%d, body=%s", resp.StatusCode, string(body))
	}
	// extract cookie
	for _, cookie := range resp.Cookies() {
		if cookie.Name == middleware.JWTCookieName {
			return cookie, resp.StatusCode, nil
		}
	}
	return nil, resp.StatusCode, errors.New("no cookie with session_id")
}
