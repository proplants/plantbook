// Package middleware contains http.Handler functions
package middleware

import (
	"context"
	"net/http"

	"github.com/kaatinga/plantbook/pkg/logging"
	"github.com/kaatinga/plantbook/utils/randutil"
	"go.uber.org/zap"
)

const (
	requestIDLength int = 32
	// XRequestIDHeader header name.
	XRequestIDHeader string = "X-Request-Id"
	JWTCookieName    string = "plantbook_token"
)

// RequestID middlware which adds random X-Request-Id header if it not exists.
func RequestID(ctx context.Context, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do some middleware logic here
		requestID := GetRequestID(r)
		if requestID == "" {
			requestID = randutil.RandStringRunes(requestIDLength)
		}
		r.Header.Set(XRequestIDHeader, requestID)
		log := logging.FromContext(r.Context()).With(zap.String("request_id", requestID))
		ctx = logging.WithLogger(ctx, log)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetRequestID extracts X-Request-Id value from htt.Request.Header.
func GetRequestID(r *http.Request) string {
	return r.Header.Get(XRequestIDHeader)
}

// GetCookie extracts session cookie from http.Requiest.
func GetCookie(r *http.Request) (*http.Cookie, error) {
	return r.Cookie(JWTCookieName)
}
