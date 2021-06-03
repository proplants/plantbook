package client

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/kaatinga/plantbook/internal/api/middleware"
	"github.com/kaatinga/plantbook/internal/api/models"
)

const (
	timeoutTestAPI time.Duration = 10 * time.Second
)

func TestAPI_Auth(t *testing.T) {
	// prepare test servers
	tsOK := httptest.NewServer(http.HandlerFunc(dummyPlantbookServer))
	tsFail := httptest.NewServer(http.HandlerFunc(dummyPlantbookServerMissCookie))
	// set context
	ctx, cancel := context.WithTimeout(context.Background(), timeoutTestAPI*10)
	defer cancel()
	// params
	uLogin := "root"
	uPassword := "love"
	uEmptyLogin := ""
	uEmptyPassword := ""
	wantCookie := &http.Cookie{Name: middleware.JWTCookieName}
	// test cases
	tests := []struct {
		name       string
		ctx        context.Context
		url        string
		apiVersion string
		params     models.UserLoginPassword
		want       *http.Cookie
		want1      int
		wantErr    bool
	}{
		{"200_fillLoginPass_errFalse", ctx, tsOK.URL, APIV1URL, models.UserLoginPassword{Login: &uLogin, Password: &uPassword}, wantCookie, http.StatusOK, false},
		{"400_emptyLoginPass_errTrue", ctx, tsOK.URL, APIV1URL, models.UserLoginPassword{Login: &uEmptyLogin, Password: &uEmptyPassword}, wantCookie, http.StatusBadRequest, true},
		{"200_missCookie_errTrue", ctx, tsFail.URL, APIV1URL, models.UserLoginPassword{Login: &uLogin, Password: &uPassword}, wantCookie, http.StatusOK, true},
		{"500_any_errTrue", ctx, "http://127.0.0.1", APIV1URL, models.UserLoginPassword{Login: &uLogin, Password: &uPassword}, wantCookie, http.StatusInternalServerError, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api, err := New(tt.url, tt.apiVersion, timeoutTestAPI)
			if err != nil {
				t.Errorf("unexpected new api error, %s", err)
				return
			}
			got, got1, err := api.Auth(tt.ctx, tt.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.Auth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got1 != tt.want1 {
				t.Errorf("API.Auth() got1 = %v, want %v", got1, tt.want1)
			}
			if tt.wantErr {
				return
			}
			if tt.want != nil && (got.Name != tt.want.Name) {
				t.Errorf("API.Auth() got = %v, want %v", got, tt.want)
			}

		})
	}
}

// dummyPlantbookServer
// need more similar as origin server,
// Auth checks only empty or not login&password
//
func dummyPlantbookServer(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case APIV1URL + cliAuthURL:
		if !methodCheck(http.MethodPost, w, r) {
			return
		}
		loginPass := &models.UserLoginPassword{}
		err := json.NewDecoder(r.Body).Decode(loginPass)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, `{"message":"`+err.Error()+`"}`)
			return
		}
		// I hope not nil pointer
		loginIsEmpty := *loginPass.Login == ""
		passwordIsEmpty := *loginPass.Password == ""
		if loginIsEmpty || passwordIsEmpty {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"message":"empty login/password not allowed"}`)
			return
		}
		c := &http.Cookie{
			Name:    middleware.JWTCookieName,
			Value:   defaultSessionID,
			Expires: time.Now().AddDate(0, 0, 1),
		}
		http.SetCookie(w, c)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"message":"unsupported path `+path+`"}`)
	}
}

// dummyPlantbookServerMissCookie
// need more similar as origin server,
// Auth checks only empty or not login&password and doesn't set cookie
func dummyPlantbookServerMissCookie(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case APIV1URL + cliAuthURL:
		if !methodCheck(http.MethodPost, w, r) {
			return
		}
		loginPass := &models.UserLoginPassword{}
		err := json.NewDecoder(r.Body).Decode(loginPass)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, `{"message":"`+err.Error()+`"}`)
			return
		}
		// I hope not nil pointer
		loginIsEmpty := *loginPass.Login == ""
		passwordIsEmpty := *loginPass.Password == ""
		if loginIsEmpty || passwordIsEmpty {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"message":"empty login/password not allowed"}`)
			return
		}
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"message":"unsupported path `+path+`"}`)
	}
}

// helpers

const (
	defaultSessionID string = "plantbook_session:customvalue"
	unauthorized     string = "Unauthorized"
)

func methodCheck(refMethod string, w http.ResponseWriter, r *http.Request) bool {
	if r.Method != refMethod {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "error: wrong http method"+r.Method+", allowed "+refMethod)
		return false
	}
	return true
}

func parameterCheck(paramName string, w http.ResponseWriter, r *http.Request) bool {
	parValue := r.URL.Query().Get(paramName)
	if parValue == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "error: missed query parameter "+paramName+" value")
		return false
	}
	return true
}

func checkAuth(w http.ResponseWriter, r *http.Request) bool {
	session, err := r.Cookie(middleware.JWTCookieName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "error: "+err.Error())
		return false
	}
	if session.Value == "" || session.Value != defaultSessionID {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, unauthorized)
		return false
	}
	return true
}
