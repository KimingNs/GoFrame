package api

import (
	"github.com/gogf/gf/net/ghttp"
)

type Object struct{}

func (c *Object) Show(r *ghttp.Request) {
	r.Response.Write("show")
}

func (c *Object) TestObject(r *ghttp.Request) {
	r.Response.Write("test-object")
}
