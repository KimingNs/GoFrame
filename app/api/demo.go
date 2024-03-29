package api

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Object struct{}

type Req struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Res struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (c *Object) Show(r *ghttp.Request) {
	r.Response.Write("show")
}

func (c *Object) TestObject(r *ghttp.Request) {
	r.Response.Write("test-object")
}

func (c *Object) User(r *ghttp.Request) {
	if j, err := gjson.DecodeToJson(r.GetBodyString()); err != nil {
		fmt.Println(err)
	} else {
		req := Req{Name: j.GetString("name"), Password: j.GetString("password")}
		//返回json格式
		err := r.Response.WriteJson(Res{Code: 200, Message: r.Cookie.Get("sessionId"), Data: req})
		if err != nil {
			return
		}
	}
}

func (c *Object) TestHttp(r *ghttp.Request) {
	if _, err := gjson.DecodeToJson(r.GetBodyString()); err != nil {
		fmt.Println(err)
	} else {
		res := g.Client().SetCookie("sessionId", "zhangtianming").ContentJson().PostContent("http://127.0.0.1:8199/demo/user", g.Map{
			"name":     "testHttp_ztm",
			"password": "http200",
		})
		err := r.Response.WriteJson(res)
		if err != nil {
			return
		}
	}
}
