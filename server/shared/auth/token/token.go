package token

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/golang-jwt/jwt"
)

type JwtVerifier struct {
	PublicKey *ecdsa.PublicKey
}

// 验证jwt
func (v *JwtVerifier) Verfier(token string) (string, error) {
	t, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return v.PublicKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("", err)
	}
	if !t.Valid {
		return "", fmt.Errorf("token not valid")
	}
	clm, ok := t.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", fmt.Errorf("token claim is not stander cliams")
	}
	if err := clm.Valid(); err != nil {
		return "", fmt.Errorf("claim not valid")
	}
	return clm.Subject, nil
}
