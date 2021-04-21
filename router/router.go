package router

import (
	"GoGF/app/api"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareErrorHandler)
		s.Group("/:route", func(group *ghttp.RouterGroup) {
			group.Middleware(MiddlewareErrorHandler)
			group.ALL("/*", func(r *ghttp.Request) {
				panic("find error page")
			})
		})
		//group.ALL("/hello", api.Hello)

		//路由注册-对象注册
		s.BindObject("POST:/demo", api.Object{})
	})
}

func MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if r.Response.Status >= http.StatusInternalServerError {
		if err := r.GetError(); err != nil {
			g.Log("error").Error(err)
		}
		r.Response.ClearBuffer()
		r.Response.Writeln("服务器外出打怪升级，至今未回，如有发现，重金奖赏")
	}
}
