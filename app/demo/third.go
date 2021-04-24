package demo

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type ThirdApi struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    g.Map  `json:"data"`
}

func (c *ThirdApi) SelfCheck(r *ghttp.Request) {
	request_json, _ := gjson.DecodeToJson(r.GetBodyString())
	partner_info, _ := g.DB().Model("t_partner_info").Where("partner_token=", request_json.GetString("partner_token")).One()
	fmt.Println(partner_info.GMap())
}
