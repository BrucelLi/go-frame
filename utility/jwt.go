package utility

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/golang-jwt/jwt/v5"
)

func getJwtKey() []byte {
	secret := g.Cfg().MustGet(gctx.GetInitCtx(), "jwt.secret").String()
	return []byte(secret)
}

func CreateToken(userID interface{}) (string, error) {
	now := time.Now().Unix()
	expTime := g.Cfg().MustGet(gctx.GetInitCtx(), "jwt.expire").Int64()
	exp := now + expTime
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     exp,
		"iat":     now,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getJwtKey())
}

func ParseToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, gerror.New(GetLocaleT("authorization.sign-method-unavailable"))
		}
		return getJwtKey(), nil
	})
	if err != nil || !token.Valid {
		return nil, gerror.New(GetLocaleT("authorization.token-unavailable"))
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, gerror.New(GetLocaleT("authorization.invalid-load"))
	}
	return claims, nil
}
