package api

import "github.com/gogf/gf/net/ghttp"

type Controller struct{}

func (c *Controller) Index(r *ghttp.Request) {
	r.Response.Write("index")
}

func (c *Controller) Show(r *ghttp.Request) {
	r.Response.Write("show")
}
