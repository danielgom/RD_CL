package api

import (
	"RD-Clone-NAPI/internal/config"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

const (
	AuthorizationTokenHeader = "Authorization"
)

var (
	errUnexpectedJWTSigningMethod = errors.New("unexpected jwt signing method")

	URLsToSkip = map[string]struct{}{
		"/v1/health": {},
	}
)

// JWTMiddleware is a middleware that forces response Content-Type.
func JWTMiddleware() func(next http.Handler) http.Handler {
	jwtConfig := config.Load().JWT
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if _, exists := URLsToSkip[r.URL.Path]; exists {
				next.ServeHTTP(w, r)
				return
			}

			authHeader := r.Header.Get(AuthorizationTokenHeader)
			headerParts := strings.Split(authHeader, " ")
			if len(headerParts) != 2 || headerParts[0] != "Bearer" {
				renderAs(w, r, userError("invalid token", authHeader))
				return
			}

			signedToken := headerParts[1]
			parse, err := jwt.Parse(signedToken, func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errUnexpectedJWTSigningMethod
				}
				return []byte(jwtConfig.Key), nil
			})
			if err != nil {
				renderAs(w, r, userError("error getting parsed token", err.Error()))
				return
			}

			if !parse.Valid {
				renderAs(w, r, userError("invalid token", err.Error()))
				return
			}

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
