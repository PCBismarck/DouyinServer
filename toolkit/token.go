package toolkit

import (
	"fmt"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TikTokClaims struct {
	Id       uint
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

type TokenMap struct {
	m sync.Map
}

func (t *TokenMap) Delete(key uint) {
	t.m.Delete(key)
}

func (t *TokenMap) Store(key uint, value string) {
	t.m.Store(key, value)
}

func (t *TokenMap) Load(key uint) (value string, ok bool) {
	v, ok := t.m.Load(key)
	if v != nil {
		value = v.(string)
	}
	return value, ok
}

var secretKey = []byte("BiteDance")

var TokenManger TokenMap

func VerifyToken(tokenStr string) (ok bool, err error) {
	ttc, err := ParseToken(tokenStr)
	if err != nil {
		return false, err
	}
	if tokenStored, ok := TokenManger.Load(ttc.Id); ok && tokenStored == tokenStr {
		return true, nil
	}
	return false, nil
}

func DeleteTokenByUid(uid uint) {
	TokenManger.Delete(uid)
}

func StoreToken(uid uint, tokenStr string) {
	TokenManger.Store(uid, tokenStr)
}

// use JWT to generate token
func GenerateToken(id uint, username string, password string) (tokenStr string, err error) {
	claims := TikTokClaims{
		Id:       id,
		Username: username,
		Password: password,
	}
	claims.RegisteredClaims = jwt.RegisteredClaims{
		NotBefore: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		Issuer:    "Tiktok",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

// parse JWT token to TikTokClaims
func ParseToken(tokenStr string) (claims *TikTokClaims, err error) {
	tokenClaims, err := jwt.ParseWithClaims(
		tokenStr, &TikTokClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
	if err != nil {
		return nil, err
	}
	if tokenClaims == nil {
		return nil, fmt.Errorf("empty claims")
	}
	claims, ok := tokenClaims.Claims.(*TikTokClaims)
	if !ok {
		return nil, fmt.Errorf("bad claims")
	}
	return claims, nil
}
