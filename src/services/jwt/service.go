package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/springmove/sptty"
	"github.com/springmove/tp/src/base"
)

type Service struct {
	sptty.BaseService

	cfg    Config
	secret string
}

func (s *Service) Init(app sptty.ISptty) error {
	if err := app.GetConfig(s.ServiceName(), &s.cfg); err != nil {
		return err
	}

	return nil
}

func (s *Service) ServiceName() string {
	return base.ServiceJwt
}

func (s *Service) Sign(claims jwt.MapClaims) (string, error) {
	claims[base.JwtClaimTimeStamp] = time.Now().Format(time.RFC3339)
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(s.secret))
}

func (s *Service) validateExpiry(claims jwt.MapClaims) error {
	if s.cfg.Expiry == (0 * time.Second) {
		return nil
	}

	ts, _ := time.Parse(time.RFC3339, claims[base.JwtClaimTimeStamp].(string))
	if time.Now().After(ts.Add(s.cfg.Expiry)) {
		return fmt.Errorf(base.ErrTokenExpiry)
	}

	return nil
}

func (s *Service) Validate(token string) (jwt.MapClaims, error) {

	claims, err := s.parse(token)
	if err != nil {
		return nil, err
	}

	if err := s.validateExpiry(claims); err != nil {
		return nil, err
	}

	return claims, nil
}

func (s *Service) parse(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected Signing Method: %v", token.Header["alg"])
		}

		return []byte(s.secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf(base.ErrTokenInvalid)
	}
}

func (s *Service) Refresh(token string) (string, error) {
	claims, err := s.parse(token)
	if err != nil {
		return "", err
	}

	return s.Sign(claims)
}

func (s *Service) SetSecret(secret string) {
	s.secret = secret
}
