package auth

import (
	"context"
	"coolcar/shared/auth/token"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/internal/metadata"
	"google.golang.org/grpc/status"
)

const (
	authorizationHeader = "authorization"
	bearerPrefix        = "Bearer "
)

// 验证token
func Interceptor(publicKeyFile string) (grpc.UnaryServerInterceptor, error) {
	f, err := os.Open(publicKeyFile)
	if err != nil {
		return nil, fmt.Errorf("cannot opne file")
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("cannot ReadAll file")
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("cannot parseECPublicKeyFromPEM pubKey")
	}
	i := &interceptor{
		PublicKey: pubKey,
		verifier: &token.JwtVerifier{
			PublicKey: pubKey,
		},
	}
	return i.HandlerRequest, nil
}

type tokenVerifiler interface {
	Verfier(token string) (string, error)
}

// UnaryServerInfo
type interceptor struct {
	PublicKey *rsa.PublicKey
	verifier  tokenVerifiler
}

func (i *interceptor) HandlerRequest(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	tkn, err := tokenFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}
	aid, err := i.verifier.Verfier(tkn)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token not valid: %v", err)
	}
	return handler(ContextWithAccountId(ctx, AccountId(aid)), req)
}

func tokenFromContext(c context.Context) (string, error) {
	m, ok := metadata.FromIncomingContext(c)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "")
	}
	tkn := ""
	for _, v := range m[authorizationHeader] {
		if strings.HasPrefix(v, bearerPrefix) {
			// 截取
			tkn = v[len(bearerPrefix):]
		}
	}
	if tkn == "" {
		return "", status.Error(codes.Unauthenticated, "")
	}
	return tkn, nil
}

type accountIdKey struct {
}

type AccountId string

func (a AccountId) String() string {
	return string(a)
}

func ContextWithAccountId(c context.Context, aid AccountId) context.Context {
	return context.WithValue(c, accountIdKey{}, aid)
}

func AccountIdFromContext(c context.Context) (AccountId, error) {
	v := c.Value(accountIdKey{})
	aid, ok := v.(AccountId)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "")
	}
	return aid, nil

}
