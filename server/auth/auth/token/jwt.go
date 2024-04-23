package token

import (
	"crypto/rsa"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateToken(accountID string, expire time.Duration) (string, error)
type JwtTokenGen struct {
	privateKey *rsa.PrivateKey
	issuer     string
	nowFunc    func() time.Time
}

func NewJwtTokenGen(issuer string, privateKey *rsa.PrivateKey) *JwtTokenGen {
	return &JwtTokenGen{
		issuer:     issuer,
		nowFunc:    time.Now,
		privateKey: privateKey,
	}
}

func (t *JwtTokenGen) GenerateToken(accountID string, expire time.Duration) (string, error) {
	nowSecd := t.nowFunc()
	// 创建jwt 明文部分
	tkn := jwt.NewWithClaims(jwt.SigningMethodES512, jwt.StandardClaims{
		Issuer:    t.issuer,
		IssuedAt:  nowSecd.Unix(),
		ExpiresAt: nowSecd.Unix() + int64(expire.Seconds()),
		Subject:   accountID,
	})
	// 签名部分
	return tkn.SignedString(t.privateKey)
}
