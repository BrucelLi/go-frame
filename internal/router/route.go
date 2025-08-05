package router

import (
	"demo/internal/controller/hello"
	"demo/internal/controller/job"
	"demo/internal/controller/token"
	"demo/internal/controller/user"
	"demo/internal/middleware"
	"github.com/gogf/gf/v2/net/ghttp"
)

func RegisterRoutes(s *ghttp.Server) {
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse, middleware.I18nMiddleware, middleware.CorsMiddleware)
		group.Group("/v1", func(v1Group *ghttp.RouterGroup) {
			v1Group.Group("/public", func(publicGroup *ghttp.RouterGroup) {
				publicGroup.Bind(token.NewV1())
				publicGroup.Bind(hello.NewV1())
			})
			v1Group.Group("", func(v1Group *ghttp.RouterGroup) {
				v1Group.Middleware(middleware.AuthMiddleware)
				v1Group.Bind(user.NewV1())
				v1Group.Bind(job.NewV1())
			})
		})
	})
}
