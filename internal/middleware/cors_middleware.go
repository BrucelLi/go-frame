package middleware

import "github.com/gogf/gf/v2/net/ghttp"

func CorsMiddleware(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
