package authenticator

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/abdussalamfaqih/rest-sqlite/internal/appconfig"
	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/auth/domain"
	"github.com/golang-jwt/jwt"
)

type jwtAuthenticator struct {
	cfg appconfig.AppConfig
}

func NewJWTAuthenticator(cfg appconfig.AppConfig) JWTAuthenticator {
	return &jwtAuthenticator{
		cfg: cfg,
	}
}

func (auth *jwtAuthenticator) GenerateAccessToken(ctx context.Context, user domain.User) (token domain.TokenData, err error) {

	expires := time.Now().Add(time.Duration(ACCESS_TOKEN_DURATION) * time.Second).Unix()
	claims := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    auth.cfg.Name,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expires,
			Subject:   strconv.Itoa(int(user.ID)),
		},
		Username: user.Name,
		UserID:   int(user.ID),
		Admin:    user.Name == "admin",
		Group:    "users",
	}

	tokenJWT := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)

	signedToken, err := tokenJWT.SignedString([]byte(auth.cfg.Secret_Key))
	if err != nil {
		return
	}

	token.TokenType = "Bearer"
	token.Token = signedToken
	token.ExpiresAt = int(expires)

	return
}

func (auth *jwtAuthenticator) ValidateAccessToken(ctx context.Context, tokenStr string) (token JWTClaims, err error) {

	tokenJwt, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("signing method invalid")
		}

		return []byte(auth.cfg.Secret_Key), nil
	})

	if err != nil {

		return token, err
	}

	claims, ok := tokenJwt.Claims.(jwt.MapClaims)
	if !ok || !tokenJwt.Valid {
		return
	}

	exp := claims["exp"].(float64)
	iat := claims["iat"].(float64)

	result := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    claims["iss"].(string),
			ExpiresAt: int64(exp),
			IssuedAt:  int64(iat),
		},
		UserID: int(claims["user_id"].(float64)),
	}
	return result, nil
}
