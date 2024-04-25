package token

import (
	"testing"

	"github.com/golang-jwt/jwt"
)

const PublicKey = ""

func TestVerfier(t *testing.T) {
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(PublicKey))
	if err != nil {
		t.Fatalf("cannot parse pubkey")
	}
	v := &JwtVerifier{
		PublicKey: pubKey,
	}
	tkn := ""

	accountId, err := v.Verfier(tkn)
	if err != nil {
		t.Errorf("Verfier is fail")
	}
	want := ""
	if want != accountId {
		t.Errorf("wromh account id . want %q,accountId %q", want, accountId)
	}

}
