package router

import (
	"GoGF/app/api"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareErrorHandler)
		s.Group("/demo", func(group *ghttp.RouterGroup) {
			group.Middleware(MiddlewareErrorHandler)
			group.ALL("/*", func(r *ghttp.Request) {
				panic("db error: sql is xxxxxxx")
			})
		})
		//group.ALL("/hello", api.Hello)

		//路由注册-对象注册
		s.BindObject("POST:/demo", api.Object{})
	})
}

func MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	fmt.Println(r.Response.Status)
	fmt.Println(http.StatusInternalServerError)
	if r.Response.Status >= http.StatusInternalServerError {
		r.Response.ClearBuffer()
		r.Response.Writeln("服务器居然开小差了，请稍后再试吧！")
	}
}
