package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	t := GenerateJwtToken("hello")
	fmt.Printf("token: %v\n\n", t)
	uis, err := ValidateJwt(t)
	fmt.Printf("userId: %s", uis)
	if err != nil {
		fmt.Println("err: ", err)
	}
}

func GenerateJwtToken(uniqueId string) string {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, newClaim(uniqueId))
	secureKey := getSecureKey()
	fmt.Printf("secure key: %+v\n\n", secureKey)
	token, er := jwtToken.SignedString(secureKey)
	if er != nil {
		fmt.Printf("%v\n\n", er)
	}
	return token
}

// ValidateJwt accepts jwt toke from header and returns the unique id for that token
// returns error if not valid
func ValidateJwt(headerToken string) (string, error) {
	t, err := jwt.ParseWithClaims(headerToken, &jwtWebClaim{}, getKeyFunc())
	if err != nil {
		fmt.Println("error here", err)
		return "", err
	}
	claims, ok := t.Claims.(*jwtWebClaim)
	if !ok {
		return "", errors.New("couldn't parse claims")
	}
	return claims.UniqueId, nil
}

func getSecureKey() interface{} {
	key, _ := jwt.ParseRSAPrivateKeyFromPEM([]byte(v))
	return key
}

func getKeyFunc() func(token *jwt.Token) (interface{}, error) {
	return func(token *jwt.Token) (interface{}, error) {
		return getSecureKey(), nil
	}
}

func newClaim(uniqueId string) jwtWebClaim {
	return jwtWebClaim{
		UniqueId: uniqueId, StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 3*60*1000,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "kreditbee.in",
		},
	}
}

type jwtWebClaim struct {
	UniqueId string `json:"unique_id"`
	jwt.StandardClaims
}

const (
	v = `-----BEGIN PRIVATE KEY-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCXn99ID9DDD6kS
GMXYKsj7h5UQWLI1YSwpNzpCOzwjP+z3UpKU1LVlAqlhzdIMNGOlf60FeeKMNOJo
kUhDRgWWSlfjq1+qNCRAAHQVMvtiLcjsGB/aU9Aa3mTgg7pOzZjIu/Zw93ryoa7M
PIjmEQL7YeRHEx8XuI3e9WyZn4m7u8e0v2IiiaI1jUKtULnzoSuc6PX9FyN5YgRf
/aS6GA4GKV085q21Oo70nCfB+j3Sfzyxv3lUfyqc4soFDykG/Is9IIFgEOb4roOK
zctR9VtxSTKyb3cVpSAtA6m5szbonF1aID5i7rKngY4elK1eSdb+jR+bBIwxI8sW
Syp6h0AfAgMBAAECggEAZ3VQ92zDN7YB2TVKGhvnk6mJzuOWhdHOPjlO6U5d1HeU
C5YuKpuRQmE0jmXQQz+kFpMtziTCroSPUGaBdlXbDhegLNsMpah6a6lji+uDmBRB
msdDRMgXoZ3KG2AGyiqDa8TFfYOAQvqLuRQ9HS7SaBD9oHmtp5PoShAQkCpJ4uxL
gOKlXgTvGUvw8T08QHN5hWwHfokTNeerFpYBVo8IUNB6AI3x5ABhj0wwCuLDhv8c
u/yL7H5HoM/38AF8wYCcrw1zK2gc0gHbnTMxqexNL8S4bqKZOWJ3O+2Td7ynNHDw
8k+Hx7w6hOP2dAnnGxK3rt5UKFNkA++BZrMO5Tc8gQKBgQDJxAPEW0ztHQ356/ux
U9WPUWf1mGZ85gsE26X1nksOInOuMgz3jWlFcw5mDQ5DpjiYmH1i79BuNRqI+SpI
b4qt2xaW1LtkEcsCycHTQcAgtW5Ou4Jkis19d2LOxe8eF8Aqkyd/67JoKnTTJXyC
+iEuWpOiboOJ6Cl22o0mV+6GPwKBgQDAYYV1Z0bkHhn9Ja+uuWvERYS0uFwkpvKm
FU9oDLJc6vPlVIRaTsdZ+OqgqEEFn7TYIWcjfz6KlQ5VFVH/l1tXrWMnDNURuP6+
WLPcJJmISuIRdQxUmENQmkhANzzFArAKgIgLuE2Z6E5fz9b3JiYxWarEgCsRGsgg
bE77H2+OIQKBgQCMepuc1WzEEtyuS+3cU2B3/tgBGXESOSEm6r+sOeBMIRSmDlcU
7TKElk1KJIDv/QWeyV1Ty2E1umeVQtZ7xJ2r2sTpk4g9bl4IvOzk5/ybSXdZ/hgV
ZMcaOktjaFDKhQLZIcf4uuRmvljEwm7kyr7bhVkRWRzgofJsvFJWhhafBQKBgQCA
mvDfuim+54ySGTZnhFbRf3OmaDRY4C4H7ukFcq9txDdFUIml4VINCIS436GQA+Ke
NT5AkZiaheht4nHNfj42z5cgDMkHLvdFAgFCokjsvrp/1xJmHt+pK2ovW33JafGL
I79OrmdaAt0Z2dQnph4UEZBCCjhAF/o/CoiGAc7OYQKBgQCfwd41qkXtYDCoNqNl
KRuutY0E9JNIlseqpfeVdzm/r4RcGaNLvCgRHhJd3rMJQirL3U9mS0J/JH+2F/zS
KWjl9fzIXa8j0n/afTAIqp0n2gM9Q1ADUW4Jr2xMzllCbDIe+kkITR9boNp3e2xN
vVoQ04MrazhrMImkQNXHdOmDBQ==
-----END PRIVATE KEY-----`
)
