package middleware

import (
	"demo/utility"
	"github.com/gogf/gf/v2/net/ghttp"
)

func AuthMiddleware(r *ghttp.Request) {
	tokenStr := r.Header.Get("Authorization")
	if tokenStr == "" {
		r.Response.WriteStatusExit(401, utility.GetLocaleT("authorization.not-found"))
	}

	claims, err := utility.ParseToken(tokenStr)
	if err != nil {
		r.Response.WriteStatusExit(401, err.Error())
	}

	userID := claims["user_id"]
	r.SetCtxVar("user_id", userID)

	r.Middleware.Next()
}
