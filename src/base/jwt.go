package base

import "github.com/dgrijalva/jwt-go"

const (
	ServiceJwt        = "jwt"
	JwtClaimTimeStamp = "timestamp"

	ErrTokenInvalid = "ErrTokenInvalid"
	ErrTokenExpiry  = "ErrTokenExpiry"

	HeaderAuth = "Authorization"
)

var IJwt IServiceJwt

type IServiceJwt interface {
	Sign(claims jwt.MapClaims) (string, error)
	Validate(token string) (jwt.MapClaims, error)
	Refresh(token string) (string, error)
	SetSecret(secret string)
}
