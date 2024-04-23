package auth

import (
	"context"
	"coolcar/shared/auth/token"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
		}
	}
	return i.HandlerRequest, nil
}

type tokenVerifiler interface {
	Verify(token string) (string, error)
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
	return handler(ctx, req)
}
