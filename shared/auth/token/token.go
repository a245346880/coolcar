package token

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

// JWTTokenVerifier JWT token 认证
type JWTTokenVerifier struct {
	PublicKey *rsa.PublicKey
}

// Verify 验证token
func (v *JWTTokenVerifier) Verify(token string) (string, error) {
	t, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return v.PublicKey, nil
		})
	if err != nil {
		return "", fmt.Errorf("无法解析token:%v", err)
	}
	if !t.Valid {
		return "", fmt.Errorf("token验证失败")
	}
	clm, ok := t.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", fmt.Errorf("token 签名解析失败")
	}
	if err := clm.Valid(); err != nil {
		return "", fmt.Errorf("token 签名验证失败：%v", err)
	}
	return clm.Subject, nil
}
