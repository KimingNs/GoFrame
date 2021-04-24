package router

import (
	"GoGF/app/api"
	"GoGF/app/demo"
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

		//路由注册-对象注册
		s.BindObject("POST:/api/demo", api.Object{})
		//新接口路径
		s.BindObject("POST:/demo/third/", demo.ThirdApi{})
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
