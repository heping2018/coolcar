package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const privateKey = "MIIEpAIBAAKCAQEAglrInXu5X5j/Or2g77/f3rsne2J90w/i7zFpVtcz2IneH+ulHrBBfkVpDnaA7ni4dPrz2m97tg6OenVH0mxHMbt1bxqhOVK/24KdD/jpOr6skLXyo702ymUek/5fE8F36z0kXfb5WU5ZOJ3b4KTCdPEey6fs03+rQqzEKF4DK3EHnIXjtJL5Xw2LoCOblQyX92ZtdAQG1E2x8ldIZumqzm+lwv5bmBETViyetwrdEpzfHpirc88UM3LogkBvNvt1yWf0GXAthC5bBZ1QZ1f/npiHuRGWIITFY5yMWbEnLZwhr4z5/od4pjXo6rxbLJvpeBKzzHB127KVDJqDeaoZgwIDAQABAoIBABhuopZd17dSyle8rhrxqCirhHFsWvAB1MOHS9qieE24PYFXHfo8B/J+WPwEexL3Xn9Sf/0rkxDi9pJ+Q+ltMQThwVeDMr+3Qk2G1CBnw3MmxNpUt/c3ojraLBqJ/VELHIpjSdswzLiP5kol0E3xI59eiCqcvVdA4R0cm6GDSW3NO5KnchA/8Tu5wBRATh0++jAHdIRdeDk0IcUctJUocqHHd2ikcJLlvGxzWN4EybyZJzMMMP5vmZDrBN4IvRu1cH9vBDG7PmusdG6ljE5SdSaLCrSPiSsrXO3Yn37aGWOwhNBzOdEudB28MqtQaqXcuH9zOmas3Mypl+UZpsrUEb0CgYEAwB9vxBToxvxy2UC2QO4qm3uSM9K4x8wZta/o9yhNhK4LcF3+2oZj19c3dr33sol3OKxNsg8nxL02cIKoRFX5vkBuJb0ZGaOzCwLNaPzOfJmWgFfYrBY6jRm4uKU5aT3uapj+swiftm7JxCDFONKYwlvXaDOytaCMp8D2di1SUZ0CgYEArbHwUECD6Y2MJLd45HOHWayIB5LkokXqitkp6cr5JUxv0CWXs4SbsCOQeiwpy0t9ynSZGynOoHhm4nP1LYGIgxIRP+BTyBcmb3040YEs5SY0HWq92p0VJXRXh2VyftLjgXGxRiZS2qtgAX6XnEP0c8/YQxsgkaFQE5uIoKJDPZ8CgYEAne65v+S4XitUdgdeSw399IambveAcq3zTIno5pQ49SLlwF9ki1sCZIJE09Xh8uBSI0JxNDiFJpOsVtzxWgubG6x3X4qNQahyHFEXboCzdXYEZEjSktRLGYbVdORNx5fjj7lpVt91+1AjiJivx8BHVy6MatpkxC3Qsm5LrGYhT3kCgYEAoRshTh54/AD5UvkGXcc2bJuV1IiQHl+BK2Y/9QT50Hm0YDjuG2lNSrPCBz+9SD3uIAHEHSjK6ZajvzlU2O+Doib47ulwbG/ki5Z3RANvk9+6iOp/zCzU91eQ9BnJeA69TEwa5HkZco9wThKQJzX0oBLRgXTA4bLg5j7mHDBE8/sCgYAGjK80X75TQ6/qNUiKzrfyliP12se6kT2tC/vP+IQKqh+yc7hVFIwUU82V/f3dSahVabpik87/YPtZXdx4JXhfogg9pDjWwmBY8pfLmKYzr0U0tYKtsfsPKjlGbZGNnKm5i8w8GKuIZb881q+QqZOM1/yKOzXeyC0+I0+Ojd0lVQ=="

func TestNewJwtTokenGen(t *testing.T) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		t.Fatalf("cannot parse private key %v", err)
	}
	g := NewJwtTokenGen("coolcar", key)
	g.nowFunc = func() time.Time {
		return time.Unix(1516239022, 0)
	}
	tkn, err := g.GenerateToken("123123", 7200)
	if err != nil {
		t.Errorf("cannnot gennerate token %v", err)
	}
	want := ""
	if tkn != want {
		t.Errorf("run token err %q,%q", want, tkn)
	}

}
