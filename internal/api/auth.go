package api

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

type JwtData struct {
	UserID string
	Role   string
}

func (h *api) GenerateAccessJWT(data JwtData) (string, error) {
	_, tokenString, err := h.auth.Encode(jwt.MapClaims{
		"sub":  data.UserID,
		"role": data.Role,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	})
	return tokenString, err
}

func (h *api) GenerateRefreshJWT(data JwtData) (string, error) {
	_, tokenString, err := h.auth.Encode(jwt.MapClaims{
		"sub": data.UserID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	return tokenString, err
}

func (h *api) JWT(next http.Handler) http.Handler {
	return jwtauth.Verifier(h.auth)(jwtauth.Authenticator(next))
}

func GetJwtData(r *http.Request) JwtData {
	_, claims, _ := jwtauth.FromContext(r.Context())
	userID, _ := claims["sub"].(string)
	role, _ := claims["role"].(string)
	return JwtData{UserID: userID, Role: role}
}

func TokenFromHeader(r *http.Request) string {
	return jwtauth.TokenFromHeader(r)
}

func (j *JwtData) IsRefreshToken() bool {
	return j.Role == ""
}

func (j *JwtData) HasAccess(roles... string) bool {
	if j.Role == "admin" {
		return true
	}

	for _, role := range roles {
		if j.Role == role {
			return true
		}
	}

	return false
}