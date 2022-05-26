package authenticator

import (
	"context"

	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/auth/domain"
	"github.com/golang-jwt/jwt"
)

type JWTAuthenticator interface {
	GenerateAccessToken(ctx context.Context, user domain.User) (token domain.TokenData, err error)
	ValidateAccessToken(ctx context.Context, tokenStr string) (token JWTClaims, err error)
}

const ACCESS_TOKEN_DURATION = 3600 * 24 //in seconds
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

type JWTClaims struct {
	jwt.StandardClaims
	Username string `json:"name"`
	Admin    bool   `json:"admin"`
	Group    string `json:"group"`
	UserID   int    `json:"user_id"`
}
